package alipay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/dysodeng/payment/support"
	"github.com/dysodeng/payment/support/crypto/rsa"
)

// PayResponse 响应参数
type PayResponse struct {
	Sign string `json:"sign"`
}

// PayResponseData 响应参数数据
type PayResponseData struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

// gateway 获取支付宝网关地址
func (pay *AliPay) gateway() string {
	if pay.config.isDev {
		return devGateway
	}
	return prodGateway
}

// publicParams 公共参数
// @params method string 接口方法
// @params bizContent string 业务数据
func (pay *AliPay) publicParams(method, bizContent string) map[string]string {
	params := map[string]string{
		"app_id":      pay.config.appId,
		"method":      method,
		"format":      "json",
		"charset":     "utf-8",
		"sign_type":   "RSA2",
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"version":     "1.0",
		"biz_content": bizContent,
	}
	if pay.config.notifyUrl != "" {
		params["notify_url"] = pay.config.notifyUrl
	}
	if pay.config.returnUrl != "" {
		params["return_url"] = pay.config.returnUrl
	}
	if pay.config.appAuthToken != "" {
		params["app_auth_token"] = pay.config.appAuthToken
	}

	return params
}

// signString 组装支付签名参数
// @params params map[string]string 待签名参数
func (pay *AliPay) signString(params map[string]string) string {
	// 对key进行排序
	sortedKeys := make([]string, 0)
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// 组装签名参数
	var sign string
	for _, key := range sortedKeys {
		value := fmt.Sprintf("%v", params[key])
		if value != "" {
			sign += key + "=" + value + "&"
		}
	}

	return strings.TrimRight(sign, "&")
}

// privateKey 获取应用私钥
func (pay *AliPay) privateKey() string {
	prefix := "-----BEGIN PRIVATE KEY-----\n"
	suffix := "-----END PRIVATE KEY-----"
	return prefix + pay.config.privateKey + "\n" + suffix
}

// alipayPublicKey 获取支付宝公钥
func (pay *AliPay) alipayPublicKey() string {
	prefix := "-----BEGIN PUBLIC KEY-----\n"
	suffix := "-----END PUBLIC KEY-----"
	return prefix + pay.config.alipayPublicKey + "\n" + suffix
}

// call 接口调用
// @params method string 接口方法
// @params bizContent string 业务数据
func (pay *AliPay) call(method, bizContent string) ([]byte, error) {
	params := pay.publicParams(method, bizContent)
	content := pay.signString(params)

	sign, err := rsa.Encrypt(content, pay.privateKey())
	if err != nil {
		return nil, fmt.Errorf("%s 生成支付宝签名错误: %v", method, err)
	}
	params["sign"] = sign

	postValues := url.Values{}
	for key, value := range params {
		postValues.Add(key, value)
	}

	response, err := http.Post(
		pay.gateway(),
		"application/x-www-form-urlencoded;charset=utf-8",
		strings.NewReader(postValues.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("请求支付宝接口失败: %v", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return support.GbkToUtf8(body)
}
