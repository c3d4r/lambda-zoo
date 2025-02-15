# basic-lambda-python-cdk

Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL.

## Deploy

Make sure you are using Python 3.12

```sh 
# Enter the correct directory
cd basiclambda/python/basic-lambda-python-cdk/cdk
# Make a virtual environment
python -m venv .venv
# Activate the virtual environment
source .venv/bin/activate
# Install the dependencies needed for the CDK setup
pip install -r requirements.txt
pip install -r requirements-dev.txt
# Deploy
cdk deploy
```
It should look something like:

``` 
> cdk deploy
╰─ cdk deploy                                                                                                                                                                       ─╯
✨  Synthesis time: 5.02s

BasicLambdaPythonStack: deploying... [1/1]

 ✅  BasicLambdaPythonStack (no changes)

✨  Deployment time: 2.4s

Outputs:
BasicLambdaPythonStack.functionArn = _SOME_ARN_
BasicLambdaPythonStack.functionUrl = _SOME_URL_
Stack ARN:
_SOME_ARN_

✨  Total time: 7.42s


```


## Invoke

Use the `functionUrl` value from the deployment output.

_or_

Find it using:

```shell
FN_URL=`aws cloudformation describe-stacks --stack-name BasicLambdaPythonStack --output json \
 | jq -r '.Stacks[].Outputs[] | select(.OutputKey=="functionUrl") | .OutputValue'`
echo $FN_URL
```

Then run a cURL command like:

```shell
curl $FN_URL
```
It should look something like:

```
❯ curl $FN_URL
"Hello world from basic-lambda-python-cdk Lambda!"%  
```

## Destroy

Destroy using:

```shell
cdk destroy
```

It should look something like:

```
> cdk destroy
Are you sure you want to delete: BasicLambdaPythonStack (y/n)? y
BasicLambdaPythonStack: destroying... [1/1]
BasicLambdaPythonStack |   0 | 10:46:47 PM | DELETE_IN_PROGRESS   | AWS::CloudFormation::Stack | BasicLambdaPythonStack User Initiated
BasicLambdaPythonStack |   0 | 10:46:49 PM | DELETE_IN_PROGRESS   | AWS::Lambda::Url        | BasicLambdaPythonCdk/FunctionUrl (BasicLambdaPythonCdkFunctionUrlB296FC69) 
BasicLambdaPythonStack |   0 | 10:46:49 PM | DELETE_IN_PROGRESS   | AWS::CDK::Metadata      | CDKMetadata/Default (CDKMetadata) 
BasicLambdaPythonStack |   0 | 10:46:49 PM | DELETE_IN_PROGRESS   | AWS::Lambda::Permission | BasicLambdaPythonCdk/invoke-function-url (BasicLambdaPythonCdkinvokefunctionurl072F0E06) 
BasicLambdaPythonStack |   1 | 10:46:49 PM | DELETE_COMPLETE      | AWS::CDK::Metadata      | CDKMetadata/Default (CDKMetadata) 
BasicLambdaPythonStack |   2 | 10:46:50 PM | DELETE_COMPLETE      | AWS::Lambda::Url        | BasicLambdaPythonCdk/FunctionUrl (BasicLambdaPythonCdkFunctionUrlB296FC69) 
BasicLambdaPythonStack |   3 | 10:46:50 PM | DELETE_COMPLETE      | AWS::Lambda::Permission | BasicLambdaPythonCdk/invoke-function-url (BasicLambdaPythonCdkinvokefunctionurl072F0E06) 
BasicLambdaPythonStack |   3 | 10:46:50 PM | DELETE_IN_PROGRESS   | AWS::Lambda::Function   | BasicLambdaPythonCdk (BasicLambdaPythonCdk915C4A44) 
BasicLambdaPythonStack |   4 | 10:46:54 PM | DELETE_COMPLETE      | AWS::Lambda::Function   | BasicLambdaPythonCdk (BasicLambdaPythonCdk915C4A44) 
BasicLambdaPythonStack |   4 | 10:46:54 PM | DELETE_IN_PROGRESS   | AWS::IAM::Role          | BasicLambdaPythonCdk/ServiceRole (BasicLambdaPythonCdkServiceRole5F89E436) 

 ✅  BasicLambdaPythonStack: destroyed

```
