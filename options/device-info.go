package options

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// DeviceInfo IoT 设备信息
type DeviceInfo struct {
	IotInstanceID string `json:"iotInstanceId"` // 服务实例标识
	ProductKey    string `json:"productKey"`    // 产品标识
	DeviceName    string `json:"deviceName"`    // 设备标识
}

// GenerateSyncRequest 生成服务同步调用请求
func (deviceInfo *DeviceInfo) GenerateSyncRequest(payload *ServiceSyncRequestPayload) (req *iot.RRpcRequest, err error) {
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

// GenerateRequest 生成服务的调用请求（异步）
func (deviceInfo *DeviceInfo) GenerateRequest(serviceId string, params interface{}) (req *iot.InvokeThingServiceRequest, err error) {
	var paramsByte []byte
	paramsByte, err = json.Marshal(params)
	if err != nil {
		return
	}

	req = iot.CreateInvokeThingServiceRequest()
	req.DeviceName = deviceInfo.DeviceName
	req.IotInstanceId = deviceInfo.IotInstanceID
	req.ProductKey = deviceInfo.ProductKey
	req.Identifier = serviceId
	req.Args = string(paramsByte)

	return
}

func (deviceInfo *DeviceInfo) GenerateSubscribeTopicRequest(topic []string) (req *iot.SubscribeTopicRequest, err error) {
	req = iot.CreateSubscribeTopicRequest()
	req.IotInstanceId = deviceInfo.IotInstanceID
	req.ProductKey = deviceInfo.ProductKey
	req.DeviceName = deviceInfo.DeviceName
	req.Topic = &topic

	return
}
