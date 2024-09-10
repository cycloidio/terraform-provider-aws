// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package dynamodb

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
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourcePolicyResource,
			Name:    "Resource Policy",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceTable,
			TypeName: "aws_dynamodb_table",
			Name:     "Table",
		},
		{
			Factory:  dataSourceTableItem,
			TypeName: "aws_dynamodb_table_item",
			Name:     "Table Item",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceContributorInsights,
			TypeName: "aws_dynamodb_contributor_insights",
			Name:     "Contributor Insights",
		},
		{
			Factory:  resourceGlobalTable,
			TypeName: "aws_dynamodb_global_table",
			Name:     "Global Table",
		},
		{
			Factory:  resourceKinesisStreamingDestination,
			TypeName: "aws_dynamodb_kinesis_streaming_destination",
			Name:     "Kinesis Streaming Destination",
		},
		{
			Factory:  resourceTable,
			TypeName: "aws_dynamodb_table",
			Name:     "Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTableExport,
			TypeName: "aws_dynamodb_table_export",
			Name:     "Table Export",
		},
		{
			Factory:  resourceTableItem,
			TypeName: "aws_dynamodb_table_item",
			Name:     "Table Item",
		},
		{
			Factory:  resourceTableReplica,
			TypeName: "aws_dynamodb_table_replica",
			Name:     "Table Replica",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceTag,
			TypeName: "aws_dynamodb_tag",
			Name:     "DynamoDB Resource Tag",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DynamoDB
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
