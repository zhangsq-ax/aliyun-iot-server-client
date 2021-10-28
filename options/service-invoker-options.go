package options

// ServiceInvokerOptions 服务调用者设置
type ServiceInvokerOptions struct {
	RegionID        string // IoT 平台服务地域标识，详情参考 https://help.aliyun.com/document_detail/40654.htm
	AccessKeyID     string // IoT 平台服务访问密钥标识
	AccessKeySecret string // IoT 平台服务访问密钥
}
