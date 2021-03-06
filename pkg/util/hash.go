package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
