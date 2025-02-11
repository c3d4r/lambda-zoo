package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(_ context.Context, r events.APIGatewayProxyRequest) (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "hello world from basic-lambda-go-cdk",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
