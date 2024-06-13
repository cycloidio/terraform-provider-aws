// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appconfig

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	appconfig_sdkv2 "github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceEnvironment,
			Name:    "Environment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceConfigurationProfile,
			TypeName: "aws_appconfig_configuration_profile",
		},
		{
			Factory:  DataSourceConfigurationProfiles,
			TypeName: "aws_appconfig_configuration_profiles",
		},
		{
			Factory:  DataSourceEnvironment,
			TypeName: "aws_appconfig_environment",
		},
		{
			Factory:  DataSourceEnvironments,
			TypeName: "aws_appconfig_environments",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceApplication,
			TypeName: "aws_appconfig_application",
			Name:     "Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceConfigurationProfile,
			TypeName: "aws_appconfig_configuration_profile",
			Name:     "Configuration Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDeployment,
			TypeName: "aws_appconfig_deployment",
			Name:     "Deployment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDeploymentStrategy,
			TypeName: "aws_appconfig_deployment_strategy",
			Name:     "Deployment Strategy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceExtension,
			TypeName: "aws_appconfig_extension",
			Name:     "Extension",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceExtensionAssociation,
			TypeName: "aws_appconfig_extension_association",
		},
		{
			Factory:  ResourceHostedConfigurationVersion,
			TypeName: "aws_appconfig_hosted_configuration_version",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppConfig
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appconfig_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return appconfig_sdkv2.NewFromConfig(cfg, func(o *appconfig_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			tflog.Debug(ctx, "setting endpoint", map[string]any{
				"tf_aws.endpoint": endpoint,
			})
			o.BaseEndpoint = aws_sdkv2.String(endpoint)

			if o.EndpointOptions.UseFIPSEndpoint == aws_sdkv2.FIPSEndpointStateEnabled {
				tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
				o.EndpointOptions.UseFIPSEndpoint = aws_sdkv2.FIPSEndpointStateDisabled
			}
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
