package model

type ClientObjectState struct {
	BarStatus string `json:"barStatus"`
}

type IotShadowState struct {
	Desired  ClientObjectState `json:"desired,omitempty"`
	Reported ClientObjectState `json:"reported,omitempty"`
}

type IotShadowDoc struct {
	Version  int               `json:"version,omitempty"`
	State    IotShadowState    `json:"state,omitempty"`
	Metadata iotObjectMetadata `json:"metadata,omitempty"`
}

type desiredMetadata struct {
	BarStatus timestampType `json:"barStatus,omitempty"`
}

type timestampType struct {
	Timestamp int64 `json:"timestamp"`
}

type iotObjectMetadata struct {
	Desired  desiredMetadata `json:"desired"`
	Reported desiredMetadata `json:"reported"`
}

type IotEvent struct {
	Previous  IotShadowDoc `json:"previous"`
	Current   IotShadowDoc `json:"current"`
	Timestamp int64        `json:"timestamp"`
	DeviceId  string       `json:"deviceid"`
}
