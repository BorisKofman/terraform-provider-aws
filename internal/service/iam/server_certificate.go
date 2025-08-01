// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iam

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	awstypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_iam_server_certificate", name="Server Certificate")
// @Tags(identifierAttribute="name", resourceType="ServerCertificate")
// @Testing(existsType="github.com/aws/aws-sdk-go-v2/service/iam/types;types.ServerCertificate", tlsKey=true, importStateId="rName", importIgnore="private_key")
func resourceServerCertificate() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceServerCertificateCreate,
		ReadWithoutTimeout:   resourceServerCertificateRead,
		UpdateWithoutTimeout: resourceServerCertificateUpdate,
		DeleteWithoutTimeout: resourceServerCertificateDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceServerCertificateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_body": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressNormalizeCertRemoval,
				StateFunc:        StateTrimSpace,
			},
			names.AttrCertificateChain: {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressNormalizeCertRemoval,
				StateFunc:        StateTrimSpace,
			},
			"expiration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrName: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{names.AttrNamePrefix},
				ValidateFunc:  validation.StringLenBetween(0, 128),
			},
			names.AttrNamePrefix: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{names.AttrName},
				ValidateFunc:  validation.StringLenBetween(0, 128-id.UniqueIDSuffixLength),
			},
			names.AttrPath: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
			names.AttrPrivateKey: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: suppressNormalizeCertRemoval,
				StateFunc:        StateTrimSpace,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"upload_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerCertificateCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMClient(ctx)

	sslCertName := create.Name(d.Get(names.AttrName).(string), d.Get(names.AttrNamePrefix).(string))
	input := &iam.UploadServerCertificateInput{
		CertificateBody:       aws.String(d.Get("certificate_body").(string)),
		PrivateKey:            aws.String(d.Get(names.AttrPrivateKey).(string)),
		ServerCertificateName: aws.String(sslCertName),
		Tags:                  getTagsIn(ctx),
	}

	if v, ok := d.GetOk(names.AttrCertificateChain); ok {
		input.CertificateChain = aws.String(v.(string))
	}

	if v, ok := d.GetOk(names.AttrPath); ok {
		input.Path = aws.String(v.(string))
	}

	output, err := conn.UploadServerCertificate(ctx, input)

	// Some partitions (e.g. ISO) may not support tag-on-create.
	partition := meta.(*conns.AWSClient).Partition(ctx)
	if input.Tags != nil && errs.IsUnsupportedOperationInPartitionError(partition, err) {
		input.Tags = nil

		output, err = conn.UploadServerCertificate(ctx, input)
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating IAM Server Certificate (%s): %s", sslCertName, err)
	}

	d.SetId(aws.ToString(output.ServerCertificateMetadata.ServerCertificateId))
	d.Set(names.AttrName, sslCertName) // Required for resource Read.

	// For partitions not supporting tag-on-create, attempt tag after create.
	if tags := getTagsIn(ctx); input.Tags == nil && len(tags) > 0 {
		err := serverCertificateCreateTags(ctx, conn, sslCertName, tags)

		// If default tags only, continue. Otherwise, error.
		if v, ok := d.GetOk(names.AttrTags); (!ok || len(v.(map[string]any)) == 0) && errs.IsUnsupportedOperationInPartitionError(partition, err) {
			return append(diags, resourceServerCertificateRead(ctx, d, meta)...)
		}

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "setting IAM Server Certificate (%s) tags: %s", d.Id(), err)
		}
	}

	return append(diags, resourceServerCertificateRead(ctx, d, meta)...)
}

func resourceServerCertificateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMClient(ctx)

	cert, err := findServerCertificateByName(ctx, conn, d.Get(names.AttrName).(string))

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] IAM Server Certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IAM Server Certificate (%s): %s", d.Id(), err)
	}

	metadata := cert.ServerCertificateMetadata
	d.SetId(aws.ToString(metadata.ServerCertificateId))
	d.Set(names.AttrARN, metadata.Arn)
	d.Set("certificate_body", cert.CertificateBody)
	d.Set(names.AttrCertificateChain, cert.CertificateChain)
	if metadata.Expiration != nil {
		d.Set("expiration", aws.ToTime(metadata.Expiration).Format(time.RFC3339))
	} else {
		d.Set("expiration", nil)
	}
	d.Set(names.AttrName, metadata.ServerCertificateName)
	d.Set(names.AttrNamePrefix, create.NamePrefixFromName(aws.ToString(metadata.ServerCertificateName)))
	d.Set(names.AttrPath, metadata.Path)
	if metadata.UploadDate != nil {
		d.Set("upload_date", aws.ToTime(metadata.UploadDate).Format(time.RFC3339))
	} else {
		d.Set("upload_date", nil)
	}

	setTagsOut(ctx, cert.Tags)

	return diags
}

func resourceServerCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).IAMClient(ctx)

	if d.HasChanges(names.AttrName, names.AttrNamePrefix, names.AttrPath) {
		input := &iam.UpdateServerCertificateInput{}

		if d.HasChange(names.AttrName) {
			oldName, newName := d.GetChange(names.AttrName)

			// Handle both a name change and a switch to using a name prefix
			newSSLCertName := create.Name(newName.(string), d.Get(names.AttrNamePrefix).(string))

			input.ServerCertificateName = aws.String(oldName.(string))
			input.NewServerCertificateName = aws.String(newSSLCertName)
		} else if d.HasChange(names.AttrNamePrefix) {
			oldName := d.Get(names.AttrName).(string)

			// Handle only a name prefix change using an empty string as name (as it hasn't been changed)
			newSSLCertName := create.Name("", d.Get(names.AttrNamePrefix).(string))

			input.ServerCertificateName = aws.String(oldName)
			input.NewServerCertificateName = aws.String(newSSLCertName)
		}
		nameChanged := input.NewServerCertificateName != nil

		if d.HasChange(names.AttrPath) {
			if !nameChanged {
				name := d.Get(names.AttrName).(string)
				input.ServerCertificateName = aws.String(name)
			}
			input.NewPath = aws.String(d.Get(names.AttrPath).(string))
		}

		_, err := conn.UpdateServerCertificate(ctx, input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating IAM Server Certificate (%s): %s", d.Id(), err)
		}

		// If the name was changed, the new name must be set in the state for tag update that precedes resource read
		if nameChanged {
			d.Set(names.AttrName, input.NewServerCertificateName)
		}
	}

	return append(diags, resourceServerCertificateRead(ctx, d, meta)...)
}

func resourceServerCertificateDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).IAMClient(ctx)

	log.Printf("[DEBUG] Deleting IAM Server Certificate: %s", d.Id())
	_, err := tfresource.RetryWhenIsAErrorMessageContains[any, *awstypes.DeleteConflictException](ctx, d.Timeout(schema.TimeoutDelete), func(ctx context.Context) (any, error) {
		return conn.DeleteServerCertificate(ctx, &iam.DeleteServerCertificateInput{
			ServerCertificateName: aws.String(d.Get(names.AttrName).(string)),
		})
	}, "currently in use by arn")

	if errs.IsA[*awstypes.NoSuchEntityException](err) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting IAM Server Certificate (%s): %s", d.Id(), err)
	}

	return diags
}

func resourceServerCertificateImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	d.Set(names.AttrName, d.Id())
	// private_key can't be fetched from any API call
	return []*schema.ResourceData{d}, nil
}

func findServerCertificateByName(ctx context.Context, conn *iam.Client, name string) (*awstypes.ServerCertificate, error) {
	input := &iam.GetServerCertificateInput{
		ServerCertificateName: aws.String(name),
	}

	output, err := conn.GetServerCertificate(ctx, input)

	if errs.IsA[*awstypes.NoSuchEntityException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.ServerCertificate == nil || output.ServerCertificate.ServerCertificateMetadata == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.ServerCertificate, nil
}

func normalizeCert(cert any) string {
	if cert == nil || cert == (*string)(nil) {
		return ""
	}

	var rawCert string
	switch cert := cert.(type) {
	case string:
		rawCert = cert
	case *string:
		rawCert = aws.ToString(cert)
	default:
		return ""
	}

	cleanVal := sha1.Sum(stripCR([]byte(strings.TrimSpace(rawCert))))
	return hex.EncodeToString(cleanVal[:])
}

// strip CRs from raw literals. Lifted from go/scanner/scanner.go
// See https://github.com/golang/go/blob/release-branch.go1.6/src/go/scanner/scanner.go#L479
func stripCR(b []byte) []byte {
	c := make([]byte, len(b))
	i := 0
	for _, ch := range b {
		if ch != '\r' {
			c[i] = ch
			i++
		}
	}
	return c[:i]
}

// Terraform AWS Provider version 3.0.0 removed state hash storage.
// This DiffSuppressFunc prevents the resource from triggering needless recreation.
func suppressNormalizeCertRemoval(k, old, new string, d *schema.ResourceData) bool {
	return normalizeCert(new) == old
}

func serverCertificateTags(ctx context.Context, conn *iam.Client, identifier string, optFns ...func(*iam.Options)) ([]awstypes.Tag, error) {
	input := iam.ListServerCertificateTagsInput{
		ServerCertificateName: aws.String(identifier),
	}
	var output []awstypes.Tag

	pages := iam.NewListServerCertificateTagsPaginator(conn, &input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx, optFns...)

		if err != nil {
			return nil, err
		}

		output = append(output, page.Tags...)
	}

	return output, nil
}
