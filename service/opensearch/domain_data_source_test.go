package opensearch_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/opensearchservice"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/acctest"
)

func TestAccOpenSearchDomainDataSource_Data_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	autoTuneStartAtTime := testAccGetValidStartAtTime(t, "24h")
	datasourceName := "data.aws_opensearch_domain.test"
	resourceName := "aws_opensearch_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t); testAccPreCheckIAMServiceLinkedRoleOpenSearch(t) },
		ErrorCheck: acctest.ErrorCheck(t, opensearchservice.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainWithDataSourceConfig(rName, autoTuneStartAtTime),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "processing", "false"),
					resource.TestCheckResourceAttrPair(datasourceName, "engine_version", resourceName, "engine_version"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.#", resourceName, "auto_tune_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.desired_state", resourceName, "auto_tune_options.0.desired_state"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.maintenance_schedule", resourceName, "auto_tune_options.0.maintenance_schedule"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.rollback_on_disable", resourceName, "auto_tune_options.0.rollback_on_disable"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.#", resourceName, "cluster_config.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.instance_type", resourceName, "cluster_config.0.instance_type"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.instance_count", resourceName, "cluster_config.0.instance_count"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.dedicated_master_enabled", resourceName, "cluster_config.0.dedicated_master_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.zone_awareness_enabled", resourceName, "cluster_config.0.zone_awareness_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.#", resourceName, "ebs_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.ebs_enabled", resourceName, "ebs_options.0.ebs_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.volume_type", resourceName, "ebs_options.0.volume_type"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.volume_size", resourceName, "ebs_options.0.volume_size"),
					resource.TestCheckResourceAttrPair(datasourceName, "snapshot_options.#", resourceName, "snapshot_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "snapshot_options.0.automated_snapshot_start_hour", resourceName, "snapshot_options.0.automated_snapshot_start_hour"),
					resource.TestCheckResourceAttrPair(datasourceName, "advanced_security_options.#", resourceName, "advanced_security_options.#"),
				),
			},
		},
	})
}

func TestAccOpenSearchDomainDataSource_Data_advanced(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	autoTuneStartAtTime := testAccGetValidStartAtTime(t, "24h")
	datasourceName := "data.aws_opensearch_domain.test"
	resourceName := "aws_opensearch_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t); testAccPreCheckIAMServiceLinkedRoleOpenSearch(t) },
		ErrorCheck: acctest.ErrorCheck(t, opensearchservice.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainAdvancedWithDataSourceConfig(rName, autoTuneStartAtTime),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "engine_version", resourceName, "engine_version"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.#", resourceName, "auto_tune_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.desired_state", resourceName, "auto_tune_options.0.desired_state"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.maintenance_schedule", resourceName, "auto_tune_options.0.maintenance_schedule"),
					resource.TestCheckResourceAttrPair(datasourceName, "auto_tune_options.0.rollback_on_disable", resourceName, "auto_tune_options.0.rollback_on_disable"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.#", resourceName, "cluster_config.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.instance_type", resourceName, "cluster_config.0.instance_type"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.instance_count", resourceName, "cluster_config.0.instance_count"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.dedicated_master_enabled", resourceName, "cluster_config.0.dedicated_master_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "cluster_config.0.zone_awareness_enabled", resourceName, "cluster_config.0.zone_awareness_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.#", resourceName, "ebs_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.ebs_enabled", resourceName, "ebs_options.0.ebs_enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.volume_type", resourceName, "ebs_options.0.volume_type"),
					resource.TestCheckResourceAttrPair(datasourceName, "ebs_options.0.volume_size", resourceName, "ebs_options.0.volume_size"),
					resource.TestCheckResourceAttrPair(datasourceName, "snapshot_options.#", resourceName, "snapshot_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "snapshot_options.0.automated_snapshot_start_hour", resourceName, "snapshot_options.0.automated_snapshot_start_hour"),
					resource.TestCheckResourceAttrPair(datasourceName, "log_publishing_options.#", resourceName, "log_publishing_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "vpc_options.#", resourceName, "vpc_options.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "advanced_security_options.0.enabled", resourceName, "advanced_security_options.0.enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "advanced_security_options.0.internal_user_database_enabled", resourceName, "advanced_security_options.0.internal_user_database_enabled"),
				),
			},
		},
	})
}

func testAccDomainWithDataSourceConfig(rName, autoTuneStartAtTime string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

locals {
  domain_substr = substr(%[1]q, 0, 28)
}

resource "aws_opensearch_domain" "test" {
  domain_name    = local.domain_substr
  engine_version = "Elasticsearch_7.10"

  access_policies = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:${data.aws_partition.current.partition}:es:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:domain/${local.domain_substr}/*",
      "Condition": {
        "IpAddress": {
          "aws:SourceIp": [
            "127.0.0.0/8"
          ]
        }
      }
    }
  ]
}
POLICY

  auto_tune_options {
    desired_state = "ENABLED"

    maintenance_schedule {
      start_at = %[2]q
      duration {
        value = "2"
        unit  = "HOURS"
      }
      cron_expression_for_recurrence = "cron(0 0 ? * 1 *)"
    }

    rollback_on_disable = "NO_ROLLBACK"

  }

  cluster_config {
    instance_type            = "t2.small.search"
    instance_count           = 2
    dedicated_master_enabled = false

    zone_awareness_config {
      availability_zone_count = 2
    }

    zone_awareness_enabled = true
  }

  ebs_options {
    ebs_enabled = true
    volume_type = "gp2"
    volume_size = 20
  }

  snapshot_options {
    automated_snapshot_start_hour = 23
  }
}

data "aws_opensearch_domain" "test" {
  domain_name = aws_opensearch_domain.test.domain_name
}
`, rName, autoTuneStartAtTime)
}

func testAccDomainAdvancedWithDataSourceConfig(rName, autoTuneStartAtTime string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

resource "aws_cloudwatch_log_group" "test" {
  name = %[1]q
}

resource "aws_cloudwatch_log_resource_policy" "test" {
  policy_name = %[1]q

  policy_document = <<CONFIG
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "opensearchservice.${data.aws_partition.current.dns_suffix}"
      },
      "Action": [
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
        "logs:CreateLogStream"
      ],
      "Resource": "arn:${data.aws_partition.current.partition}:logs:*"
    }
  ]
}
CONFIG
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = "10.0.0.0/24"
  vpc_id            = aws_vpc.test.id
}

resource "aws_subnet" "test2" {
  availability_zone = data.aws_availability_zones.available.names[1]
  cidr_block        = "10.0.1.0/24"
  vpc_id            = aws_vpc.test.id
}

resource "aws_security_group" "test" {
  name   = %[1]q
  vpc_id = aws_vpc.test.id
}

resource "aws_security_group_rule" "test" {
  type        = "ingress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.test.id
}

locals {
  domain_substr = substr(%[1]q, 0, 28)
}

resource "aws_opensearch_domain" "test" {
  domain_name    = local.domain_substr
  engine_version = "Elasticsearch_6.7"

  access_policies = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:${data.aws_partition.current.partition}:es:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:domain/${local.domain_substr}/*"
    }
  ]
}
POLICY

  auto_tune_options {
    desired_state = "ENABLED"

    maintenance_schedule {
      start_at = %[2]q
      duration {
        value = "2"
        unit  = "HOURS"
      }
      cron_expression_for_recurrence = "cron(0 0 ? * 1 *)"
    }

    rollback_on_disable = "NO_ROLLBACK"

  }

  cluster_config {
    instance_type            = "t2.small.search"
    instance_count           = 2
    dedicated_master_enabled = false

    zone_awareness_config {
      availability_zone_count = 2
    }

    zone_awareness_enabled = true
  }

  ebs_options {
    ebs_enabled = true
    volume_type = "gp2"
    volume_size = 20
  }

  snapshot_options {
    automated_snapshot_start_hour = 23
  }

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.test.arn
    log_type                 = "INDEX_SLOW_LOGS"
  }

  vpc_options {
    security_group_ids = [
      aws_security_group.test.id
    ]
    subnet_ids = [
      aws_subnet.test.id,
      aws_subnet.test2.id
    ]
  }

  advanced_security_options {
    enabled                        = false
    internal_user_database_enabled = false
  }

  tags = {
    Domain = %[1]q
  }
}

data "aws_opensearch_domain" "test" {
  domain_name = aws_opensearch_domain.test.domain_name
}
`, rName, autoTuneStartAtTime))
}
