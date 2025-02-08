# basic-lambda-python-cdk

Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL. Explicitly creates a role.

## Deploy

Make sure you are using Python 3.12

```sh 
# Enter the correct directory
cd basiclambda/python/basic-lambda-python-cdk-explicit-role/cdk
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
❯ cdk deploy
╰─ cdk deploy
✨  Synthesis time: 11.79s

BasicLambdaPythonStack: start: Building 885ffbf05605b5e3b263bffb2a812df29e7d205168e3e5de2358c286a0af2261:336205929843-eu-west-1
BasicLambdaPythonStack: success: Built 885ffbf05605b5e3b263bffb2a812df29e7d205168e3e5de2358c286a0af2261:336205929843-eu-west-1
BasicLambdaPythonStack: start: Building b7ee6773b19c0042d31f71b287e50dccaab3b9b231909b624c1e588367863471:336205929843-eu-west-1
BasicLambdaPythonStack: success: Built b7ee6773b19c0042d31f71b287e50dccaab3b9b231909b624c1e588367863471:336205929843-eu-west-1
BasicLambdaPythonStack: start: Publishing b7ee6773b19c0042d31f71b287e50dccaab3b9b231909b624c1e588367863471:336205929843-eu-west-1
BasicLambdaPythonStack: start: Publishing 885ffbf05605b5e3b263bffb2a812df29e7d205168e3e5de2358c286a0af2261:336205929843-eu-west-1
BasicLambdaPythonStack: success: Published 885ffbf05605b5e3b263bffb2a812df29e7d205168e3e5de2358c286a0af2261:336205929843-eu-west-1
BasicLambdaPythonStack: success: Published b7ee6773b19c0042d31f71b287e50dccaab3b9b231909b624c1e588367863471:336205929843-eu-west-1
Stack undefined
This deployment will make potentially sensitive changes according to your current security approval level (--require-approval broadening).
Please confirm you intend to make the following modifications:

IAM Statement Changes
┌───┬─────────────────────────────────────────┬────────┬──────────────────────────┬──────────────────────────────┬───────────┐
│   │ Resource                                │ Effect │ Action                   │ Principal                    │ Condition │
├───┼─────────────────────────────────────────┼────────┼──────────────────────────┼──────────────────────────────┼───────────┤
│ + │ ${BasicLambdaPythonCdkExplicitRole.Arn} │ Allow  │ lambda:InvokeFunctionUrl │ *                            │           │
├───┼─────────────────────────────────────────┼────────┼──────────────────────────┼──────────────────────────────┼───────────┤
│ + │ ${BasicLambdaPythonRole.Arn}            │ Allow  │ sts:AssumeRole           │ Service:lambda.amazonaws.com │           │
└───┴─────────────────────────────────────────┴────────┴──────────────────────────┴──────────────────────────────┴───────────┘
IAM Policy Changes
┌───┬──────────────────────────┬────────────────────────────────────────────────────────────────────────────────┐
│   │ Resource                 │ Managed Policy ARN                                                             │
├───┼──────────────────────────┼────────────────────────────────────────────────────────────────────────────────┤
│ + │ ${BasicLambdaPythonRole} │ arn:${AWS::Partition}:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole │
└───┴──────────────────────────┴────────────────────────────────────────────────────────────────────────────────┘
(NOTE: There may be security-related changes not in this list. See https://github.com/aws/aws-cdk/issues/1299)

Do you wish to deploy these changes (y/n)? y
BasicLambdaPythonStack: deploying... [1/1]
BasicLambdaPythonStack: creating CloudFormation changeset...
BasicLambdaPythonStack | 0/6 | 10:21:28 PM | REVIEW_IN_PROGRESS   | AWS::CloudFormation::Stack | BasicLambdaPythonStack User Initiated
BasicLambdaPythonStack | 0/6 | 10:21:35 PM | CREATE_IN_PROGRESS   | AWS::CloudFormation::Stack | BasicLambdaPythonStack User Initiated
BasicLambdaPythonStack | 0/6 | 10:21:38 PM | CREATE_IN_PROGRESS   | AWS::IAM::Role          | BasicLambdaPythonRole (BasicLambdaPythonRole3D584C23) 
BasicLambdaPythonStack | 0/6 | 10:21:38 PM | CREATE_IN_PROGRESS   | AWS::CDK::Metadata      | CDKMetadata/Default (CDKMetadata) 
BasicLambdaPythonStack | 0/6 | 10:21:39 PM | CREATE_IN_PROGRESS   | AWS::IAM::Role          | BasicLambdaPythonRole (BasicLambdaPythonRole3D584C23) Resource creation Initiated
BasicLambdaPythonStack | 0/6 | 10:21:39 PM | CREATE_IN_PROGRESS   | AWS::CDK::Metadata      | CDKMetadata/Default (CDKMetadata) Resource creation Initiated
BasicLambdaPythonStack | 1/6 | 10:21:39 PM | CREATE_COMPLETE      | AWS::CDK::Metadata      | CDKMetadata/Default (CDKMetadata) 
BasicLambdaPythonStack | 2/6 | 10:21:57 PM | CREATE_COMPLETE      | AWS::IAM::Role          | BasicLambdaPythonRole (BasicLambdaPythonRole3D584C23) 
BasicLambdaPythonStack | 2/6 | 10:21:58 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Function   | BasicLambdaPythonCdkExplicitRole (BasicLambdaPythonCdkExplicitRoleAD853CE5) 
BasicLambdaPythonStack | 2/6 | 10:21:59 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Function   | BasicLambdaPythonCdkExplicitRole (BasicLambdaPythonCdkExplicitRoleAD853CE5) Resource creation Initiated
BasicLambdaPythonStack | 2/6 | 10:22:00 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Function   | BasicLambdaPythonCdkExplicitRole (BasicLambdaPythonCdkExplicitRoleAD853CE5) Eventual consistency check initiated
BasicLambdaPythonStack | 2/6 | 10:22:00 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Url        | BasicLambdaPythonCdkExplicitRole/FunctionUrl (BasicLambdaPythonCdkExplicitRoleFunctionUrlA77F332E) 
BasicLambdaPythonStack | 2/6 | 10:22:00 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Permission | BasicLambdaPythonCdkExplicitRole/invoke-function-url (BasicLambdaPythonCdkExplicitRoleinvokefunctionurlCD7A1CCB) 
BasicLambdaPythonStack | 2/6 | 10:22:01 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Permission | BasicLambdaPythonCdkExplicitRole/invoke-function-url (BasicLambdaPythonCdkExplicitRoleinvokefunctionurlCD7A1CCB) Resource creation Initiated
BasicLambdaPythonStack | 2/6 | 10:22:01 PM | CREATE_IN_PROGRESS   | AWS::Lambda::Url        | BasicLambdaPythonCdkExplicitRole/FunctionUrl (BasicLambdaPythonCdkExplicitRoleFunctionUrlA77F332E) Resource creation Initiated
BasicLambdaPythonStack | 3/6 | 10:22:01 PM | CREATE_COMPLETE      | AWS::Lambda::Url        | BasicLambdaPythonCdkExplicitRole/FunctionUrl (BasicLambdaPythonCdkExplicitRoleFunctionUrlA77F332E) 
BasicLambdaPythonStack | 4/6 | 10:22:01 PM | CREATE_COMPLETE      | AWS::Lambda::Permission | BasicLambdaPythonCdkExplicitRole/invoke-function-url (BasicLambdaPythonCdkExplicitRoleinvokefunctionurlCD7A1CCB) 
BasicLambdaPythonStack | 5/6 | 10:22:05 PM | CREATE_COMPLETE      | AWS::Lambda::Function   | BasicLambdaPythonCdkExplicitRole (BasicLambdaPythonCdkExplicitRoleAD853CE5) 
BasicLambdaPythonStack | 6/6 | 10:22:06 PM | CREATE_COMPLETE      | AWS::CloudFormation::Stack | BasicLambdaPythonStack 

 ✅  BasicLambdaPythonStack

✨  Deployment time: 52.26s

Outputs:
BasicLambdaPythonStack.functionArn = arn:aws:lambda:eu-west-1:[SOME_ACCOUNT]:function:BasicLambdaPythonStack-BasicLambdaPythonCdkExplici-KbnNn2Sfigbp
BasicLambdaPythonStack.functionUrl = https://[SOME_ID].lambda-url.eu-west-1.on.aws/
Stack ARN:
arn:aws:cloudformation:eu-west-1:[SOME_ACCOUNT]:stack/BasicLambdaPythonStack/[SOME_UID]

✨  Total time: 64.05s
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
