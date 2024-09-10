// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cur

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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceReportDefinition,
			TypeName: "aws_cur_report_definition",
			Name:     "Report Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "report_name",
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceReportDefinition,
			TypeName: "aws_cur_report_definition",
			Name:     "Report Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "report_name",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CUR
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
