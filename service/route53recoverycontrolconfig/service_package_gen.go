// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/conns"
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
			Factory:  resourceCluster,
			TypeName: "aws_route53recoverycontrolconfig_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  resourceControlPanel,
			TypeName: "aws_route53recoverycontrolconfig_control_panel",
			Name:     "Control Panel",
		},
		{
			Factory:  resourceRoutingControl,
			TypeName: "aws_route53recoverycontrolconfig_routing_control",
			Name:     "Routing Control",
		},
		{
			Factory:  resourceSafetyRule,
			TypeName: "aws_route53recoverycontrolconfig_safety_rule",
			Name:     "Safety Rule",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53RecoveryControlConfig
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
