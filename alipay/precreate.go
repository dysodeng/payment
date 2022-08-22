package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// PreCreateResponse 预下单响应参数
type preCreateResponse struct {
	PayResponse
	Response preCreateResponseData `json:"alipay_trade_precreate_response"`
}

// preCreateResponseData 预下单响应参数数据
type preCreateResponseData struct {
	PayResponseData
	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
}

// PreCreate 预下单接口（当面付-生成二维码）
// @params bizContent string 业务数据
func (pay *AliPay) PreCreate(bizContent map[string]interface{}) (outTradeNo, qrCode string, err error) {
	biz, err := json.Marshal(bizContent)
	if err != nil {
		return "", "", err
	}

	response, err := pay.call("alipay.trade.precreate", string(biz))
	if err != nil {
		return
	}

	var result preCreateResponse
	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		return
	}

	if result.Response.Code != "10000" {
		err = errors.New(
			fmt.Sprintf("msg: %s, code:%s, sub_code:%s, sub_msg:%s",
				result.Response.Msg,
				result.Response.Code,
				result.Response.SubCode,
				result.Response.SubMsg),
		)
		return
	}

	return result.Response.OutTradeNo, result.Response.QrCode, nil
}
