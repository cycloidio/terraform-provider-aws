// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sweep_test

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-provider-aws/conns"
	"github.com/hashicorp/terraform-provider-aws/service/accessanalyzer"
	"github.com/hashicorp/terraform-provider-aws/service/account"
	"github.com/hashicorp/terraform-provider-aws/service/acm"
	"github.com/hashicorp/terraform-provider-aws/service/acmpca"
	"github.com/hashicorp/terraform-provider-aws/service/amp"
	"github.com/hashicorp/terraform-provider-aws/service/amplify"
	"github.com/hashicorp/terraform-provider-aws/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/service/apigatewayv2"
	"github.com/hashicorp/terraform-provider-aws/service/appautoscaling"
	"github.com/hashicorp/terraform-provider-aws/service/appconfig"
	"github.com/hashicorp/terraform-provider-aws/service/appfabric"
	"github.com/hashicorp/terraform-provider-aws/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/service/appintegrations"
	"github.com/hashicorp/terraform-provider-aws/service/applicationinsights"
	"github.com/hashicorp/terraform-provider-aws/service/applicationsignals"
	"github.com/hashicorp/terraform-provider-aws/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/service/appsync"
	"github.com/hashicorp/terraform-provider-aws/service/athena"
	"github.com/hashicorp/terraform-provider-aws/service/auditmanager"
	"github.com/hashicorp/terraform-provider-aws/service/autoscaling"
	"github.com/hashicorp/terraform-provider-aws/service/autoscalingplans"
	"github.com/hashicorp/terraform-provider-aws/service/backup"
	"github.com/hashicorp/terraform-provider-aws/service/batch"
	"github.com/hashicorp/terraform-provider-aws/service/bcmdataexports"
	"github.com/hashicorp/terraform-provider-aws/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/service/bedrockagent"
	"github.com/hashicorp/terraform-provider-aws/service/budgets"
	"github.com/hashicorp/terraform-provider-aws/service/ce"
	"github.com/hashicorp/terraform-provider-aws/service/chatbot"
	"github.com/hashicorp/terraform-provider-aws/service/chime"
	"github.com/hashicorp/terraform-provider-aws/service/chimesdkmediapipelines"
	"github.com/hashicorp/terraform-provider-aws/service/chimesdkvoice"
	"github.com/hashicorp/terraform-provider-aws/service/cleanrooms"
	"github.com/hashicorp/terraform-provider-aws/service/cloud9"
	"github.com/hashicorp/terraform-provider-aws/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-aws/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws/service/cloudfront"
	"github.com/hashicorp/terraform-provider-aws/service/cloudfrontkeyvaluestore"
	"github.com/hashicorp/terraform-provider-aws/service/cloudhsmv2"
	"github.com/hashicorp/terraform-provider-aws/service/cloudsearch"
	"github.com/hashicorp/terraform-provider-aws/service/cloudtrail"
	"github.com/hashicorp/terraform-provider-aws/service/cloudwatch"
	"github.com/hashicorp/terraform-provider-aws/service/codeartifact"
	"github.com/hashicorp/terraform-provider-aws/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/service/codecatalyst"
	"github.com/hashicorp/terraform-provider-aws/service/codecommit"
	"github.com/hashicorp/terraform-provider-aws/service/codeguruprofiler"
	"github.com/hashicorp/terraform-provider-aws/service/codegurureviewer"
	"github.com/hashicorp/terraform-provider-aws/service/codepipeline"
	"github.com/hashicorp/terraform-provider-aws/service/codestarconnections"
	"github.com/hashicorp/terraform-provider-aws/service/codestarnotifications"
	"github.com/hashicorp/terraform-provider-aws/service/cognitoidentity"
	"github.com/hashicorp/terraform-provider-aws/service/cognitoidp"
	"github.com/hashicorp/terraform-provider-aws/service/comprehend"
	"github.com/hashicorp/terraform-provider-aws/service/computeoptimizer"
	"github.com/hashicorp/terraform-provider-aws/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/service/connect"
	"github.com/hashicorp/terraform-provider-aws/service/connectcases"
	"github.com/hashicorp/terraform-provider-aws/service/controltower"
	"github.com/hashicorp/terraform-provider-aws/service/costoptimizationhub"
	"github.com/hashicorp/terraform-provider-aws/service/cur"
	"github.com/hashicorp/terraform-provider-aws/service/customerprofiles"
	"github.com/hashicorp/terraform-provider-aws/service/databrew"
	"github.com/hashicorp/terraform-provider-aws/service/dataexchange"
	"github.com/hashicorp/terraform-provider-aws/service/datapipeline"
	"github.com/hashicorp/terraform-provider-aws/service/datasync"
	"github.com/hashicorp/terraform-provider-aws/service/datazone"
	"github.com/hashicorp/terraform-provider-aws/service/dax"
	"github.com/hashicorp/terraform-provider-aws/service/deploy"
	"github.com/hashicorp/terraform-provider-aws/service/detective"
	"github.com/hashicorp/terraform-provider-aws/service/devicefarm"
	"github.com/hashicorp/terraform-provider-aws/service/devopsguru"
	"github.com/hashicorp/terraform-provider-aws/service/directconnect"
	"github.com/hashicorp/terraform-provider-aws/service/dlm"
	"github.com/hashicorp/terraform-provider-aws/service/dms"
	"github.com/hashicorp/terraform-provider-aws/service/docdb"
	"github.com/hashicorp/terraform-provider-aws/service/docdbelastic"
	"github.com/hashicorp/terraform-provider-aws/service/drs"
	"github.com/hashicorp/terraform-provider-aws/service/ds"
	"github.com/hashicorp/terraform-provider-aws/service/dynamodb"
	"github.com/hashicorp/terraform-provider-aws/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/service/ecrpublic"
	"github.com/hashicorp/terraform-provider-aws/service/ecs"
	"github.com/hashicorp/terraform-provider-aws/service/efs"
	"github.com/hashicorp/terraform-provider-aws/service/eks"
	"github.com/hashicorp/terraform-provider-aws/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-provider-aws/service/elasticsearch"
	"github.com/hashicorp/terraform-provider-aws/service/elastictranscoder"
	"github.com/hashicorp/terraform-provider-aws/service/elb"
	"github.com/hashicorp/terraform-provider-aws/service/elbv2"
	"github.com/hashicorp/terraform-provider-aws/service/emr"
	"github.com/hashicorp/terraform-provider-aws/service/emrcontainers"
	"github.com/hashicorp/terraform-provider-aws/service/emrserverless"
	"github.com/hashicorp/terraform-provider-aws/service/events"
	"github.com/hashicorp/terraform-provider-aws/service/evidently"
	"github.com/hashicorp/terraform-provider-aws/service/finspace"
	"github.com/hashicorp/terraform-provider-aws/service/firehose"
	"github.com/hashicorp/terraform-provider-aws/service/fis"
	"github.com/hashicorp/terraform-provider-aws/service/fms"
	"github.com/hashicorp/terraform-provider-aws/service/fsx"
	"github.com/hashicorp/terraform-provider-aws/service/gamelift"
	"github.com/hashicorp/terraform-provider-aws/service/glacier"
	"github.com/hashicorp/terraform-provider-aws/service/globalaccelerator"
	"github.com/hashicorp/terraform-provider-aws/service/glue"
	"github.com/hashicorp/terraform-provider-aws/service/grafana"
	"github.com/hashicorp/terraform-provider-aws/service/greengrass"
	"github.com/hashicorp/terraform-provider-aws/service/groundstation"
	"github.com/hashicorp/terraform-provider-aws/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/service/healthlake"
	"github.com/hashicorp/terraform-provider-aws/service/iam"
	"github.com/hashicorp/terraform-provider-aws/service/identitystore"
	"github.com/hashicorp/terraform-provider-aws/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/service/inspector"
	"github.com/hashicorp/terraform-provider-aws/service/inspector2"
	"github.com/hashicorp/terraform-provider-aws/service/internetmonitor"
	"github.com/hashicorp/terraform-provider-aws/service/iot"
	"github.com/hashicorp/terraform-provider-aws/service/iotanalytics"
	"github.com/hashicorp/terraform-provider-aws/service/iotevents"
	"github.com/hashicorp/terraform-provider-aws/service/ivs"
	"github.com/hashicorp/terraform-provider-aws/service/ivschat"
	"github.com/hashicorp/terraform-provider-aws/service/kafka"
	"github.com/hashicorp/terraform-provider-aws/service/kafkaconnect"
	"github.com/hashicorp/terraform-provider-aws/service/kendra"
	"github.com/hashicorp/terraform-provider-aws/service/keyspaces"
	"github.com/hashicorp/terraform-provider-aws/service/kinesis"
	"github.com/hashicorp/terraform-provider-aws/service/kinesisanalytics"
	"github.com/hashicorp/terraform-provider-aws/service/kinesisanalyticsv2"
	"github.com/hashicorp/terraform-provider-aws/service/kinesisvideo"
	"github.com/hashicorp/terraform-provider-aws/service/kms"
	"github.com/hashicorp/terraform-provider-aws/service/lakeformation"
	"github.com/hashicorp/terraform-provider-aws/service/lambda"
	"github.com/hashicorp/terraform-provider-aws/service/launchwizard"
	"github.com/hashicorp/terraform-provider-aws/service/lexmodels"
	"github.com/hashicorp/terraform-provider-aws/service/lexv2models"
	"github.com/hashicorp/terraform-provider-aws/service/licensemanager"
	"github.com/hashicorp/terraform-provider-aws/service/lightsail"
	"github.com/hashicorp/terraform-provider-aws/service/location"
	"github.com/hashicorp/terraform-provider-aws/service/logs"
	"github.com/hashicorp/terraform-provider-aws/service/lookoutmetrics"
	"github.com/hashicorp/terraform-provider-aws/service/m2"
	"github.com/hashicorp/terraform-provider-aws/service/macie2"
	"github.com/hashicorp/terraform-provider-aws/service/mediaconnect"
	"github.com/hashicorp/terraform-provider-aws/service/mediaconvert"
	"github.com/hashicorp/terraform-provider-aws/service/medialive"
	"github.com/hashicorp/terraform-provider-aws/service/mediapackage"
	"github.com/hashicorp/terraform-provider-aws/service/mediapackagev2"
	"github.com/hashicorp/terraform-provider-aws/service/mediastore"
	"github.com/hashicorp/terraform-provider-aws/service/memorydb"
	"github.com/hashicorp/terraform-provider-aws/service/meta"
	"github.com/hashicorp/terraform-provider-aws/service/mq"
	"github.com/hashicorp/terraform-provider-aws/service/mwaa"
	"github.com/hashicorp/terraform-provider-aws/service/neptune"
	"github.com/hashicorp/terraform-provider-aws/service/neptunegraph"
	"github.com/hashicorp/terraform-provider-aws/service/networkfirewall"
	"github.com/hashicorp/terraform-provider-aws/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/service/networkmonitor"
	"github.com/hashicorp/terraform-provider-aws/service/oam"
	"github.com/hashicorp/terraform-provider-aws/service/opensearch"
	"github.com/hashicorp/terraform-provider-aws/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/service/opsworks"
	"github.com/hashicorp/terraform-provider-aws/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/service/osis"
	"github.com/hashicorp/terraform-provider-aws/service/outposts"
	"github.com/hashicorp/terraform-provider-aws/service/paymentcryptography"
	"github.com/hashicorp/terraform-provider-aws/service/pcaconnectorad"
	"github.com/hashicorp/terraform-provider-aws/service/pcs"
	"github.com/hashicorp/terraform-provider-aws/service/pinpoint"
	"github.com/hashicorp/terraform-provider-aws/service/pipes"
	"github.com/hashicorp/terraform-provider-aws/service/polly"
	"github.com/hashicorp/terraform-provider-aws/service/pricing"
	"github.com/hashicorp/terraform-provider-aws/service/qbusiness"
	"github.com/hashicorp/terraform-provider-aws/service/qldb"
	"github.com/hashicorp/terraform-provider-aws/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/service/ram"
	"github.com/hashicorp/terraform-provider-aws/service/rbin"
	"github.com/hashicorp/terraform-provider-aws/service/rds"
	"github.com/hashicorp/terraform-provider-aws/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/service/redshiftdata"
	"github.com/hashicorp/terraform-provider-aws/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/service/rekognition"
	"github.com/hashicorp/terraform-provider-aws/service/resiliencehub"
	"github.com/hashicorp/terraform-provider-aws/service/resourceexplorer2"
	"github.com/hashicorp/terraform-provider-aws/service/resourcegroups"
	"github.com/hashicorp/terraform-provider-aws/service/resourcegroupstaggingapi"
	"github.com/hashicorp/terraform-provider-aws/service/rolesanywhere"
	"github.com/hashicorp/terraform-provider-aws/service/route53"
	"github.com/hashicorp/terraform-provider-aws/service/route53domains"
	"github.com/hashicorp/terraform-provider-aws/service/route53profiles"
	"github.com/hashicorp/terraform-provider-aws/service/route53recoverycontrolconfig"
	"github.com/hashicorp/terraform-provider-aws/service/route53recoveryreadiness"
	"github.com/hashicorp/terraform-provider-aws/service/route53resolver"
	"github.com/hashicorp/terraform-provider-aws/service/rum"
	"github.com/hashicorp/terraform-provider-aws/service/s3"
	"github.com/hashicorp/terraform-provider-aws/service/s3control"
	"github.com/hashicorp/terraform-provider-aws/service/s3outposts"
	"github.com/hashicorp/terraform-provider-aws/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/service/scheduler"
	"github.com/hashicorp/terraform-provider-aws/service/schemas"
	"github.com/hashicorp/terraform-provider-aws/service/secretsmanager"
	"github.com/hashicorp/terraform-provider-aws/service/securityhub"
	"github.com/hashicorp/terraform-provider-aws/service/securitylake"
	"github.com/hashicorp/terraform-provider-aws/service/serverlessrepo"
	"github.com/hashicorp/terraform-provider-aws/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/service/servicecatalogappregistry"
	"github.com/hashicorp/terraform-provider-aws/service/servicediscovery"
	"github.com/hashicorp/terraform-provider-aws/service/servicequotas"
	"github.com/hashicorp/terraform-provider-aws/service/ses"
	"github.com/hashicorp/terraform-provider-aws/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/service/sfn"
	"github.com/hashicorp/terraform-provider-aws/service/shield"
	"github.com/hashicorp/terraform-provider-aws/service/signer"
	"github.com/hashicorp/terraform-provider-aws/service/simpledb"
	"github.com/hashicorp/terraform-provider-aws/service/sns"
	"github.com/hashicorp/terraform-provider-aws/service/sqs"
	"github.com/hashicorp/terraform-provider-aws/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/service/ssmcontacts"
	"github.com/hashicorp/terraform-provider-aws/service/ssmincidents"
	"github.com/hashicorp/terraform-provider-aws/service/ssmsap"
	"github.com/hashicorp/terraform-provider-aws/service/sso"
	"github.com/hashicorp/terraform-provider-aws/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/service/storagegateway"
	"github.com/hashicorp/terraform-provider-aws/service/sts"
	"github.com/hashicorp/terraform-provider-aws/service/swf"
	"github.com/hashicorp/terraform-provider-aws/service/synthetics"
	"github.com/hashicorp/terraform-provider-aws/service/timestreaminfluxdb"
	"github.com/hashicorp/terraform-provider-aws/service/timestreamwrite"
	"github.com/hashicorp/terraform-provider-aws/service/transcribe"
	"github.com/hashicorp/terraform-provider-aws/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/service/verifiedpermissions"
	"github.com/hashicorp/terraform-provider-aws/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/service/waf"
	"github.com/hashicorp/terraform-provider-aws/service/wafregional"
	"github.com/hashicorp/terraform-provider-aws/service/wafv2"
	"github.com/hashicorp/terraform-provider-aws/service/wellarchitected"
	"github.com/hashicorp/terraform-provider-aws/service/worklink"
	"github.com/hashicorp/terraform-provider-aws/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/service/workspacesweb"
	"github.com/hashicorp/terraform-provider-aws/service/xray"
)

func servicePackages(ctx context.Context) []conns.ServicePackage {
	v := []conns.ServicePackage{
		accessanalyzer.ServicePackage(ctx),
		account.ServicePackage(ctx),
		acm.ServicePackage(ctx),
		acmpca.ServicePackage(ctx),
		amp.ServicePackage(ctx),
		amplify.ServicePackage(ctx),
		apigateway.ServicePackage(ctx),
		apigatewayv2.ServicePackage(ctx),
		appautoscaling.ServicePackage(ctx),
		appconfig.ServicePackage(ctx),
		appfabric.ServicePackage(ctx),
		appflow.ServicePackage(ctx),
		appintegrations.ServicePackage(ctx),
		applicationinsights.ServicePackage(ctx),
		applicationsignals.ServicePackage(ctx),
		appmesh.ServicePackage(ctx),
		apprunner.ServicePackage(ctx),
		appstream.ServicePackage(ctx),
		appsync.ServicePackage(ctx),
		athena.ServicePackage(ctx),
		auditmanager.ServicePackage(ctx),
		autoscaling.ServicePackage(ctx),
		autoscalingplans.ServicePackage(ctx),
		backup.ServicePackage(ctx),
		batch.ServicePackage(ctx),
		bcmdataexports.ServicePackage(ctx),
		bedrock.ServicePackage(ctx),
		bedrockagent.ServicePackage(ctx),
		budgets.ServicePackage(ctx),
		ce.ServicePackage(ctx),
		chatbot.ServicePackage(ctx),
		chime.ServicePackage(ctx),
		chimesdkmediapipelines.ServicePackage(ctx),
		chimesdkvoice.ServicePackage(ctx),
		cleanrooms.ServicePackage(ctx),
		cloud9.ServicePackage(ctx),
		cloudcontrol.ServicePackage(ctx),
		cloudformation.ServicePackage(ctx),
		cloudfront.ServicePackage(ctx),
		cloudfrontkeyvaluestore.ServicePackage(ctx),
		cloudhsmv2.ServicePackage(ctx),
		cloudsearch.ServicePackage(ctx),
		cloudtrail.ServicePackage(ctx),
		cloudwatch.ServicePackage(ctx),
		codeartifact.ServicePackage(ctx),
		codebuild.ServicePackage(ctx),
		codecatalyst.ServicePackage(ctx),
		codecommit.ServicePackage(ctx),
		codeguruprofiler.ServicePackage(ctx),
		codegurureviewer.ServicePackage(ctx),
		codepipeline.ServicePackage(ctx),
		codestarconnections.ServicePackage(ctx),
		codestarnotifications.ServicePackage(ctx),
		cognitoidentity.ServicePackage(ctx),
		cognitoidp.ServicePackage(ctx),
		comprehend.ServicePackage(ctx),
		computeoptimizer.ServicePackage(ctx),
		configservice.ServicePackage(ctx),
		connect.ServicePackage(ctx),
		connectcases.ServicePackage(ctx),
		controltower.ServicePackage(ctx),
		costoptimizationhub.ServicePackage(ctx),
		cur.ServicePackage(ctx),
		customerprofiles.ServicePackage(ctx),
		databrew.ServicePackage(ctx),
		dataexchange.ServicePackage(ctx),
		datapipeline.ServicePackage(ctx),
		datasync.ServicePackage(ctx),
		datazone.ServicePackage(ctx),
		dax.ServicePackage(ctx),
		deploy.ServicePackage(ctx),
		detective.ServicePackage(ctx),
		devicefarm.ServicePackage(ctx),
		devopsguru.ServicePackage(ctx),
		directconnect.ServicePackage(ctx),
		dlm.ServicePackage(ctx),
		dms.ServicePackage(ctx),
		docdb.ServicePackage(ctx),
		docdbelastic.ServicePackage(ctx),
		drs.ServicePackage(ctx),
		ds.ServicePackage(ctx),
		dynamodb.ServicePackage(ctx),
		ec2.ServicePackage(ctx),
		ecr.ServicePackage(ctx),
		ecrpublic.ServicePackage(ctx),
		ecs.ServicePackage(ctx),
		efs.ServicePackage(ctx),
		eks.ServicePackage(ctx),
		elasticache.ServicePackage(ctx),
		elasticbeanstalk.ServicePackage(ctx),
		elasticsearch.ServicePackage(ctx),
		elastictranscoder.ServicePackage(ctx),
		elb.ServicePackage(ctx),
		elbv2.ServicePackage(ctx),
		emr.ServicePackage(ctx),
		emrcontainers.ServicePackage(ctx),
		emrserverless.ServicePackage(ctx),
		events.ServicePackage(ctx),
		evidently.ServicePackage(ctx),
		finspace.ServicePackage(ctx),
		firehose.ServicePackage(ctx),
		fis.ServicePackage(ctx),
		fms.ServicePackage(ctx),
		fsx.ServicePackage(ctx),
		gamelift.ServicePackage(ctx),
		glacier.ServicePackage(ctx),
		globalaccelerator.ServicePackage(ctx),
		glue.ServicePackage(ctx),
		grafana.ServicePackage(ctx),
		greengrass.ServicePackage(ctx),
		groundstation.ServicePackage(ctx),
		guardduty.ServicePackage(ctx),
		healthlake.ServicePackage(ctx),
		iam.ServicePackage(ctx),
		identitystore.ServicePackage(ctx),
		imagebuilder.ServicePackage(ctx),
		inspector.ServicePackage(ctx),
		inspector2.ServicePackage(ctx),
		internetmonitor.ServicePackage(ctx),
		iot.ServicePackage(ctx),
		iotanalytics.ServicePackage(ctx),
		iotevents.ServicePackage(ctx),
		ivs.ServicePackage(ctx),
		ivschat.ServicePackage(ctx),
		kafka.ServicePackage(ctx),
		kafkaconnect.ServicePackage(ctx),
		kendra.ServicePackage(ctx),
		keyspaces.ServicePackage(ctx),
		kinesis.ServicePackage(ctx),
		kinesisanalytics.ServicePackage(ctx),
		kinesisanalyticsv2.ServicePackage(ctx),
		kinesisvideo.ServicePackage(ctx),
		kms.ServicePackage(ctx),
		lakeformation.ServicePackage(ctx),
		lambda.ServicePackage(ctx),
		launchwizard.ServicePackage(ctx),
		lexmodels.ServicePackage(ctx),
		lexv2models.ServicePackage(ctx),
		licensemanager.ServicePackage(ctx),
		lightsail.ServicePackage(ctx),
		location.ServicePackage(ctx),
		logs.ServicePackage(ctx),
		lookoutmetrics.ServicePackage(ctx),
		m2.ServicePackage(ctx),
		macie2.ServicePackage(ctx),
		mediaconnect.ServicePackage(ctx),
		mediaconvert.ServicePackage(ctx),
		medialive.ServicePackage(ctx),
		mediapackage.ServicePackage(ctx),
		mediapackagev2.ServicePackage(ctx),
		mediastore.ServicePackage(ctx),
		memorydb.ServicePackage(ctx),
		meta.ServicePackage(ctx),
		mq.ServicePackage(ctx),
		mwaa.ServicePackage(ctx),
		neptune.ServicePackage(ctx),
		neptunegraph.ServicePackage(ctx),
		networkfirewall.ServicePackage(ctx),
		networkmanager.ServicePackage(ctx),
		networkmonitor.ServicePackage(ctx),
		oam.ServicePackage(ctx),
		opensearch.ServicePackage(ctx),
		opensearchserverless.ServicePackage(ctx),
		opsworks.ServicePackage(ctx),
		organizations.ServicePackage(ctx),
		osis.ServicePackage(ctx),
		outposts.ServicePackage(ctx),
		paymentcryptography.ServicePackage(ctx),
		pcaconnectorad.ServicePackage(ctx),
		pcs.ServicePackage(ctx),
		pinpoint.ServicePackage(ctx),
		pipes.ServicePackage(ctx),
		polly.ServicePackage(ctx),
		pricing.ServicePackage(ctx),
		qbusiness.ServicePackage(ctx),
		qldb.ServicePackage(ctx),
		quicksight.ServicePackage(ctx),
		ram.ServicePackage(ctx),
		rbin.ServicePackage(ctx),
		rds.ServicePackage(ctx),
		redshift.ServicePackage(ctx),
		redshiftdata.ServicePackage(ctx),
		redshiftserverless.ServicePackage(ctx),
		rekognition.ServicePackage(ctx),
		resiliencehub.ServicePackage(ctx),
		resourceexplorer2.ServicePackage(ctx),
		resourcegroups.ServicePackage(ctx),
		resourcegroupstaggingapi.ServicePackage(ctx),
		rolesanywhere.ServicePackage(ctx),
		route53.ServicePackage(ctx),
		route53domains.ServicePackage(ctx),
		route53profiles.ServicePackage(ctx),
		route53recoverycontrolconfig.ServicePackage(ctx),
		route53recoveryreadiness.ServicePackage(ctx),
		route53resolver.ServicePackage(ctx),
		rum.ServicePackage(ctx),
		s3.ServicePackage(ctx),
		s3control.ServicePackage(ctx),
		s3outposts.ServicePackage(ctx),
		sagemaker.ServicePackage(ctx),
		scheduler.ServicePackage(ctx),
		schemas.ServicePackage(ctx),
		secretsmanager.ServicePackage(ctx),
		securityhub.ServicePackage(ctx),
		securitylake.ServicePackage(ctx),
		serverlessrepo.ServicePackage(ctx),
		servicecatalog.ServicePackage(ctx),
		servicecatalogappregistry.ServicePackage(ctx),
		servicediscovery.ServicePackage(ctx),
		servicequotas.ServicePackage(ctx),
		ses.ServicePackage(ctx),
		sesv2.ServicePackage(ctx),
		sfn.ServicePackage(ctx),
		shield.ServicePackage(ctx),
		signer.ServicePackage(ctx),
		simpledb.ServicePackage(ctx),
		sns.ServicePackage(ctx),
		sqs.ServicePackage(ctx),
		ssm.ServicePackage(ctx),
		ssmcontacts.ServicePackage(ctx),
		ssmincidents.ServicePackage(ctx),
		ssmsap.ServicePackage(ctx),
		sso.ServicePackage(ctx),
		ssoadmin.ServicePackage(ctx),
		storagegateway.ServicePackage(ctx),
		sts.ServicePackage(ctx),
		swf.ServicePackage(ctx),
		synthetics.ServicePackage(ctx),
		timestreaminfluxdb.ServicePackage(ctx),
		timestreamwrite.ServicePackage(ctx),
		transcribe.ServicePackage(ctx),
		transfer.ServicePackage(ctx),
		verifiedpermissions.ServicePackage(ctx),
		vpclattice.ServicePackage(ctx),
		waf.ServicePackage(ctx),
		wafregional.ServicePackage(ctx),
		wafv2.ServicePackage(ctx),
		wellarchitected.ServicePackage(ctx),
		worklink.ServicePackage(ctx),
		workspaces.ServicePackage(ctx),
		workspacesweb.ServicePackage(ctx),
		xray.ServicePackage(ctx),
	}

	return slices.Clone(v)
}
