# The bartender as a function

## This repository is in construction

## requirements

* go > 1.10 
* go dep
* [sam local](https://github.com/awslabs/aws-sam-cli)
* docker

## How to run tests

If you want to run all test: 

```
$ go test ./...

```

if you want to run a particular test:

```
$ go test bartenderAsFunction/functions/FUNCTION_FOLDER/FILE -run TESTNAME

```

example:

```
$ go test bartenderAsFunction/functions/getCommand/ -run TestHandlerShouldReturn404

```

if you want to clean cache for tests:

```
$ go clean -testcache

```

to build 

```
$ ./build.sh

```

test with sam local

```
$ sam local start-api --template sam.yml
```

to deploy: create profile xebia and:

```
$ ./deploy.sh USER
```