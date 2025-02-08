# lambda-zoo

A zoo of lambda implementations.

What I mean is: a set of similar little lambda functions with different
configurations and deployment options.

I'm doing this for practice, but I thought it might be useful to someone 
to have working examples.

## List of Functions
### Basic-lambda group

The Basic lambdas are exposed to be accessible as a web endpoint somehow
and return a _hello world_ message that include their individual name
when HTTP GET'ed.

#### Python

  * [basic-lambda-python-cdk](basiclambda/python/basic-lambda-python-cdk): 

    Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL.

  * [basic-lambda-python-cdk-explicit-role](basiclambda/python/basic-lambda-python-cdk-explicit-role): 

    Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL. Explicitly creates a role.

  * [basic-lambda-python-pulumi](basiclambda/python/basic-lambda-python-pulumi): 

    Implemented in Python. Deployed with Pulumi in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL.

#### Golang

* [basic-lambda-go-cdk](basiclambda/go/basic-lambda-go-cdk):

  Implemented in Go. Deployed with CDK in Go. Packaged as a simple zip without dependencies (using.
  Exposed using a Lambda URL. CDK provides teh default role.

  Uses the [@aws-cdk/aws-lambda-go-alpha](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-go-alpha-readme.html) module for packaging

* [basic-lambda-go-cdk-explicit-role](basiclambda/go/basic-lambda-go-cdk-explicit-role):

  Implemented in Go. Deployed with CDK in Go. Packaged as a simple zip without dependencies.
  Exposed using a Lambda URL. Explicitly creates a role.

  Uses the [@aws-cdk/aws-lambda-go-alpha](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-go-alpha-readme.html) module for packaging

