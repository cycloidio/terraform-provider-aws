// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package worklink

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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceFleet,
			TypeName: "aws_worklink_fleet",
		},
		{
			Factory:  ResourceWebsiteCertificateAuthorityAssociation,
			TypeName: "aws_worklink_website_certificate_authority_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.WorkLink
}

var ServicePackage = &servicePackage{}
