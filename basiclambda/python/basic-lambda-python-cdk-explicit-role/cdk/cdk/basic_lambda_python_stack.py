from aws_cdk import (
    Stack,
    aws_lambda as lambda_,
    aws_iam as iam,
    CfnOutput,
)
from constructs import Construct

class BasicLambdaPythonStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # Define the IAM role with trust policy for Lambda
        lambda_role = iam.Role(
            self, "BasicLambdaPythonRole",
            role_name="basic-lambda-python-cdk-explicit-role-role",
            assumed_by=iam.ServicePrincipal("lambda.amazonaws.com")
        )

        # Attach Managed Policies (e.g., AWSLambdaBasicExecutionRole)
        lambda_role.add_managed_policy(
            iam.ManagedPolicy.from_aws_managed_policy_name("service-role/AWSLambdaBasicExecutionRole")
        )

        fn = lambda_.Function(self, "BasicLambdaPythonCdkExplicitRole",
                              runtime=lambda_.Runtime.PYTHON_3_12,
                              handler="lambda_handler.lambda_handler",
                              code=lambda_.Code.from_asset("../app"),
                              role=lambda_role,
                         )

        fn_url = fn.add_function_url(auth_type=lambda_.FunctionUrlAuthType.NONE)

        CfnOutput(self, "functionArn", value=fn.function_arn, description="The ARN of the provisioned function")
        CfnOutput(self, "functionUrl", value=fn_url.url, description="The URL of the provisioned function")



