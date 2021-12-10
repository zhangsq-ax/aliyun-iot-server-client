package service_invoker

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
	"github.com/zhangsq-ax/aliyun-iot-server-client/options"
)

// ServiceInvoker 服务调用者类
type ServiceInvoker struct {
	client *iot.Client
}

// NewServiceInvoker 创建服务调用者实例
func NewServiceInvoker(opts *options.ServiceInvokerOptions) (invoker *ServiceInvoker, err error) {
	client, err := iot.NewClientWithAccessKey(opts.RegionID, opts.AccessKeyID, opts.AccessKeySecret)
	if err != nil {
		return
	}

	invoker = &ServiceInvoker{client}
	return
}

// InvokeServiceSync 以同步方式（RRpc）调用服务，注意：此方式不会经过物模型，因此不产生服务调用记录也不检查数据结构合法性，只适合需要同步从设备获取时使用
func (invoker *ServiceInvoker) InvokeServiceSync(device *options.DeviceInfo, payload *options.ServiceSyncRequestPayload) (response *options.ServiceSyncResponse, err error) {
	req, err := device.GenerateSyncRequest(payload)
	if err != nil {
		return
	}

	res, err := invoker.client.RRpc(req)
	if err != nil {
		return
	}

	if !res.Success {
		err = fmt.Errorf("failed to invoke service: %s", res.String())
		return
	}

	response, err = options.NewServiceResponseByBase64(res.PayloadBase64Byte)
	return
}

// InvokeService 以异步的方式调用服务，不返回设备端响应，如需响应请订阅相应的 topic
func (invoker *ServiceInvoker) InvokeService(device *options.DeviceInfo, serviceId string, params interface{}) (err error) {
	req, err := device.GenerateRequest(serviceId, params)
	if err != nil {
		return
	}

	res, err := invoker.client.InvokeThingService(req)
	if !res.Success {
		data, _ := json.Marshal(res)
		fmt.Println(string(data))
		err = fmt.Errorf("failed to invoke thing service: %s - %s", res.Code, res.ErrorMessage)
	}
	return
}
