package options

import (
	"encoding/base64"
	"encoding/json"
)

// ServiceResponse 服务响应
type ServiceResponse struct {
	ID   string `json:"id"`   // 响应标识
	Code int    `json:"code"` // 响应状态标识，详情参考 https://help.aliyun.com/document_detail/89309.htm
}

// NewServiceResponseByBase64 通过 Base64 字符串创建服务响应对象
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
