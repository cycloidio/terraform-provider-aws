// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package globalaccelerator

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
			Factory: newAcceleratorDataSource,
			Name:    "Accelerator",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCrossAccountAttachmentResource,
			Name:    "Cross-account Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
			Name:     "Custom Routing Accelerator",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccelerator,
			TypeName: "aws_globalaccelerator_accelerator",
			Name:     "Accelerator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceCustomRoutingAccelerator,
			TypeName: "aws_globalaccelerator_custom_routing_accelerator",
			Name:     "Custom Routing Accelerator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceCustomRoutingEndpointGroup,
			TypeName: "aws_globalaccelerator_custom_routing_endpoint_group",
			Name:     "Custom Routing Endpoint Group",
		},
		{
			Factory:  resourceCustomRoutingListener,
			TypeName: "aws_globalaccelerator_custom_routing_listener",
			Name:     "Custom Routing Listener",
		},
		{
			Factory:  resourceEndpointGroup,
			TypeName: "aws_globalaccelerator_endpoint_group",
			Name:     "Endpoint Group",
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_globalaccelerator_listener",
			Name:     "Listener",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GlobalAccelerator
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
