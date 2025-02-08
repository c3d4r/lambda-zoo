# basic-lambda-python-pulumi

Implemented in Python. Deployed with Pulumi in Python. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL.

## Deploy

```shell
pulumi up
```

It should look something like:

```
❯ pulumi up
Previewing update (dev)                                                                                                                                                                                ─╯

View in Browser (Ctrl+O): https://app.pulumi.com/[SOME_ACCOUNT]/basic-lambda-python-pulumi/dev/previews/[SOME_UUID]

     Type                       Name                            Plan       
 +   pulumi:pulumi:Stack        basic-lambda-python-pulumi-dev  create     
 +   ├─ aws:iam:Role            iam_for_lambda                  create     
 +   ├─ aws:lambda:FunctionUrl  basic-lambda-python-pulumi      create     
 +   └─ aws:lambda:Function     basic-lambda-python-pulumi      create     

Outputs:
    function_url: output<string>
    lambda_arn  : output<string>

Resources:
    + 4 to create

Do you want to perform this update? yes
Updating (dev)

View in Browser (Ctrl+O): https://app.pulumi.com/[SOME_ACCOUNT]/basic-lambda-python-pulumi/dev/updates/7

     Type                       Name                            Status            
 +   pulumi:pulumi:Stack        basic-lambda-python-pulumi-dev  created (19s)     
 +   ├─ aws:iam:Role            iam_for_lambda                  created (2s)      
 +   ├─ aws:lambda:Function     basic-lambda-python-pulumi      created (12s)     
 +   └─ aws:lambda:FunctionUrl  basic-lambda-python-pulumi      created (1s)      

Outputs:
    function_url: "https://[SOME_ID].lambda-url.eu-west-1.on.aws/"
    lambda_arn  : "arn:aws:lambda:eu-west-1:[SOME_ACCOUNT]:function:basic-lambda-python-pulumi-df7ad63"

Resources:
    + 4 created

Duration: 22s

```

## Invoke

Use the `function_url` value from the deployment output.

_or_

Find it using:

```shell
export FN_URL=`pulumi stack output function_url`
echo $FN_URL
```

Then run a cURL command like:

```shell
curl $FN_URL
```
It should look something like:

``` 
❯ curl $FN_URL
"Hello world from basic-lambda-python-pulumi Lambda!"%   
```


## Destroy

```shell
pulumi destroy
```

It should look something like:

```
❯ pulumi destroy
Previewing destroy (dev)                                                                                                                                                                               ─╯

View in Browser (Ctrl+O): https://app.pulumi.com/[SOME_ACCOUNT]/basic-lambda-python-pulumi/dev/previews/[SOME_UUID]

     Type                       Name                            Plan       
 -   pulumi:pulumi:Stack        basic-lambda-python-pulumi-dev  delete     
 -   ├─ aws:lambda:Function     basic-lambda-python-pulumi      delete     
 -   ├─ aws:iam:Role            iam_for_lambda                  delete     
 -   └─ aws:lambda:FunctionUrl  basic-lambda-python-pulumi      delete     

Outputs:
  - function_url: "https://[SOME_ID].lambda-url.eu-west-1.on.aws/"
  - lambda_arn  : "arn:aws:lambda:eu-west-1:[SOME_ACCOUNT]:function:basic-lambda-python-pulumi-df7ad63"

Resources:
    - 4 to delete

Do you want to perform this destroy? yes
Destroying (dev)

View in Browser (Ctrl+O): https://app.pulumi.com/[SOME_ACCOUNT]/basic-lambda-python-pulumi/dev/updates/8

     Type                       Name                            Status              
 -   pulumi:pulumi:Stack        basic-lambda-python-pulumi-dev  deleted (0.72s)     
 -   ├─ aws:lambda:FunctionUrl  basic-lambda-python-pulumi      deleted (1s)        
 -   ├─ aws:lambda:Function     basic-lambda-python-pulumi      deleted (0.73s)     
 -   └─ aws:iam:Role            iam_for_lambda                  deleted (1s)        

Outputs:
  - function_url: "https://[SOME_ID].lambda-url.eu-west-1.on.aws/"
  - lambda_arn  : "arn:aws:lambda:eu-west-1:[SOME_ACCOUNT]:function:basic-lambda-python-pulumi-df7ad63"

Resources:
    - 4 deleted

Duration: 9s

The resources in the stack have been deleted, but the history and configuration associated with the stack are still maintained. 
If you want to remove the stack completely, run `pulumi stack rm dev`.

```