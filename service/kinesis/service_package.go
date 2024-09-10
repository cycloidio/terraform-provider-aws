// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/hashicorp/terraform-provider-aws/conns"
	"github.com/hashicorp/terraform-provider-aws/errs"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kinesis.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return kinesis.NewFromConfig(cfg,
		kinesis.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *kinesis.Options) {
			o.Retryer = conns.AddIsErrorRetryables(cfg.Retryer().(aws.RetryerV2), retry.IsErrorRetryableFunc(func(err error) aws.Ternary {
				if errs.IsAErrorMessageContains[*types.LimitExceededException](err, "simultaneously be in CREATING or DELETING") ||
					errs.IsAErrorMessageContains[*types.LimitExceededException](err, "Rate exceeded for stream") {
					return aws.TrueTernary
				}
				return aws.UnknownTernary // Delegate to configured Retryer.
			}))
		},
	), nil
}
