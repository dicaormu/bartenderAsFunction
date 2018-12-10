package testUtils

import (
	"bartenderAsFunction/model"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type IotConnectionMock struct {
	mock.Mock
	DrunkClient       model.DrunkClient
	ExpectedError error
}

func (con *IotConnectionMock) RegisterDevice(drunkClient  *model.DrunkClient) error {
	bytes, _ := json.Marshal(drunkClient)
	drunkClient.CertificateArn="an arn"
	drunkClient.CertificatePem="a pem"
	drunkClient.PrivateKey="a private key"
	fmt.Println(string(bytes))
	return nil
}

func (con *IotConnectionMock) UpdateShadow(idClient, barStatus string) error {
	return nil
}
