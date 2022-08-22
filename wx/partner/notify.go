package partner

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	payNotify "github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
)

// notify 支付回调
type notify struct {
	payment  *Payment
	subAppID string // 子商户AppID
	subMchID string // 子商户号
}

// notifyResponse 支付通知响应体
type notifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Handler 获取支付回调Handler
func (notify *notify) Handler(request *http.Request, bizCallback func(transaction *partnerpayments.Transaction) error) (*notifyResponse, error) {
	// 获取平台证书访问器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(notify.payment.config.MchID)
	handler := payNotify.NewNotifyHandler(notify.payment.config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	transaction := new(partnerpayments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(notify.payment.ctx, request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		log.Printf("%+v", err)
		return &notifyResponse{Code: "fail", Message: "微信支付通知验签失败"}, errors.Wrap(err, "微信支付通知验签失败")
	}

	if notifyReq.EventType != "TRANSACTION.SUCCESS" {
		log.Printf("%+v", notifyReq)
		return &notifyResponse{Code: "fail", Message: "支付失败"}, errors.New("支付失败")
	}

	err = bizCallback(transaction)
	if err != nil {
		return &notifyResponse{Code: "fail", Message: "支付业务处理失败"}, errors.Wrap(err, "支付业务处理失败")
	}

	return &notifyResponse{Code: "success", Message: "支付成功"}, nil
}
