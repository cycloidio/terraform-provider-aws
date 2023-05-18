// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceAccelerator,
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccelerator,
			TypeName: "aws_globalaccelerator_accelerator",
			Name:     "Accelerator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
			Name:     "Custom Routing Accelerator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceCustomRoutingEndpointGroup,
			TypeName: "aws_globalaccelerator_custom_routing_endpoint_group",
		},
		{
			Factory:  ResourceCustomRoutingListener,
			TypeName: "aws_globalaccelerator_custom_routing_listener",
		},
		{
			Factory:  ResourceEndpointGroup,
			TypeName: "aws_globalaccelerator_endpoint_group",
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_globalaccelerator_listener",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GlobalAccelerator
}

var ServicePackage = &servicePackage{}
