// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kendra

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
			Factory:  DataSourceExperience,
			TypeName: "aws_kendra_experience",
		},
		{
			Factory:  DataSourceFaq,
			TypeName: "aws_kendra_faq",
		},
		{
			Factory:  DataSourceIndex,
			TypeName: "aws_kendra_index",
		},
		{
			Factory:  DataSourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
		},
		{
			Factory:  DataSourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_kendra_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceExperience,
			TypeName: "aws_kendra_experience",
		},
		{
			Factory:  ResourceFaq,
			TypeName: "aws_kendra_faq",
			Name:     "FAQ",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceIndex,
			TypeName: "aws_kendra_index",
			Name:     "Index",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
			Name:     "Query Suggestions Block List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
			Name:     "Thesaurus",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kendra
}

var ServicePackage = &servicePackage{}
