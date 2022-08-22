package partner

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	payNative "github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/native"
)

// native 原生扫码支付
type native struct {
	payment  *Payment
	subAppID string // 子商户AppID
	subMchID string // 子商户号
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
		SpAppid:     core.String(native.payment.config.AppID),
		SpMchid:     core.String(native.payment.config.MchID),
		SubAppid:    core.String(native.subAppID),
		SubMchid:    core.String(native.subMchID),
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
func (native *native) CloseOrder(outTradeNo string) (result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.CloseOrder(native.payment.ctx, payNative.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		SpMchid:    core.String(native.payment.config.MchID),
		SubMchid:   core.String(native.subMchID),
	})
}

// QueryOrderById 微信支付订单号查询订单
func (native *native) QueryOrderById(transactionId string) (resp *partnerpayments.Transaction, result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.QueryOrderById(native.payment.ctx, payNative.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		SpMchid:       core.String(native.payment.config.MchID),
		SubMchid:      core.String(native.subMchID),
	})
}

// QueryOrderByOutTradeNo 商户订单号查询订单
func (native *native) QueryOrderByOutTradeNo(outTradeNo string) (resp *partnerpayments.Transaction, result *core.APIResult, err error) {
	svc := payNative.NativeApiService{Client: native.payment.client}
	return svc.QueryOrderByOutTradeNo(native.payment.ctx, payNative.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		SpMchid:    core.String(native.payment.config.MchID),
		SubMchid:   core.String(native.subMchID),
	})
}
