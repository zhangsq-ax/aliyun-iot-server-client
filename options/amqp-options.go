package options

// AMQPOptions AMQP 服务连接设置
type AMQPOptions struct {
	Host            string `json:"host"`            // 服务终端节点
	Port            int    `json:"port"`            // 服务终端端口
	InstanceID      string `json:"instanceId"`      // IoT 服务实例标识
	ClientID        string `json:"clientId"`        // 客户端标识
	ConsumerGroupID string `json:"consumerGroupId"` // 消费组标识
	AccessKeyID     string `json:"accessKeyId"`     // 密钥 ID
	AccessKeySecret string `json:"accessKeySecret"` // 密钥
}
