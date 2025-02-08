#!/usr/bin/env python3
import os

import aws_cdk as cdk

from cdk.basic_lambda_python_stack import BasicLambdaPythonStack


app = cdk.App()
BasicLambdaPythonStack(app, "BasicLambdaPythonExplicitRoleStack",
                       env=cdk.Environment(account=os.getenv('CDK_DEFAULT_ACCOUNT'), region=os.getenv('CDK_DEFAULT_REGION')),
                       )

tags = {
    "Project": "lambda-zoo",
    "Item": "basic-lambda-python-cdk-explicit-role",
    "IAC": "cdk",
    "Language": "python"
}

for key, value in tags.items():
    cdk.Tags.of(app).add(key, value)

app.synth()