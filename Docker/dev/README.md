# Firebase

## Build

```
docker build -t hcst/lambda-go-demo --build-arg AWS_KEY=<aws IAM key for serverless> --build-arg AWS_SECRET=<aws IAM secret for serverless> .
```

## Run

```
docker run -it --rm -v "$Env:USERPROFILE\Repo\go:/go/src" -v "$Env:USERPROFILE\.ssh:/root/.ssh" hcst/lambda-go-demo /bin/bash
```

## Connect to stopped container

```
docker start <container name>
docker exec -it <container name> /bin/bash
```

## Serverless
* [Serverless](https://serverless.com/)
* [Serverless go lambda support](https://serverless.com/blog/framework-example-golang-lambda-support/)

```
$ serverless create -t aws-go-dep -p myservice
$ cd myservice
$ make
$ serverless deploy
$ serverless invoke -f hello
$ serverless remove
```
