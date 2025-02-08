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

  * [basic-lambda-python-cdk](basiclambda/python/basic-lambda-python-cdk): 

    Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL.

  * [basic-lambda-python-cdk-explicit-role](basiclambda/python/basic-lambda-python-cdk-explicit-role): 

    Implemented in Python. Deployed with CDK in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL. Explicitly creates a role.

  * [basic-lambda-python-pulumi](basiclambda/python/basic-lambda-python-pulumi): 

    Implemented in Python. Deployed with Pulumi in Python. Packaged as a simple zip without dependencies.
    Exposed using a Lambda URL.
