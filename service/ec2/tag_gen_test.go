// Code generated by internal/generate/tagresource/main.go; DO NOT EDIT.

package ec2_test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/acctest"
	"github.com/hashicorp/terraform-provider-aws/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/service/ec2"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
	"github.com/hashicorp/terraform-provider-aws/tfresource"
)

func testAccCheckTagDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_ec2_tag" {
			continue
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)

		if err != nil {
			return err
		}

		_, err = tfec2.GetTag(conn, identifier, key)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("%s resource (%s) tag (%s) still exists", ec2.ServiceID, identifier, key)
	}

	return nil
}

func testAccCheckTagExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s: missing resource ID", resourceName)
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)

		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn

		_, err = tfec2.GetTag(conn, identifier, key)

		if err != nil {
			return err
		}

		return nil
	}
}
