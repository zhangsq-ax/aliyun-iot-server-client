package options

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

type DeviceInfo struct {
	IotInstanceID string `json:"iotInstanceId"`
	ProductKey    string `json:"productKey"`
	DeviceName    string `json:"deviceName"`
}

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
