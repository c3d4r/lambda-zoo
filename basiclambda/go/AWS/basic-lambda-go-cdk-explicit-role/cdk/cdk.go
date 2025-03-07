package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Define the IAM role with trust policy for Lambda
	lambdaRole := awsiam.NewRole(stack, jsii.String("BasicLambdaGoRole"), &awsiam.RoleProps{
		RoleName:  jsii.String("basic-lambda-go-cdk-explicit-role-role"),
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})

	// Attach Managed Policies (e.g., AWSLambdaBasicExecutionRole)
	lambdaRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))

	fn := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("basic-lambda-go-cdk-explicit-role"),
		&awscdklambdagoalpha.GoFunctionProps{
			FunctionName: jsii.String("basic-lambda-go-cdk-explicit-role"),
			Role:         lambdaRole,
			Entry:        jsii.String("../lambda/main.go"),
			Bundling: &awscdklambdagoalpha.BundlingOptions{
				GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
			},
			Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		})

	fn_url := fn.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	awscdk.NewCfnOutput(stack, jsii.String("functionArn"), &awscdk.CfnOutputProps{
		Value:       fn.FunctionArn(),
		Description: jsii.String("The ARN of the provisioned function"),
	})

	awscdk.NewCfnOutput(stack, jsii.String("functionUrl"), &awscdk.CfnOutputProps{
		Value:       fn_url.Url(),
		Description: jsii.String("The URL of the provisioned function"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
