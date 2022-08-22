package cryption

import (
	"crypto/sha512"
	"encoding/hex"
)

func SHA512(s string) (h2 string) {
	h := sha512.New()
	h.Write([]byte(s))
	h2 = hex.EncodeToString(h.Sum(nil))
	return
}
