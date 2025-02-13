package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
	"os/exec"
)

const (
	shell      = "sh"
	shellFlag  = "-c"
	rootFolder = "../lambda"
)

func runCmd(args string) error {
	cmd := exec.Command(shell, shellFlag, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = rootFolder
	return cmd.Run()
}

func main() {
	if err := runCmd("GOOS=linux GOARCH=arm64 go build -o ./bootstrap ./handler.go"); err != nil {
		fmt.Printf("Error building code: %s", err.Error())
		os.Exit(1)
	}

	if err := runCmd("zip -r -j ./handler.zip ./bootstrap"); err != nil {
		fmt.Printf("Error creating zipfile: %s", err.Error())
		os.Exit(1)
	}
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Policy for the lambda function
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "basic-lambda-go-pulumi-role", &iam.RoleArgs{
			Name:             pulumi.String("basic-lambda-go-pulumi-role"),
			AssumeRolePolicy: pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}

		// Attach a policy to allow writing logs to CloudWatch
		logPolicy, err := iam.NewRolePolicy(ctx, "lambda-log-policy", &iam.RolePolicyArgs{
			Role: role.Name,
			Policy: pulumi.String(`{
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "logs:CreateLogGroup",
                        "logs:CreateLogStream",
                        "logs:PutLogEvents"
                    ],
                    "Resource": "arn:aws:logs:*:*:*"
                }]
            }`),
		})
		if err != nil {
			return err
		}

		// Create the lambda using the args.
		function, err := lambda.NewFunction(
			ctx,
			"basicLambda",
			&lambda.FunctionArgs{
				Handler:       pulumi.String("bootstrap"),
				Role:          role.Arn,
				Runtime:       pulumi.String("provided.al2"),
				Architectures: pulumi.StringArray{pulumi.String("arm64")},
				Code:          pulumi.NewFileArchive("../lambda/handler.zip"),
			},
			pulumi.DependsOn([]pulumi.Resource{logPolicy}),
		)
		if err != nil {
			return err
		}

		fnUrl, err := lambda.NewFunctionUrl(ctx, "basic-lambda-go-pulumi", &lambda.FunctionUrlArgs{
			FunctionName:      function.Name,
			AuthorizationType: pulumi.String("NONE"),
		})
		if err != nil {
			return err
		}

		// Export the lambda ARN.
		ctx.Export("lambda", function.Arn)
		ctx.Export("lambdaUrl", fnUrl.FunctionUrl)

		return nil

	})
}
