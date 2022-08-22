package partner

import (
	"context"
	"crypto/rsa"

	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Payment 微信支付服务商模式
type Payment struct {
	client *core.Client
	ctx    context.Context
	config PaymentConfig
}

// PaymentConfig 支付配置
type PaymentConfig struct {
	MchCertificateSerialNumber string // 服务商商户证书序列号
	MchAPIv3Key                string // 服务商商户api v3密钥
	MchPrivateKey              string // 服务商商户私钥
	MchID                      string // 服务商商户号
	AppID                      string // 服务商关联公众号/小程序AppID
}

// NewPayment 新建服务商模式支付
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
		client: client,
		ctx:    ctx,
		config: config,
	}, nil
}

// Native 原生扫码支付
func (p *Payment) Native(subAppID, subMchID string) *native {
	return &native{
		payment:  p,
		subMchID: subMchID,
		subAppID: subAppID,
	}
}

// JsApi JSAPI支付
func (p *Payment) JsApi(subAppID, subMchID string) *jsApi {
	return &jsApi{
		payment:  p,
		subMchID: subMchID,
		subAppID: subAppID,
	}
}

// H5 h5支付
func (p *Payment) H5(subAppID, subMchID string) *h5 {
	return &h5{
		payment:  p,
		subMchID: subMchID,
		subAppID: subAppID,
	}
}

// Notify 支付通知
func (p *Payment) Notify(subAppID, subMchID string) *notify {
	return &notify{
		payment:  p,
		subMchID: subMchID,
		subAppID: subAppID,
	}
}
