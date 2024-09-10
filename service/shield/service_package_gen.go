// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package shield

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/conns"
	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceProtection,
			Name:    "Protection",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newApplicationLayerAutomaticResponseResource,
			Name:    "Application Layer Automatic Response",
		},
		{
			Factory: newDRTAccessLogBucketAssociationResource,
			Name:    "DRT Log Bucket Association",
		},
		{
			Factory: newDRTAccessRoleARNAssociationResource,
			Name:    "DRT Role ARN Association",
		},
		{
			Factory: newProactiveEngagementResource,
			Name:    "Proactive Engagement",
		},
		{
			Factory: newResourceSubscription,
			Name:    "Subscription",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceProtection,
			TypeName: "aws_shield_protection",
			Name:     "Protection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceProtectionGroup,
			TypeName: "aws_shield_protection_group",
			Name:     "Protection Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "protection_group_arn",
			},
		},
		{
			Factory:  ResourceProtectionHealthCheckAssociation,
			TypeName: "aws_shield_protection_health_check_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Shield
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
