// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package opensearch

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
			Factory:  DataSourceDomain,
			TypeName: "aws_opensearch_domain",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDomain,
			TypeName: "aws_opensearch_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceDomainPolicy,
			TypeName: "aws_opensearch_domain_policy",
		},
		{
			Factory:  ResourceDomainSAMLOptions,
			TypeName: "aws_opensearch_domain_saml_options",
		},
		{
			Factory:  ResourceInboundConnectionAccepter,
			TypeName: "aws_opensearch_inbound_connection_accepter",
		},
		{
			Factory:  ResourceOutboundConnection,
			TypeName: "aws_opensearch_outbound_connection",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpenSearch
}

var ServicePackage = &servicePackage{}
