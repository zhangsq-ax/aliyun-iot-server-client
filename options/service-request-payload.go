package options

import (
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
)

// ServiceRequestPayload 服务调用载荷
type ServiceRequestPayload struct {
	ID      string      `json:"id"`      // 调用标识
	Version string      `json:"version"` // 协议版本，固定为 "1.0"
	Params  interface{} `json:"params"`  // 输入参数
	Method  string      `json:"method"`  // 服务标识，格式如：thing.service.<serviceId>
}

// NewServiceRequestPayload 创建新的服务调用载荷
func NewServiceRequestPayload(method string, params interface{}) *ServiceRequestPayload {
	return &ServiceRequestPayload{
		ID:      uuid.Must(uuid.NewRandom()).String(),
		Version: "1.0",
		Params:  params,
		Method:  method,
	}
}

// Serialize 序列化服务调用载荷
func (rp *ServiceRequestPayload) Serialize() ([]byte, error) {
	return json.Marshal(rp)
}

// SerializeBase64 将调用载荷序列化为 Base64 字符串
func (rp *ServiceRequestPayload) SerializeBase64() (string, error) {
	b, err := rp.Serialize()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
