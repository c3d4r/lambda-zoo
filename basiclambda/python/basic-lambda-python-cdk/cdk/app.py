#!/usr/bin/env python3
import os

import aws_cdk as cdk

from cdk.basic_lambda_python_stack import BasicLambdaPythonStack


app = cdk.App()
BasicLambdaPythonStack(app, "BasicLambdaPythonStack",
                       env=cdk.Environment(account=os.getenv('CDK_DEFAULT_ACCOUNT'), region=os.getenv('CDK_DEFAULT_REGION')),
                       )

app.synth()
