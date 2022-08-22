package alipay

const (
	prodGateway = "https://openapi.alipay.com/gateway.do"    // 线上环境
	devGateway  = "https://openapi.alipaydev.com/gateway.do" // 沙箱环境
)

// config 支付宝配置
type config struct {
	isDev           bool   // 是否为沙箱环境
	appId           string // 应用ID
	alipayPublicKey string // 支付宝公钥
	privateKey      string // 商户私钥
	notifyUrl       string // 支付结果异步通知地址
	returnUrl       string // 支付完成跳转地址
	appAuthToken    string // app auth token
}

type Option func(*config)

// WithDev 启用沙箱环境
func WithDev() Option {
	return func(c *config) {
		c.isDev = true
	}
}

// WithNotifyUrl 设置支付结果异步通知地址
func WithNotifyUrl(notifyUrl string) Option {
	return func(c *config) {
		c.notifyUrl = notifyUrl
	}
}

// WithReturnUrl 设置支付完成跳转页面地址
func WithReturnUrl(returnUrl string) Option {
	return func(c *config) {
		c.returnUrl = returnUrl
	}
}

// WithAppAuthToken setting app auth token
func WithAppAuthToken(appAuthToken string) Option {
	return func(c *config) {
		c.appAuthToken = appAuthToken
	}
}
