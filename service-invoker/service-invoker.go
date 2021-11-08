package service_invoker

import (
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

// Invoke 调用服务
func (invoker *ServiceInvoker) Invoke(device *options.DeviceInfo, payload *options.ServiceRequestPayload) (response *options.ServiceResponse, err error) {
	req, err := device.GenerateRequest(payload)
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
