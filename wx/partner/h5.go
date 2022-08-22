package partner

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	payH5 "github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/h5"
)

// h5 支付
type h5 struct {
	payment  *Payment
	subAppID string // 子商户AppID
	subMchID string // 子商户号
}

// h5SceneType H5场景类型
type h5SceneType string

const (
	H5SceneTypeWap     h5SceneType = "Wap"
	H5SceneTypeIOS     h5SceneType = "iOS"
	H5SceneTypeAndroid h5SceneType = "Android"
)

type h5SceneOption struct {
	deviceId      string
	storeInfo     *h5StoreInfo
	appName       string
	appUrl        string
	bundleId      string
	packageName   string
	profitSharing bool
}

type h5StoreInfo struct {
	id       string
	name     string
	areaCode string
	address  string
}

type H5Option func(*h5SceneOption)

// WithH5StoreInfo 设置门店信息
func WithH5StoreInfo(storeId, storeName, areaCode, address string) H5Option {
	return func(option *h5SceneOption) {
		option.storeInfo = &h5StoreInfo{
			id:       storeId,
			name:     storeName,
			areaCode: areaCode,
			address:  address,
		}
	}
}

// WithH5DeviceId 设置商户终端设备号
func WithH5DeviceId(deviceId string) H5Option {
	return func(option *h5SceneOption) {
		option.deviceId = deviceId
	}
}

// WithH5AppInfo 设置应用信息
func WithH5AppInfo(appName, appUrl string) H5Option {
	return func(option *h5SceneOption) {
		option.appName = appName
		option.appUrl = appUrl
	}
}

// WithH5iOS 设置iOS平台BundleID
func WithH5iOS(bundleId string) H5Option {
	return func(option *h5SceneOption) {
		option.bundleId = bundleId
	}
}

// WithH5Android 设置安卓平台PackageName
func WithH5Android(packageName string) H5Option {
	return func(option *h5SceneOption) {
		option.packageName = packageName
	}
}

// WithH5ProfitSharing 指定分账
func WithH5ProfitSharing() H5Option {
	return func(option *h5SceneOption) {
		option.profitSharing = true
	}
}

// Prepay H5支付预下单
func (h5 *h5) Prepay(
	description,
	outTradeNo string,
	amount int64,
	attach,
	notifyUrl,
	payerClientIp string,
	h5SceneType h5SceneType,
	opts ...H5Option,
) (resp *payH5.PrepayResponse, result *core.APIResult, err error) {

	sceneInfo := &payH5.SceneInfo{
		PayerClientIp: core.String(payerClientIp),
		H5Info: &payH5.H5Info{
			Type: core.String(string(h5SceneType)),
		},
	}

	o := &h5SceneOption{}
	for _, opt := range opts {
		opt(o)
	}

	if o.storeInfo != nil {
		sceneInfo.StoreInfo = &payH5.StoreInfo{
			Id:       core.String(o.storeInfo.id),
			Name:     core.String(o.storeInfo.name),
			AreaCode: core.String(o.storeInfo.areaCode),
			Address:  core.String(o.storeInfo.address),
		}
	}
	if o.deviceId != "" {
		sceneInfo.DeviceId = core.String(o.deviceId)
	}
	if o.appUrl != "" {
		sceneInfo.H5Info.AppUrl = core.String(o.appUrl)
	}
	if o.appName != "" {
		sceneInfo.H5Info.AppName = core.String(o.appName)
	}
	if o.bundleId != "" {
		sceneInfo.H5Info.BundleId = core.String(o.bundleId)
	}
	if o.packageName != "" {
		sceneInfo.H5Info.PackageName = core.String(o.packageName)
	}

	svc := payH5.H5ApiService{Client: h5.payment.client}

	return svc.Prepay(h5.payment.ctx, payH5.PrepayRequest{
		SpAppid:     core.String(h5.payment.config.AppID),
		SpMchid:     core.String(h5.payment.config.MchID),
		SubAppid:    core.String(h5.subAppID),
		SubMchid:    core.String(h5.subMchID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Attach:      core.String(attach),
		NotifyUrl:   core.String(notifyUrl),
		SceneInfo:   sceneInfo,
		Amount: &payH5.Amount{
			Total:    core.Int64(amount),
			Currency: core.String("CNY"),
		},
		SettleInfo: &payH5.SettleInfo{
			ProfitSharing: core.Bool(o.profitSharing),
		},
	})
}

// CloseOrder 关闭订单
// @param outTradeNo string 商户订单号
func (h5 *h5) CloseOrder(outTradeNo string) (result *core.APIResult, err error) {
	svc := payH5.H5ApiService{Client: h5.payment.client}
	return svc.CloseOrder(h5.payment.ctx, payH5.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		SpMchid:    core.String(h5.payment.config.MchID),
		SubMchid:   core.String(h5.subMchID),
	})
}

// QueryOrderById 微信支付订单号查询订单
// @param transactionId string 微信支付订单号
func (h5 *h5) QueryOrderById(transactionId string) (resp *partnerpayments.Transaction, result *core.APIResult, err error) {
	svc := payH5.H5ApiService{Client: h5.payment.client}
	return svc.QueryOrderById(h5.payment.ctx, payH5.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		SpMchid:       core.String(h5.payment.config.MchID),
		SubMchid:      core.String(h5.subMchID),
	})
}

// QueryOrderByOutTradeNo 商户订单号查询订单
// @param outTradeNo string 商户订单号
func (h5 *h5) QueryOrderByOutTradeNo(outTradeNo string) (resp *partnerpayments.Transaction, result *core.APIResult, err error) {
	svc := payH5.H5ApiService{Client: h5.payment.client}
	return svc.QueryOrderByOutTradeNo(h5.payment.ctx, payH5.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		SpMchid:    core.String(h5.payment.config.MchID),
		SubMchid:   core.String(h5.subMchID),
	})
}
