# The bartender as a function

## Objectives

The main objective of this Workshop is to learn how to use an IoT platform (aws) for the things fleet industrialization.

It is presented as a simple game: You are the waiter of a bar. Each client is represented by a "thing".
You should wait for the clients' commands and generate the bill.

![the bartender](https://github.com/dicaormu/bartenderAsFunction/blob/solution/bartenderHL.png "The bartender")


Don't worry, I have coded for you the client and the bartender. :relaxed: .

I give you also the structure of the *waiter* project. It includes the *aws-sam* templates and the unit tests of the services (functions) you should code in Nodejs or Go languages.

## Before you start: Requirements

* go > 1.10 
* go dep
* [sam local](https://github.com/awslabs/aws-sam-cli)
* a profile "epf" for aws-cli with your aws credentials

### Creating a profile for aws cli

Go to your security credentials in your aws console, and create an access key. Copy your aws_access_key_id
and your aws_secret_access_key  and  create your ~/.aws/credentials and  ~/.aws/config files as stated in this instructions: https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html
Your profile should be called *epf* (or modify the provided scripts to use the name of the profile you want to use).

## The exercise 

Create your GOPATH environment variable with the path to your go projects.
Inside this path, create a folder src.

```
$ cd $GOPATH/src

$ git clone https://github.com/dicaormu/bartenderAsFunction
```


During this exercise, if you have any question, you can go to the [faq](FAQ.md).

![the exercise](https://docs.google.com/drawings/d/e/2PACX-1vQo9d9tz8Mm0s_NxGLRni0yA6V7r6YDlaJtOHQLblMqXi9jWjkIfv-v8L0eHsnF_XSIbTK2Yg7tecY0/pub?w=480&h=360)

### Step 1
The client is an IOT device who is going to send a command.
As waiter you have to:
* Announce where is the client going to register
* Allow clients to register to the IoT Platform

Don't worry, I've coded the client for you.

### Step 2
When the client send a command, as waiter you have to:
* *listen* to those messages (see the sam.yml file, LambdaRuleReadCommandBeer,LambdaRuleReadCommandFood )
* create 2 rules to send the message to be treated by the right lambda (you have to modify the file sam.yml, see [the AWS documentation](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) for more information)
* generate and id for the command
* limit the beer commands to 1 per 2 minutes (we don't want a drunk client)
* save the command to a Dynamodb table

Again, don't worry, you have to execute all tests in file *readCommand_test.go* , once they "pass", you can be sure your function works
There are also "todo" comments in the *readCommand.go* file.

### Step 3
The bartender ask you for the unattended commands, it means, you should expose a rest api to allow the bartender to ask you for those commands.
As waiter you should:
* *listen* to the rest calls of the bartender (see the sam.yml file, LambdaGetCommands)
* read from the database all commands you have not served
* return them to the bartender

This time, if you execute all tests in *getCommand_test.go* file, and make them pass, everything should be ok.

### Step 4
The client is very drunk. As a waiter you are going to close the bar for him (no more service)
As waiter you should:

* *Know* the status of the client (you should use the shadow for that)
* change the property "barStatus" of the shadow of the client to "CLOSED" (LambdaStatusBar)
* verify that the client does not send more information and clean all commands (pass them to "served", for this point you should see the file sam.yml to get the event from the shadow update and LambdaGetFacture)

To know more about aws shadow and how to update the thing, see [the aws documentation](https://docs.aws.amazon.com/iot/latest/developerguide/device-shadow-mqtt.html)

Tests are implemented in *FILE_test.go*.

Deploy your solution and enjoy!!!

## Register your api:
Make a post to the URL i'll give you in the workshop with the body:

```json
{ 
"name":"USERXX",
"url":"URL_OF_YOUR_DEPLOYEDAPI"
}
```