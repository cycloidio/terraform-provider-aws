// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDetector,
			TypeName: "aws_guardduty_detector",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDetector,
			TypeName: "aws_guardduty_detector",
			Name:     "Detector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFilter,
			TypeName: "aws_guardduty_filter",
			Name:     "Filter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceInviteAccepter,
			TypeName: "aws_guardduty_invite_accepter",
		},
		{
			Factory:  ResourceIPSet,
			TypeName: "aws_guardduty_ipset",
			Name:     "IP Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_guardduty_member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_guardduty_organization_admin_account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_guardduty_organization_configuration",
		},
		{
			Factory:  ResourcePublishingDestination,
			TypeName: "aws_guardduty_publishing_destination",
		},
		{
			Factory:  ResourceThreatIntelSet,
			TypeName: "aws_guardduty_threatintelset",
			Name:     "Threat Intel Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GuardDuty
}

var ServicePackage = &servicePackage{}
