// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package batch

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/internal/vcr"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newJobDefinitionDataSource,
			TypeName: "aws_batch_job_definition",
			Name:     "Job Definition",
			Tags:     unique.Make(inttypes.ServicePackageResourceTags{}),
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newJobQueueResource,
			TypeName: "aws_batch_job_queue",
			Name:     "Job Queue",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
			Identity: inttypes.RegionalARNIdentity(inttypes.WithIdentityDuplicateAttrs(names.AttrID)),
			Import: inttypes.FrameworkImport{
				WrappedImport: true,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceComputeEnvironment,
			TypeName: "aws_batch_compute_environment",
			Name:     "Compute Environment",
			Tags:     unique.Make(inttypes.ServicePackageResourceTags{}),
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceJobQueue,
			TypeName: "aws_batch_job_queue",
			Name:     "Job Queue",
			Tags:     unique.Make(inttypes.ServicePackageResourceTags{}),
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceSchedulingPolicy,
			TypeName: "aws_batch_scheduling_policy",
			Name:     "Scheduling Policy",
			Tags:     unique.Make(inttypes.ServicePackageResourceTags{}),
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceComputeEnvironment,
			TypeName: "aws_batch_compute_environment",
			Name:     "Compute Environment",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceJobDefinition,
			TypeName: "aws_batch_job_definition",
			Name:     "Job Definition",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
			Identity: inttypes.RegionalARNIdentity(
				inttypes.WithIdentityDuplicateAttrs(names.AttrID),
				inttypes.WithMutableIdentity(),
			),
			Import: inttypes.SDKv2Import{
				CustomImport: true,
			},
		},
		{
			Factory:  resourceSchedulingPolicy,
			TypeName: "aws_batch_scheduling_policy",
			Name:     "Scheduling Policy",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Batch
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*batch.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*batch.Options){
		batch.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *batch.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *batch.Options) {
			if inContext, ok := conns.FromContext(ctx); ok && inContext.VCREnabled() {
				tflog.Info(ctx, "overriding retry behavior to immediately return VCR errors")
				o.Retryer = conns.AddIsErrorRetryables(cfg.Retryer().(aws.RetryerV2), retry.IsErrorRetryableFunc(vcr.InteractionNotFoundRetryableFunc))
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return batch.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*batch.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*batch.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *batch.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*batch.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
