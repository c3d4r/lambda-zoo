# basic-lambda-python-terraform

Implemented in Python. Deployed with Terraform. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL.


# basic-lambda-go-pulumi

Implemented in Go. Deployed with Pulumi in Terraform. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL.

## Deploy

```shell
cd terraform
terraform init
terraform plan
terraform apply
```

## Invoke

Use the `lambda_function_url` value from the deployment output.

_or_

Find it using:

```shell
export FN_URL=`terraform output -json | jq -r '.lambda_function_url.value'`
echo $FN_URL
```

Then run a cURL command like:

```shell
curl $FN_URL
```
It should look something like:

``` 
❯ curl $FN_URL
"hello world from basic-lambda-golang-pulumi"%   
```


## Destroy

```shell
terraform destroy
```

