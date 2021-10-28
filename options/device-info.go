package options

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// DeviceInfo IoT 设备信息
type DeviceInfo struct {
	IotInstanceID string `json:"iotInstanceId"` // 服务实例标识
	ProductKey    string `json:"productKey"`    // 产品标识
	DeviceName    string `json:"deviceName"`    // 设备标识
}

// GenerateRequest 生成服务调用请求
func (deviceInfo *DeviceInfo) GenerateRequest(payload *ServiceRequestPayload) (req *iot.RRpcRequest, err error) {
	base64Payload, err := payload.SerializeBase64()
	if err != nil {
		return
	}

	req = iot.CreateRRpcRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = deviceInfo.IotInstanceID
	req.ProductKey = deviceInfo.ProductKey
	req.DeviceName = deviceInfo.DeviceName
	req.Timeout = requests.NewInteger(8000) // 8 秒
	req.RequestBase64Byte = base64Payload

	return
}
