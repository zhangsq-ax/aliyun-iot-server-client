package service_invoker

import (
	"github.com/zhangsq-ax/aliyun-iot-server-client/options"
	"os"
	"testing"
)

var invoker *ServiceInvoker

var (
	regionId        string
	accessKeyId     string
	accessKeySecret string
	instanceId      string
	productKey      string
	deviceName      string
)

func init() {
	regionId = os.Getenv("IOT_REGION")
	accessKeyId = os.Getenv("IOT_ACCESS_KEY_ID")
	accessKeySecret = os.Getenv("IOT_ACCESS_KEY_SECRET")
	instanceId = os.Getenv("IOT_INSTANCE_ID")
	productKey = os.Getenv("IOT_PRODUCT_KEY")
	deviceName = os.Getenv("IOT_DEVICE_NAME")
}

func TestNewServiceInvoker(t *testing.T) {
	var err error
	invoker, err = NewServiceInvoker(&options.ServiceInvokerOptions{
		RegionID:        regionId,
		AccessKeyID:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestServiceInvoker_InvokeService(t *testing.T) {
	err := invoker.InvokeService(&options.DeviceInfo{
		IotInstanceID: instanceId,
		ProductKey:    productKey,
		DeviceName:    deviceName,
	}, "open", map[string]int{
		"duration": 10,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestServiceInvoker_InvokeServiceSync(t *testing.T) {
	res, err := invoker.InvokeServiceSync(&options.DeviceInfo{
		IotInstanceID: instanceId,
		ProductKey:    productKey,
		DeviceName:    deviceName,
	}, options.NewServiceSyncRequestPayloadByIdentifier("call", map[string]int{
		"floor": 10,
	}))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

/* func TestServiceInvoker_Invoke(t *testing.T) {
	res, err := invoker.InvokeServiceSync(&options.DeviceInfo{
		IotInstanceID: instanceId,
		ProductKey:    productKey,
		DeviceName:    deviceName,
	}, options.NewServiceSyncRequestPayload("thing.service.reserve", map[string]interface{}{
		"order_id": "test-order",
		"detail": []map[string]int{
			{
				"product_id": 2468,
				"amount":     2,
			},
		},
	}))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)
} */
