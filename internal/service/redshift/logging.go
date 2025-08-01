// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshift

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	awstypes "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkResource("aws_redshift_logging", name="Logging")
func newLoggingResource(_ context.Context) (resource.ResourceWithConfigure, error) {
	return &loggingResource{}, nil
}

const (
	ResNameLogging = "Logging"
)

type loggingResource struct {
	framework.ResourceWithModel[loggingResourceModel]
}

func (r *loggingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrBucketName: schema.StringAttribute{
				Optional: true,
			},
			names.AttrClusterIdentifier: schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			names.AttrID: framework.IDAttributeDeprecatedWithAlternate(path.Root(names.AttrClusterIdentifier)),
			"log_destination_type": schema.StringAttribute{
				Optional:   true,
				CustomType: fwtypes.StringEnumType[awstypes.LogDestinationType](),
			},
			"log_exports": schema.SetAttribute{
				Optional:    true,
				CustomType:  fwtypes.SetOfStringType,
				ElementType: types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						enum.FrameworkValidate[LogExports](),
					),
				},
			},
			names.AttrS3KeyPrefix: schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *loggingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	conn := r.Meta().RedshiftClient(ctx)

	var plan loggingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = types.StringValue(plan.ClusterIdentifier.ValueString())

	in := &redshift.EnableLoggingInput{}
	resp.Diagnostics.Append(flex.Expand(ctx, plan, in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retry InvalidClusterState faults, which can occur when logging is enabled
	// immediately after being disabled (ie. resource replacement).
	out, err := tfresource.RetryWhenIsAErrorMessageContains[any, *awstypes.InvalidClusterStateFault](ctx, propagationTimeout,
		func(ctx context.Context) (any, error) {
			return conn.EnableLogging(ctx, in)
		},
		"There is an operation running on the Cluster",
	)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.Redshift, create.ErrActionCreating, ResNameLogging, plan.ID.String(), err),
			err.Error(),
		)
		return
	}
	if out == nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.Redshift, create.ErrActionCreating, ResNameLogging, plan.ID.String(), nil),
			errors.New("empty output").Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *loggingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	conn := r.Meta().RedshiftClient(ctx)

	var state loggingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	out, err := findLoggingByID(ctx, conn, state.ID.ValueString())
	if tfresource.NotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.Redshift, create.ErrActionSetting, ResNameLogging, state.ID.String(), err),
			err.Error(),
		)
		return
	}

	// LogDestinationType is not returned correctly from the DescribeLoggingStatus API when
	// the type is "s3". Set attributes individually (rather than with AutoFlex) and skip setting
	// log_destination_type to avoid persistent differences.
	state.BucketName = flex.StringToFramework(ctx, out.BucketName)
	state.S3KeyPrefix = flex.StringToFramework(ctx, out.S3KeyPrefix)
	resp.Diagnostics.Append(flex.Flatten(ctx, out.LogExports, &state.LogExports)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *loggingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	conn := r.Meta().RedshiftClient(ctx)

	var plan, state loggingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.BucketName.Equal(state.BucketName) ||
		!plan.LogDestinationType.Equal(state.LogDestinationType) ||
		!plan.LogExports.Equal(state.LogExports) ||
		!plan.S3KeyPrefix.Equal(state.S3KeyPrefix) {
		in := &redshift.EnableLoggingInput{}
		resp.Diagnostics.Append(flex.Expand(ctx, plan, in)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Retry InvalidClusterState faults, which can occur when logging is enabled
		// immediately after being disabled (ie. resource replacement).
		out, err := tfresource.RetryWhenIsAErrorMessageContains[any, *awstypes.InvalidClusterStateFault](ctx, propagationTimeout,
			func(ctx context.Context) (any, error) {
				return conn.EnableLogging(ctx, in)
			},
			"There is an operation running on the Cluster",
		)
		if err != nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.Redshift, create.ErrActionUpdating, ResNameLogging, plan.ID.String(), err),
				err.Error(),
			)
			return
		}
		if out == nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.Redshift, create.ErrActionUpdating, ResNameLogging, plan.ID.String(), nil),
				errors.New("empty output").Error(),
			)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *loggingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	conn := r.Meta().RedshiftClient(ctx)

	var state loggingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	in := &redshift.DisableLoggingInput{
		ClusterIdentifier: state.ID.ValueStringPointer(),
	}

	// Retry InvalidClusterState faults, which can occur when logging is being enabled.
	_, err := tfresource.RetryWhenIsAErrorMessageContains[any, *awstypes.InvalidClusterStateFault](ctx, propagationTimeout,
		func(ctx context.Context) (any, error) {
			return conn.DisableLogging(ctx, in)
		},
		"There is an operation running on the Cluster",
	)
	if err != nil {
		if errs.IsA[*awstypes.ClusterNotFoundFault](err) {
			return
		}
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.Redshift, create.ErrActionDeleting, ResNameLogging, state.ID.String(), err),
			err.Error(),
		)
		return
	}
}

func (r *loggingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root(names.AttrID), req, resp)
	resource.ImportStatePassthroughID(ctx, path.Root(names.AttrClusterIdentifier), req, resp)
}

func findLoggingByID(ctx context.Context, conn *redshift.Client, id string) (*redshift.DescribeLoggingStatusOutput, error) {
	in := &redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: aws.String(id),
	}

	out, err := conn.DescribeLoggingStatus(ctx, in)
	if err != nil {
		if errs.IsA[*awstypes.ClusterNotFoundFault](err) {
			return nil, &retry.NotFoundError{
				LastError:   err,
				LastRequest: in,
			}
		}

		return nil, err
	}

	if out == nil {
		return nil, tfresource.NewEmptyResultError(in)
	}
	if !aws.ToBool(out.LoggingEnabled) {
		return nil, &retry.NotFoundError{
			LastError:   errors.New("logging not enabled"),
			LastRequest: in,
		}
	}

	return out, nil
}

type loggingResourceModel struct {
	framework.WithRegionModel
	BucketName         types.String                                    `tfsdk:"bucket_name"`
	ClusterIdentifier  types.String                                    `tfsdk:"cluster_identifier"`
	ID                 types.String                                    `tfsdk:"id"`
	LogDestinationType fwtypes.StringEnum[awstypes.LogDestinationType] `tfsdk:"log_destination_type"`
	LogExports         fwtypes.SetValueOf[types.String]                `tfsdk:"log_exports"`
	S3KeyPrefix        types.String                                    `tfsdk:"s3_key_prefix"`
}

// This enum should be defined in the AWS SDK, but is missing.
//
// See the Redshift documentation for valid values.
// https://docs.aws.amazon.com/redshift/latest/APIReference/API_EnableLogging.html
type LogExports string

// Enum values for LogExports
const (
	LogExportsConnectionLog   LogExports = "connectionlog"
	LogExportsUserActivityLog LogExports = "useractivitylog"
	LogExportsUserLog         LogExports = "userlog"
)

func (LogExports) Values() []LogExports {
	return []LogExports{
		LogExportsConnectionLog,
		LogExportsUserActivityLog,
		LogExportsUserLog,
	}
}
