# basic-lambda-go-pulumi

Implemented in Go. Deployed with Pulumi in Go. Packaged as a simple zip without dependencies.
Exposed using a Lambda URL.

## Deploy

```shell
pulumi up
```



## Invoke

Use the `functionUrl` value from the deployment output.

_or_

Find it using:

```shell
export FN_URL=`pulumi stack output functionUrl`
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
pulumi destroy
```

