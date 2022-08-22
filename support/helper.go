package support

import (
	"bytes"
	"io/ioutil"
	"math/rand"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	letterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// RandStringBytesMask 生成随机字符串
// @param int length 生成字符串长度
// @return string
func RandStringBytesMask(length int) string {

	str := make([]byte, length)

	for i, cache, reMain := length-1, rand.Int63(), letterIdxMax; i >= 0; {
		if reMain == 0 {
			cache, reMain = rand.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			str[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		reMain--
	}

	return string(str)
}
