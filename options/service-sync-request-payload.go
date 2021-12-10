package options

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ServiceSyncRequestPayload 服务调用载荷
type ServiceSyncRequestPayload struct {
	ID      string      `json:"id"`      // 调用标识
	Version string      `json:"version"` // 协议版本，固定为 "1.0"
	Params  interface{} `json:"params"`  // 输入参数
	Method  string      `json:"method"`  // 服务标识，格式如：thing.service.<serviceId>
}

// NewServiceSyncRequestPayload 创建新的服务同步调用载荷
func NewServiceSyncRequestPayload(method string, params interface{}) *ServiceSyncRequestPayload {
	return &ServiceSyncRequestPayload{
		ID:      uuid.Must(uuid.NewRandom()).String(),
		Version: "1.0",
		Params:  params,
		Method:  method,
	}
}

// NewServiceSyncRequestPayloadByIdentifier 使用服务标识创建服务同步调用载荷
func NewServiceSyncRequestPayloadByIdentifier(serviceId string, params interface{}) *ServiceSyncRequestPayload {
	return NewServiceSyncRequestPayload(fmt.Sprintf("thing.service.%s", serviceId), params)
}

// Serialize 序列化服务同步调用载荷
func (rp *ServiceSyncRequestPayload) Serialize() ([]byte, error) {
	return json.Marshal(rp)
}

// SerializeBase64 将同步调用载荷序列化为 Base64 字符串
func (rp *ServiceSyncRequestPayload) SerializeBase64() (string, error) {
	b, err := rp.Serialize()
	if err != nil {
		return "", err
	}
	fmt.Println(string(b))
	return base64.StdEncoding.EncodeToString(b), nil
}
