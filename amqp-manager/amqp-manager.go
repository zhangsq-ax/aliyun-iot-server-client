package amqp_manager

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/zhangsq-ax/aliyun-iot-server-client/options"
	"time"
)

type AMQPManager struct {
	address  string
	username string
	password string
}

func NewAMQPManager(opts *options.AMQPOptions) *AMQPManager {
	address := fmt.Sprintf("amqps://%s:%d", opts.Host, opts.Port)
	timestamp := time.Now().UnixNano() / 1e6
	username := fmt.Sprintf("%s|authMode=aksign,signMethod=Hmacsha1,consumerGroupId=%s,authId=%s,iotInstanceId=%s,timestamp=%d|", opts.ClientID, opts.ConsumerGroupID, opts.AccessKeyID, opts.InstanceID, timestamp)
	stringToSign := fmt.Sprintf("authId=%s&timestamp=%d", opts.AccessKeyID, timestamp)
	hmacKey := hmac.New(sha1.New, []byte(opts.AccessKeySecret))
	hmacKey.Write([]byte(stringToSign))
	password := base64.StdEncoding.EncodeToString(hmacKey.Sum(nil))

	return &AMQPManager{
		address:  address,
		username: username,
		password: password,
	}
}

func (manager *AMQPManager) CreateAMQPConsumer() *AMQPConsumer {
	return NewAMQPConsumer(manager)
}
