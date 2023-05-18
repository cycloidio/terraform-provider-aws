package events_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/eventbridge"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/acctest"
	"github.com/hashicorp/terraform-provider-aws/conns"
	tfevents "github.com/hashicorp/terraform-provider-aws/service/events"
	"github.com/hashicorp/terraform-provider-aws/tfresource"
)

func TestAccEventsEndpoint_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v eventbridge.DescribeEndpointOutput
	var providers []*schema.Provider
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_cloudwatch_event_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, eventbridge.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
		CheckDestroy:             testAccCheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccEventsEndpoint_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v eventbridge.DescribeEndpointOutput
	var providers []*schema.Provider
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_cloudwatch_event_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, eventbridge.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
		CheckDestroy:             testAccCheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfevents.ResourceEndpoint(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccEventsEndpoint_roleARN(t *testing.T) {
	ctx := acctest.Context(t)
	var v eventbridge.DescribeEndpointOutput
	var providers []*schema.Provider
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_cloudwatch_event_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, eventbridge.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
		CheckDestroy:             testAccCheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_roleARN(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "ENABLED"),
					resource.TestCheckResourceAttrPair(resourceName, "role_arn", "aws_iam_role.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccEventsEndpoint_description(t *testing.T) {
	ctx := acctest.Context(t)
	var v eventbridge.DescribeEndpointOutput
	var providers []*schema.Provider
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_cloudwatch_event_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, eventbridge.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
		CheckDestroy:             testAccCheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_description(rName, "description 1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", "description 1"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccEndpointConfig_description(rName, "description 2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", "description 2"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
		},
	})
}

func TestAccEventsEndpoint_updateRoutingConfig(t *testing.T) {
	ctx := acctest.Context(t)
	var v eventbridge.DescribeEndpointOutput
	var providers []*schema.Provider
	resourceName := "aws_cloudwatch_event_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, eventbridge.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
		CheckDestroy:             testAccCheckEndpointDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
			{
				Config: testAccEndpointConfig_updateRoutingConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointExists(ctx, resourceName, &v),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "events", fmt.Sprintf("endpoint/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_url"),
					resource.TestCheckResourceAttr(resourceName, "event_bus.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.0.event_bus_arn", "aws_cloudwatch_event_bus.primary", "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "event_bus.1.event_bus_arn", "aws_cloudwatch_event_bus.secondary", "arn"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "replication_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_config.0.state", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
					resource.TestCheckResourceAttr(resourceName, "routing_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.primary.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "routing_config.0.failover_config.0.primary.0.health_check", "aws_route53_health_check.test2", "arn"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "routing_config.0.failover_config.0.secondary.0.route", acctest.AlternateRegion()),
				),
			},
		},
	})
}

func testAccCheckEndpointDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EventsConn()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_cloudwatch_event_endpoint" {
				continue
			}

			_, err := tfevents.FindEndpointByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("EventBridge Global Endpoint %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckEndpointExists(ctx context.Context, n string, v *eventbridge.DescribeEndpointOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EventsConn()

		output, err := tfevents.FindEndpointByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccEndpointConfig_base(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateRegionProvider(), fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_cloudwatch_event_bus" "primary" {
  name = %[1]q
}

resource "aws_cloudwatch_event_bus" "secondary" {
  provider = "awsalternate"

  name = %[1]q
}

data "aws_iam_policy_document" "test_assume" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["events.${data.aws_partition.current.dns_suffix}"]
    }
  }
}

resource "aws_iam_role" "test" {
  name               = %[1]q
  assume_role_policy = data.aws_iam_policy_document.test_assume.json

  inline_policy {
    name   = %[1]q
    policy = data.aws_iam_policy_document.test.json
  }
}

data "aws_iam_policy_document" "test" {
  statement {
    actions = [
      "events:PutRule",
      "events:PutTargets",
      "events:DeleteRule",
      "events:RemoveTargets",
    ]

    resources = ["arn:${data.aws_partition.current.partition}:events:*:*:rule/%[1]s/GlobalEndpointManagedRule-*"]
  }

  statement {
    actions = ["events:PutEvents"]

    resources = [
      aws_cloudwatch_event_bus.primary.arn,
      aws_cloudwatch_event_bus.secondary.arn,
    ]
  }

  statement {
    actions = ["iam:PassRole"]

    resources = ["arn:${data.aws_partition.current.partition}:iam::*:role/%[1]s"]

    condition {
      test     = "StringLike"
      variable = "iam:PassedToService"

      values = ["events.${data.aws_partition.current.dns_suffix}"]
    }
  }
}

resource "aws_route53_health_check" "test" {
  fqdn             = "example.com"
  type             = "HTTP"
  request_interval = "30"
  disabled         = true
  port             = 80

  tags = {
    Name = %[1]q
  }
}
`, rName, acctest.AlternateRegion()))
}

func testAccEndpointConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_base(rName), fmt.Sprintf(`
resource "aws_cloudwatch_event_endpoint" "test" {
  name = %[1]q

  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.primary.arn
  }
  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.secondary.arn
  }

  replication_config {
    state = "DISABLED"
  }

  routing_config {
    failover_config {
      primary {
        health_check = aws_route53_health_check.test.arn
      }

      secondary {
        route = %[2]q
      }
    }
  }
}
`, rName, acctest.AlternateRegion()))
}

func testAccEndpointConfig_roleARN(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_base(rName), fmt.Sprintf(`
resource "aws_cloudwatch_event_endpoint" "test" {
  name = %[1]q

  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.primary.arn
  }
  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.secondary.arn
  }

  role_arn = aws_iam_role.test.arn

  routing_config {
    failover_config {
      primary {
        health_check = aws_route53_health_check.test.arn
      }

      secondary {
        route = %[2]q
      }
    }
  }
}
`, rName, acctest.AlternateRegion()))
}

func testAccEndpointConfig_description(rName, description string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_base(rName), fmt.Sprintf(`
resource "aws_cloudwatch_event_endpoint" "test" {
  name        = %[1]q
  description = %[2]q

  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.primary.arn
  }
  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.secondary.arn
  }

  replication_config {
    state = "DISABLED"
  }

  routing_config {
    failover_config {
      primary {
        health_check = aws_route53_health_check.test.arn
      }

      secondary {
        route = %[3]q
      }
    }
  }
}
`, rName, description, acctest.AlternateRegion()))
}

func testAccEndpointConfig_updateRoutingConfig(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_base(rName), fmt.Sprintf(`
resource "aws_route53_health_check" "test2" {
  fqdn             = "example.com"
  type             = "HTTPS"
  request_interval = "10"
  disabled         = true
  port             = 443

  tags = {
    Name = %[1]q
  }
}

resource "aws_cloudwatch_event_endpoint" "test" {
  name = %[1]q

  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.primary.arn
  }
  event_bus {
    event_bus_arn = aws_cloudwatch_event_bus.secondary.arn
  }

  replication_config {
    state = "DISABLED"
  }

  routing_config {
    failover_config {
      primary {
        health_check = aws_route53_health_check.test2.arn
      }

      secondary {
        route = %[2]q
      }
    }
  }
}
`, rName, acctest.AlternateRegion()))
}
