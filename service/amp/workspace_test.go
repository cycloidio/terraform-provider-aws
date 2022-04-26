package amp_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/acctest"
	"github.com/hashicorp/terraform-provider-aws/conns"
	tfamp "github.com/hashicorp/terraform-provider-aws/service/amp"
)

func TestAccAMPWorkspace_basic(t *testing.T) {
	workspaceAlias := sdkacctest.RandomWithPrefix("tf_amp_workspace")
	resourceName := "aws_prometheus_workspace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(prometheusservice.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, prometheusservice.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAMPWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAMPWorkspaceWithAliasConfig(workspaceAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "alias", workspaceAlias),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAMPWorkspaceWithoutAliasConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "alias", ""),
				),
			},
			{
				Config: testAccAMPWorkspaceWithAliasConfig(workspaceAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "alias", workspaceAlias),
				),
			},
		},
	})
}

func TestAccAMPWorkspace_disappears(t *testing.T) {
	resourceName := "aws_prometheus_workspace.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(prometheusservice.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, prometheusservice.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAMPWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAMPWorkspaceWithoutAliasConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					acctest.CheckResourceDisappears(acctest.Provider, tfamp.ResourceWorkspace(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAMPWorkspace_tags(t *testing.T) {
	rInt := sdkacctest.RandInt()
	resourceName := "aws_prometheus_workspace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, cloudwatchlogs.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAMPWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAMPWorkspaceWithTagsConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Environment", "production"),
					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "mh9"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAMPWorkspaceWithTagsAddedConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Environment", "production"),
					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "mh9"),
					resource.TestCheckResourceAttr(resourceName, "tags.App", "foo42"),
				),
			},
			{
				Config: testAccAMPWorkspaceWithTagsUpdatedConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.Environment", "production"),
					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "abhi"),
					resource.TestCheckResourceAttr(resourceName, "tags.App", "foo42"),
				),
			},
			{
				Config: testAccAMPWorkspaceWithTagsConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAMPWorkspaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Environment", "production"),
					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "mh9"),
				),
			},
		},
	})
}

func testAccCheckAMPWorkspaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AMP Workspace ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AMPConn

		req := &prometheusservice.DescribeWorkspaceInput{
			WorkspaceId: aws.String(rs.Primary.ID),
		}
		describe, err := conn.DescribeWorkspace(req)
		if err != nil {
			return err
		}
		if describe == nil {
			return fmt.Errorf("Got nil account ?!")
		}

		return nil
	}
}

func testAccCheckAMPWorkspaceDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).AMPConn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_prometheus_workspace" {
			continue
		}

		_, err := conn.DescribeWorkspace(&prometheusservice.DescribeWorkspaceInput{
			WorkspaceId: aws.String(rs.Primary.ID),
		})
		if tfawserr.ErrCodeEquals(err, prometheusservice.ErrCodeResourceNotFoundException) {
			continue
		}

		if err != nil {
			return fmt.Errorf("error reading Prometheus WorkSpace (%s): %w", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccAMPWorkspaceWithAliasConfig(randName string) string {
	return fmt.Sprintf(`
resource "aws_prometheus_workspace" "test" {
  alias = %q
}
`, randName)
}

func testAccAMPWorkspaceWithoutAliasConfig() string {
	return `
resource "aws_prometheus_workspace" "test" {
}
`
}

func testAccAMPWorkspaceWithTagsConfig(randInt int) string {
	return fmt.Sprintf(`
resource "aws_prometheus_workspace" "test" {
  alias = "test-%d"

  tags = {
    Environment = "production"
    Owner       = "mh9"
  }
}
`, randInt)
}

func testAccAMPWorkspaceWithTagsAddedConfig(randInt int) string {
	return fmt.Sprintf(`
resource "aws_prometheus_workspace" "test" {
  alias = "test-%d"

  tags = {
    Environment = "production"
    Owner       = "mh9"
    App         = "foo42"
  }
}
`, randInt)
}

func testAccAMPWorkspaceWithTagsUpdatedConfig(randInt int) string {
	return fmt.Sprintf(`
resource "aws_prometheus_workspace" "test" {
  alias = "test-%d"

  tags = {
    Environment = "production"
    Owner       = "abhi"
    App         = "foo42"
  }
}
`, randInt)
}
