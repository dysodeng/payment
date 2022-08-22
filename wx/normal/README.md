微信支付普通模式
==========

Usage
-----

```go
package main

import (
	"context"
	"github.com/dysodeng/layout/support/payment/wx/normal"
	"log"
)

func main() {
	payment := normal.NewPayment(context.Background(), normal.PaymentConfig{
		MchCertificateSerialNumber: "",
		MchAPIv3Key:                "",
		MchPrivateKey:              "",
		MchID:                      "",
		AppID:                      "",
	})
	native := payment.Native()
	resp, result, err := native.Prepay("测试支付", "201211111111", 1, "attach", "https://callback")
	if err != nil {
		log.Printf("%+v", err)
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
    }
}

```