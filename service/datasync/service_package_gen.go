// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package datasync

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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAgent,
			TypeName: "aws_datasync_agent",
			Name:     "Agent",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationEFS,
			TypeName: "aws_datasync_location_efs",
			Name:     "Location EFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationFSxLustreFileSystem,
			TypeName: "aws_datasync_location_fsx_lustre_file_system",
			Name:     "Location FSx Lustre File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationFSxOpenZFSFileSystem,
			TypeName: "aws_datasync_location_fsx_openzfs_file_system",
			Name:     "Location OpenZFS File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationFSxWindowsFileSystem,
			TypeName: "aws_datasync_location_fsx_windows_file_system",
			Name:     "Location FSx Windows File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationHDFS,
			TypeName: "aws_datasync_location_hdfs",
			Name:     "Location HDFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationNFS,
			TypeName: "aws_datasync_location_nfs",
			Name:     "Location NFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationObjectStorage,
			TypeName: "aws_datasync_location_object_storage",
			Name:     "Location Object Storage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationS3,
			TypeName: "aws_datasync_location_s3",
			Name:     "Location S3",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceLocationSMB,
			TypeName: "aws_datasync_location_smb",
			Name:     "Location SMB",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceTask,
			TypeName: "aws_datasync_task",
			Name:     "Task",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataSync
}

var ServicePackage = &servicePackage{}
