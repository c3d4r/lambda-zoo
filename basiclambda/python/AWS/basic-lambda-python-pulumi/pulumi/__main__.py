import pulumi
from pulumi_aws import lambda_, iam

assume_role = iam.get_policy_document(statements=[{
    "effect": "Allow",
    "principals": [{
        "type": "Service",
        "identifiers": ["lambda.amazonaws.com"],
    }],
    "actions": ["sts:AssumeRole"],
}])
iam_for_lambda = iam.Role("iam_for_lambda",
    name="basic-lambda-python-pulumi-role",
    assume_role_policy=assume_role.json)

fn = lambda_.Function('basic-lambda-python-pulumi',
                      handler="lambda_handler.lambda_handler",
                      code= pulumi.AssetArchive({".": pulumi.FileArchive("../app")}),
                      role=iam_for_lambda.arn,
                      runtime="python3.12",
                      )

# Export the ARN of the function
pulumi.export('lambda_arn', fn.arn)

fn_url = lambda_.FunctionUrl("basic-lambda-python-pulumi",
    function_name=fn.name,
    authorization_type="NONE",
    cors={
        "allow_credentials": True,
        "allow_origins": ["*"],
        "allow_methods": ["*"],
        "allow_headers": [
            "date",
            "keep-alive",
        ],
        "expose_headers": [
            "keep-alive",
            "date",
        ],
        "max_age": 86400,
    })

# Export the URL of the function
pulumi.export('function_url', fn_url.function_url)