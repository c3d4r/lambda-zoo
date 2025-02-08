from aws_cdk import (
    Stack,
    aws_lambda as lambda_, CfnOutput,
)
from constructs import Construct

class BasicLambdaPythonStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        fn = lambda_.Function(self, "BasicLambdaPythonCdk",
                         runtime=lambda_.Runtime.PYTHON_3_12,
                         handler="lambda_handler.lambda_handler",
                         code=lambda_.Code.from_asset("../app")
                         )

        fn_url = fn.add_function_url(auth_type=lambda_.FunctionUrlAuthType.NONE)

        CfnOutput(self, "functionArn", value=fn.function_arn, description="The ARN of the provisioned function")
        CfnOutput(self, "functionUrl", value=fn_url.url, description="The URL of the provisioned function")



