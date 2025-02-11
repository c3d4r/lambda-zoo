package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type BasicLambdaCdkStackProps struct {
	awscdk.StackProps
}

func NewBasicLambdaCdkStack(scope constructs.Construct, id string, props *BasicLambdaCdkStackProps) awscdk.Stack {

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Add a Lambda function to the stack
	fn := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("basic-lambda-go-cdk"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("basic-lambda-go-cdk"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Entry:        jsii.String("../lambda/main.go"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
		},
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

	NewBasicLambdaCdkStack(app, "BasicLambdaCdkStack", &BasicLambdaCdkStackProps{
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
