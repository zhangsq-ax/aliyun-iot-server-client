package amqp_manager

import (
	"context"
	"github.com/zhangsq-ax/aliyun-iot-server-client/options"
	"os"
	"pack.ag/amqp"
	"strconv"
	"testing"
	"time"
)

var manager *AMQPManager
var consumer *AMQPConsumer

var (
	Host            string
	Port            int
	InstanceID      string
	ClientID        string
	ConsumerGroupID string
	AccessKeyID     string
	AccessKeySecret string
)

func init() {
	Host = os.Getenv("AMQP_HOST")
	port := os.Getenv("AMQP_PORT")
	InstanceID = os.Getenv("AMQP_INSTANCE_ID")
	ClientID = os.Getenv("AMQP_CLIENT_ID")
	ConsumerGroupID = os.Getenv("AMQP_CONSUMER_GROUP_ID")
	AccessKeyID = os.Getenv("AMQP_ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("AMQP_ACCESS_KEY_SECRET")

	var err error
	Port, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

}

func TestNewAMQPManager(t *testing.T) {
	manager = NewAMQPManager(&options.AMQPOptions{
		Host:            Host,
		Port:            Port,
		InstanceID:      InstanceID,
		ClientID:        ClientID,
		ConsumerGroupID: ConsumerGroupID,
		AccessKeyID:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
	})
}

func TestAMQPManager_CreateAMQPConsumer(t *testing.T) {
	consumer = manager.CreateAMQPConsumer()
}

func TestAMQPConsumer_Start(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)
	err := consumer.Start(ctx, func(msg *amqp.Message) {
		t.Log("-------------------------")
		t.Log(string(msg.GetData()))
		t.Log(msg.ApplicationProperties)
	})
	if err != nil {
		t.Error(err)
	}
}
