package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
)

func Provider() *schema.Provider {
	return provider.Provider()
}

func Config(accessKey, secretKey, region, sessionToken string) conns.Config {
	return conns.Config{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
		Token:     sessionToken,
	}
}
