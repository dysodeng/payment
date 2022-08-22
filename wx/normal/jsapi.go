package normal

import (
	"fmt"
	"time"

	"github.com/dysodeng/payment/support"
	"github.com/dysodeng/payment/support/crypto/rsa"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	payJsApi "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

// jsApi JSAPI支付
type jsApi struct {
	payment *Payment
}

// Prepay 支付下单
// @param description string 商品描述
// @param outTradeNo string 商户订单号
// @param amount int64 支付金额,单位为分
// @param openid string 支付者公众账号openid
// @param attach string 附加数据,最大长度128位字符
// @param notifyUrl string 微信支付结果通知回调地址
func (jsApi *jsApi) Prepay(description, outTradeNo string, amount int64, openid, attach, notifyUrl string) (resp *payJsApi.PrepayResponse, result *core.APIResult, err error) {
	svc := payJsApi.JsapiApiService{Client: jsApi.payment.client}
	return svc.Prepay(jsApi.payment.ctx, payJsApi.PrepayRequest{
		Appid:       core.String(jsApi.payment.config.AppID),
		Mchid:       core.String(jsApi.payment.config.MchID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Attach:      core.String(attach),
		NotifyUrl:   core.String(notifyUrl),
		Amount: &payJsApi.Amount{
			Total:    core.Int64(amount),
			Currency: core.String("CNY"),
		},
		Payer: &payJsApi.Payer{
			Openid: core.String(openid),
		},
	})
}

// JsSdkConfig 构建微信支付jssdk配置
func (jsApi *jsApi) JsSdkConfig(prePayId string) (map[string]interface{}, error) {
	timestamp := time.Now().Unix()
	config := map[string]interface{}{
		"appId":     jsApi.payment.config.AppID,
		"timeStamp": fmt.Sprintf("%d", timestamp),
		"nonceStr":  support.RandStringBytesMask(32),
		"package":   fmt.Sprintf("prepay_id=%s", prePayId),
		"signType":  "RSA",
	}

	signString := fmt.Sprintf("%s\n%d\n%s\n%s\n", config["appId"], timestamp, config["nonceStr"], config["package"])

	sign, err := rsa.Encrypt(signString, jsApi.payment.config.MchPrivateKey)
	if err != nil {
		return nil, err
	}

	config["paySign"] = sign

	return config, nil
}

// CloseOrder 关闭订单
// @param outTradeNo string 商户订单号
func (jsApi *jsApi) CloseOrder(outTradeNo string) (result *core.APIResult, err error) {
	svc := payJsApi.JsapiApiService{Client: jsApi.payment.client}
	return svc.CloseOrder(jsApi.payment.ctx, payJsApi.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(jsApi.payment.config.MchID),
	})
}

// QueryOrderById 微信支付订单号查询订单
// @param transactionId string 微信支付订单号
func (jsApi *jsApi) QueryOrderById(transactionId string) (resp *payments.Transaction, result *core.APIResult, err error) {
	svc := payJsApi.JsapiApiService{Client: jsApi.payment.client}
	return svc.QueryOrderById(jsApi.payment.ctx, payJsApi.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(jsApi.payment.config.MchID),
	})
}

// QueryOrderByOutTradeNo 商户订单号查询订单
// @param outTradeNo string 商户订单号
func (jsApi *jsApi) QueryOrderByOutTradeNo(outTradeNo string) (resp *payments.Transaction, result *core.APIResult, err error) {
	svc := payJsApi.JsapiApiService{Client: jsApi.payment.client}
	return svc.QueryOrderByOutTradeNo(jsApi.payment.ctx, payJsApi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(jsApi.payment.config.MchID),
	})
}
