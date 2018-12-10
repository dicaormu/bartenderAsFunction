package model

type DrunkClient struct {
	IdClient       string `json:"id"`
	CertificateArn string `json:"certificateArn"`
	PrivateKey     string `json:"privateKey"`
	PublicKey      string `json:"publicKey"`
	CertificatePem string `json:"certificatePem"`
}
