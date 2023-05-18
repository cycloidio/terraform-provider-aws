// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package acmpca

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
			Factory:  DataSourceCertificate,
			TypeName: "aws_acmpca_certificate",
		},
		{
			Factory:  DataSourceCertificateAuthority,
			TypeName: "aws_acmpca_certificate_authority",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_acmpca_certificate",
		},
		{
			Factory:  ResourceCertificateAuthority,
			TypeName: "aws_acmpca_certificate_authority",
			Name:     "Certificate Authority",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceCertificateAuthorityCertificate,
			TypeName: "aws_acmpca_certificate_authority_certificate",
		},
		{
			Factory:  ResourcePermission,
			TypeName: "aws_acmpca_permission",
		},
		{
			Factory:  ResourcePolicy,
			TypeName: "aws_acmpca_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ACMPCA
}

var ServicePackage = &servicePackage{}
