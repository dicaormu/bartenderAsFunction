package dao

import (
	"bartenderAsFunction/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"os"
)

type IotConnection struct {
	Iot *iot.IoT
}

type IotConnectionInterface interface {
	RegisterDevice(drunkClient  *model.DrunkClient) error
	UpdateShadow(idClient, status string)error
}

func (con *IotConnection) RegisterDevice(drunkClient  *model.DrunkClient) error {
	var input iot.CreateKeysAndCertificateInput
	input.SetSetAsActive(true)
	output, errCert := con.Iot.CreateKeysAndCertificate(&input)
	if errCert != nil {
		return errCert
	}
	drunkClient.CertificateArn = *output.CertificateArn
	drunkClient.PrivateKey = *output.KeyPair.PrivateKey
	drunkClient.PublicKey = *output.KeyPair.PublicKey
	drunkClient.CertificatePem = *output.CertificatePem
	return nil
}

/*
func (con *IotConnection) CloseBarForClient(idClient string) error {
	errChangeStatus := con.UpdateShadow(idClient, "CLOSED")
	return errChangeStatus
}*/

func (con *IotConnection) UpdateShadow(idClient string, desiredStatus string) error {
	input := iotdataplane.UpdateThingShadowInput{}
	input.SetThingName(idClient)
	var desiredShadow model.ClientObjectState

	desiredShadow.BarStatus = desiredStatus
	shadow := model.IotShadowDoc{
		State: model.IotShadowState{
			Desired: desiredShadow,
		},
	}
	payload, _ := json.Marshal(shadow)
	input.SetPayload(payload)
	iotShadowConn := initializeIotDataClient(con.Iot)
	_, err := iotShadowConn.UpdateThingShadow(&input)
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

func CreateIotConnection() IotConnectionInterface {
	return &IotConnection{initializeIotClient()}
}

func initializeIotDataClient(iotSvc *iot.IoT) *iotdataplane.IoTDataPlane {
	res, err := iotSvc.DescribeEndpoint(&iot.DescribeEndpointInput{})
	sessionVar, err := session.NewSession(&iotSvc.Config)
	if err != nil {
		fmt.Println("error during aws session initialization :  " + err.Error())
		os.Exit(1)
	}
	return iotdataplane.New(sessionVar, &aws.Config{Endpoint: res.EndpointAddress})
}

func initializeIotClient() *iot.IoT {
	sessionVar, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})
	if err != nil {
		fmt.Println("error during aws session initialization :  ", err)
		os.Exit(1)
	}
	return iot.New(sessionVar, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
}
