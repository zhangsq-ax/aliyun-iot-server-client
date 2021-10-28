package options

import (
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
)

type ServiceRequestPayload struct {
	ID      string      `json:"id"`
	Version string      `json:"version"`
	Params  interface{} `json:"params"`
	Method  string      `json:"method"`
}

func NewServiceRequestPayload(method string, params interface{}) *ServiceRequestPayload {
	return &ServiceRequestPayload{
		ID:      uuid.Must(uuid.NewRandom()).String(),
		Version: "1.0",
		Params:  params,
		Method:  method,
	}
}

func (rp *ServiceRequestPayload) Serialize() ([]byte, error) {
	return json.Marshal(rp)
}

func (rp *ServiceRequestPayload) SerializeBase64() (string, error) {
	b, err := rp.Serialize()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
