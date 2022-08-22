package alipay

import (
	"net/url"

	"github.com/dysodeng/payment/support/crypto/rsa"
)

// AliPay 支付宝
type AliPay struct {
	config *config
}

// New create alipay
// @param appId string 应用ID
// @param alipayPublicKey string 支付宝公钥
// @param privateKey string 应用私钥
func New(appId, alipayPublicKey, privateKey string, opts ...Option) *AliPay {
	c := &config{
		isDev:           false,
		appId:           appId,
		alipayPublicKey: alipayPublicKey,
		privateKey:      privateKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return &AliPay{
		config: c,
	}
}

// CheckCallbackSign 检查支付回调参数签名
func (pay *AliPay) CheckCallbackSign(params url.Values) (bool, error) {
	m := make(map[string]string)
	for key, value := range params {
		if key == "sign" || key == "sign_type" {
			continue
		}
		m[key] = value[0]
	}

	sign := params["sign"][0]

	return rsa.Check(pay.signString(m), sign, pay.alipayPublicKey())
}
