package normal

import (
	"context"
	"crypto/rsa"

	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Payment 微信支付普通模式
type Payment struct {
	client *core.Client
	ctx    context.Context
	config PaymentConfig
}

// PaymentConfig 支付配置
type PaymentConfig struct {
	MchCertificateSerialNumber string // 商户证书序列号
	MchAPIv3Key                string // 商户api v3密钥
	MchPrivateKey              string // 商户私钥
	MchID                      string // 商户号
	AppID                      string // 关联公众号/小程序AppID
}

// NewPayment 新建普通模式支付
func NewPayment(ctx context.Context, config PaymentConfig) (*Payment, error) {

	var mchPrivateKey *rsa.PrivateKey
	var err error
	if config.MchPrivateKey != "" {
		mchPrivateKey, err = utils.LoadPrivateKey(config.MchPrivateKey)
		if err != nil {
			return nil, errors.Wrap(err, "load merchant private key error")
		}
	} else {
		mchPrivateKey = nil
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MchCertificateSerialNumber, mchPrivateKey, config.MchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "new wechat pay client err")
	}

	return &Payment{
		ctx:    ctx,
		config: config,
		client: client,
	}, nil
}

// Native 原生扫码支付
func (p *Payment) Native() *native {
	return &native{
		payment: p,
	}
}

// JsApi JSAPI支付
func (p *Payment) JsApi() *jsApi {
	return &jsApi{
		payment: p,
	}
}

// H5 h5支付
func (p *Payment) H5() *h5 {
	return &h5{
		payment: p,
	}
}

// Notify 支付通知
func (p *Payment) Notify() *notify {
	return &notify{
		payment: p,
	}
}
