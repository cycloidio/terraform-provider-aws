// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iam

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	awstypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKDataSource("aws_iam_openid_connect_provider", name="OIDC Provider")
func dataSourceOpenIDConnectProvider() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceOpenIDConnectProviderRead,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: verify.ValidARN,
				ExactlyOneOf: []string{"arn", "url"},
			},
			"client_id_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"thumbprint_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": tftags.TagsSchemaComputed(),
			"url": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validOpenIDURL,
				DiffSuppressFunc: suppressOpenIDURL,
				ExactlyOneOf:     []string{"arn", "url"},
			},
		},
	}
}

func dataSourceOpenIDConnectProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).IAMClient(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &iam.GetOpenIDConnectProviderInput{}

	if v, ok := d.GetOk("arn"); ok {
		input.OpenIDConnectProviderArn = aws.String(v.(string))
	} else if v, ok := d.GetOk("url"); ok {
		url := v.(string)

		oidcpEntry, err := dataSourceGetOpenIDConnectProviderByURL(ctx, conn, url)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "finding IAM OIDC Provider by url (%s): %s", url, err)
		}

		if oidcpEntry == nil {
			return sdkdiag.AppendErrorf(diags, "finding IAM OIDC Provider by url (%s): not found", url)
		}
		input.OpenIDConnectProviderArn = oidcpEntry.Arn
	}

	resp, err := conn.GetOpenIDConnectProvider(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IAM OIDC Provider: %s", err)
	}

	d.SetId(aws.ToString(input.OpenIDConnectProviderArn))
	d.Set("arn", input.OpenIDConnectProviderArn)
	d.Set("url", resp.Url)
	d.Set("client_id_list", flex.FlattenStringValueList(resp.ClientIDList))
	d.Set("thumbprint_list", flex.FlattenStringValueList(resp.ThumbprintList))

	if err := d.Set("tags", KeyValueTags(ctx, resp.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	return diags
}

func dataSourceGetOpenIDConnectProviderByURL(ctx context.Context, conn *iam.Client, url string) (*awstypes.OpenIDConnectProviderListEntry, error) {
	var result *awstypes.OpenIDConnectProviderListEntry

	input := &iam.ListOpenIDConnectProvidersInput{}

	output, err := conn.ListOpenIDConnectProviders(ctx, input)

	if err != nil {
		return nil, err
	}

	for _, oidcp := range output.OpenIDConnectProviderList {
		if reflect.ValueOf(oidcp).IsZero() {
			continue
		}

		arnUrl, err := urlFromOpenIDConnectProviderARN(aws.ToString(oidcp.Arn))
		if err != nil {
			return nil, err
		}

		if arnUrl == strings.TrimPrefix(url, "https://") {
			return &oidcp, nil
		}
	}

	return result, nil
}

func urlFromOpenIDConnectProviderARN(arn string) (string, error) {
	parts := strings.SplitN(arn, "/", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("reading OpenID Connect Provider expected the arn to be like: arn:PARTITION:iam::ACCOUNT:oidc-provider/URL but got: %s", arn)
	}
	return parts[1], nil
}
