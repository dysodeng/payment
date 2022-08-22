微信支付服务商模式
==========

Usage
-----

```go
package main

import (
	"context"
	"github.com/dysodeng/layout/support/payment/wx/partner"
	"log"
)

func main() {
	payment := partner.NewPayment(context.Background(), partner.PaymentConfig{
		MchCertificateSerialNumber: "",
		MchAPIv3Key:                "",
		MchPrivateKey:              "",
		MchID:                      "",
		AppID:                      "",
	})
	native := payment.Native("supAppID", "subMchID")
	resp, result, err := native.Prepay("测试支付", "201211111111", 1, "attach", "https://callback")
	if err != nil {
		log.Printf("%+v", err)
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
    }
}

```
