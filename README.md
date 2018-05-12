# The bartender as a function

## Objectives

The main objective of this Workshop is to make you to code in go language by creating a serverless application
and to discover the go language and the architectures microservices over aws using lambdas.

It is presented as a simple game: You are the waiter of a bar. You should wait for the commands of the clients and generate the bill.

Don't worry, I have coded for you the client and the bartender. :relaxed:
I give you also the structure of the *waiter* project with the *aws-sam* templates and the unit tests of the services (functions) you should code 

## Before you start: Requirements

* go > 1.10 
* go dep
* [sam local](https://github.com/awslabs/aws-sam-cli)
* docker
* a profile "xebia" for aws-cli with your aws credentials

### Creation of a profile for aws cli

Go to you aws console, to your security credentials, and create an access key. Copy your aws_access_key_id
and your aws_secret_access_key  and  create your ~/.aws/credentials and  ~/.aws/config files as stated in this instructions https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html
Your profile should be called *xebia* (or modify the provided scripts to use the name of the profile you want to use)

## The exercise 

During this exercise, if you have any question, you can go to the [faq](FAQ.md)

### Step 1
The client is an IOT device who asks for a list of food and/or beer
As waiter you have to:  
* *listen* to those messages (see the sam.yml file, LambdaRuleReadCommand)
* generate and id for the command
* save the command to a Dynamodb table

Again, don't worry, you have to execute all tests in file *readCommand_test.go* , once they "pass", you can be sure your function works
There are also "todo" comments in the *readCommand.go* file

### Step 2
The bartender ask you for the unattended commands, it means, you should expose a rest api to allow the bartender to ask you for those commands.
As waiter you should:
* *listen* to the rest calls of the bartender (see the sam.yml file, LambdaGetCommands)
* read from the database all commands you have not served
* return them to the bartender

This time, if you execute all tests in *getCommand_test.go* file, and make them pass, everything should be ok.

### Step 3
The bartender gives you a command to serve, it means, you should expose a rest api to allow the bartender to ask you to serve those commands.
As waiter you should:
* *Know* the command the bartender ask you to "attend" (see the sam.yml file, LambdaServeCommands). The bartender gives you the "id" of the command and the item he wants you to serve
* read the command from the database
* change the status of the item to "served" 

Tests are implemented in *serveCommand_test.go*

### Step 4
The bartender ask you for the served commands, it means, you should expose a rest api to allow the bartender to ask you for those commands.
As waiter you should:
* *listen* to the rest calls of the bartender (see the sam.yml file, LambdaGetFacture)
* read from the database all commands you have served
* return them to the bartender

Tests are implemented in *getFacture_test.go*

Deploy your solution and enjoy!!!

## Register your api:
post to https://dpcrgtoori.execute-api.eu-west-1.amazonaws.com/test/registry
body:
{ "name":"USERXX",
"url":"URL_API"
}