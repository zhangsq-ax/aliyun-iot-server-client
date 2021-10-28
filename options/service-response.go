package options

import (
	"encoding/base64"
	"encoding/json"
)

type ServiceResponse struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
}

func NewServiceResponseByBase64(base64Payload string) (*ServiceResponse, error) {
	b, err := base64.StdEncoding.DecodeString(base64Payload)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(b))

	resPayload := &ServiceResponse{}
	err = json.Unmarshal(b, resPayload)

	return resPayload, err
}
