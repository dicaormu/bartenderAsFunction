package dao

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"os"
	"time"
)

type IotShadowDoc struct {
	Version  int               `json:"version,omitempty"`
	State    IotShadowStates   `json:"state,omitempty"`
	Metadata iotObjectMetadata `json:"metadata,omitempty"`
}

type IotShadowStates struct {
	Desired  IotObjectState `json:"desired,omitempty"`
	Reported IotObjectState `json:"reported,omitempty"`
}

type IotObjectState struct {
	BarStatus string `json:"barStatus,omitempty"`
}

type iotObjectMetadata struct {
	Desired  desiredMetadata `json:"desired"`
	Reported desiredMetadata `json:"reported"`
}

type desiredMetadata struct {
	BarStatus timestampType `json:"barStatus,omitempty"`
}

type timestampType struct {
	Timestamp Timestamp `json:"timestamp"`
}

type Timestamp struct {
	time.Time
}

func UpdateBarShadow(deviceId string, desiredBarStatus string) error {
	clientIot := initializeIotDataClient()
	input := iotdataplane.UpdateThingShadowInput{}
	input.SetThingName(deviceId)
	desiredShadow := IotObjectState{BarStatus: desiredBarStatus}
	shadow := IotShadowDoc{
		State: IotShadowStates{
			Desired: desiredShadow,
		},
	}
	payload, _ := json.Marshal(shadow)
	fmt.Println("payload " + string(payload))
	input.SetPayload(payload)
	_, err := clientIot.UpdateThingShadow(&input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr.Message())
		} else {
			fmt.Println(err)
		}
		return err
	}
	return nil
}

func initializeIotClient() *iot.IoT {
	sessionVar, err := session.NewSession(&aws.Config{Region: aws.String( /*os.Getenv("AWS_DEFAULT_REGION")*/ "eu-west-1")})
	if err != nil {
		fmt.Println("error during aws session initialization :  " + err.Error())
		os.Exit(1)
	}
	return iot.New(sessionVar, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
}

func initializeIotDataClient() *iotdataplane.IoTDataPlane {
	awsConfig := &aws.Config{Region: aws.String( /*os.Getenv("AWS_DEFAULT_REGION")*/ "eu-west-1")}
	sessionVar, err := session.NewSession(awsConfig.WithLogLevel(aws.LogDebugWithHTTPBody))
	if err != nil {
		fmt.Println("error during aws session initialization :  " + err.Error())
		os.Exit(1)
	}
	iotSvc := iot.New(sessionVar)
	return InitializeIotDataClientFromIotClient(iotSvc)
}

func InitializeIotDataClientFromIotClient(iotSvc *iot.IoT) *iotdataplane.IoTDataPlane {
	res, err := iotSvc.DescribeEndpoint(&iot.DescribeEndpointInput{})
	sessionVar, err := session.NewSession(&iotSvc.Config)
	if err != nil {
		fmt.Println("error during aws session initialization :  " + err.Error())
		os.Exit(1)
	}
	return iotdataplane.New(sessionVar, &aws.Config{Endpoint: res.EndpointAddress})
}
