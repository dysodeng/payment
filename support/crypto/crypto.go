package crypto

import (
	"crypto/sha256"

	"github.com/pkg/errors"
)

// Sha256 sha256信息摘要
func Sha256(content []byte) ([]byte, error) {
	h := sha256.New()
	n, err := h.Write(content)
	if err != nil {
		return nil, err
	}

	if n != len(content) {
		return nil, errors.New("write length error")
	}

	return h.Sum(nil), nil
}
