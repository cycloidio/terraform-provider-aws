// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ssmcontacts_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
	"github.com/hashicorp/terraform-provider-aws/conns"
	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceRotation,
			Name:    "Rotation",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceRotation,
			Name:    "Rotation",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceContact,
			TypeName: "aws_ssmcontacts_contact",
		},
		{
			Factory:  DataSourceContactChannel,
			TypeName: "aws_ssmcontacts_contact_channel",
		},
		{
			Factory:  DataSourcePlan,
			TypeName: "aws_ssmcontacts_plan",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceContact,
			TypeName: "aws_ssmcontacts_contact",
			Name:     "Context",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceContactChannel,
			TypeName: "aws_ssmcontacts_contact_channel",
			Name:     "Contact Channel",
		},
		{
			Factory:  ResourcePlan,
			TypeName: "aws_ssmcontacts_plan",
			Name:     "Plan",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSMContacts
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ssmcontacts_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ssmcontacts_sdkv2.NewFromConfig(cfg,
		ssmcontacts_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
