// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elbv2

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	elasticloadbalancingv2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_alb",
			Name:     "Load Balancer",
		},
		{
			Factory:  dataSourceListener,
			TypeName: "aws_alb_listener",
			Name:     "Listener",
		},
		{
			Factory:  dataSourceTargetGroup,
			TypeName: "aws_alb_target_group",
			Name:     "Target Group",
		},
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_lb",
			Name:     "Load Balancer",
		},
		{
			Factory:  dataSourceHostedZoneID,
			TypeName: "aws_lb_hosted_zone_id",
			Name:     "Hosted Zone ID",
		},
		{
			Factory:  dataSourceListener,
			TypeName: "aws_lb_listener",
			Name:     "Listener",
		},
		{
			Factory:  dataSourceTargetGroup,
			TypeName: "aws_lb_target_group",
			Name:     "Target Group",
		},
		{
			Factory:  dataSourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
		},
		{
			Factory:  dataSourceLoadBalancers,
			TypeName: "aws_lbs",
			Name:     "Load Balancers",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_alb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_alb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceListenerCertificate,
			TypeName: "aws_alb_listener_certificate",
			Name:     "Listener Certificate",
		},
		{
			Factory:  resourceListenerRule,
			TypeName: "aws_alb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTargetGroup,
			TypeName: "aws_alb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_alb_target_group_attachment",
			Name:     "Target Group Attachment",
		},
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_lb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_lb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceListenerCertificate,
			TypeName: "aws_lb_listener_certificate",
			Name:     "Listener Certificate",
		},
		{
			Factory:  resourceListenerRule,
			TypeName: "aws_lb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTargetGroup,
			TypeName: "aws_lb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_lb_target_group_attachment",
			Name:     "Target Group Attachment",
		},
		{
			Factory:  resourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTrustStoreRevocation,
			TypeName: "aws_lb_trust_store_revocation",
			Name:     "Trust Store Revocation",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELBV2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticloadbalancingv2_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return elasticloadbalancingv2_sdkv2.NewFromConfig(cfg,
		elasticloadbalancingv2_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
