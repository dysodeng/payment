package normal

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	payNative "github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

// native 原生扫码支付
type native struct {
	payment *Payment
}

// Prepay 支付下单
// @param description string 商品描述
// @param outTradeNo string 商户订单号
// @param amount int64 支付金额,单位为分
// @param attach string 附加数据,最大长度128位字符
// @param notifyUrl string 微信支付结果通知回调地址
func (native *native) Prepay(description, outTradeNo string, amount int64, attach, notifyUrl string) (resp *payNative.PrepayResponse, result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.Prepay(native.payment.ctx, payNative.PrepayRequest{
		Appid:       core.String(native.payment.config.AppID),
		Mchid:       core.String(native.payment.config.MchID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Attach:      core.String(attach),
		NotifyUrl:   core.String(notifyUrl),
		Amount: &payNative.Amount{
			Total:    core.Int64(amount),
			Currency: core.String("CNY"),
		},
	})
}

// CloseOrder 关闭订单
// @param outTradeNo string 商户订单号
func (native *native) CloseOrder(outTradeNo string) (result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.CloseOrder(native.payment.ctx, payNative.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(native.payment.config.MchID),
	})
}

// QueryOrderById 微信支付订单号查询订单
// @param transactionId string 微信支付订单号
func (native *native) QueryOrderById(transactionId string) (resp *payments.Transaction, result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.QueryOrderById(native.payment.ctx, payNative.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(native.payment.config.MchID),
	})
}

// QueryOrderByOutTradeNo 商户订单号查询订单
// @param outTradeNo string 商户订单号
func (native *native) QueryOrderByOutTradeNo(outTradeNo string) (resp *payments.Transaction, result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.QueryOrderByOutTradeNo(native.payment.ctx, payNative.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(native.payment.config.MchID),
	})
}
