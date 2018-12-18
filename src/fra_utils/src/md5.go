package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(message []byte) string {
	h := md5.New()
	h.Write(message) // 需要加密的字符串为 message
	return hex.EncodeToString(h.Sum(nil))
}

func GetMD5ByString(message string) string {
	h := md5.New()
	h.Write([]byte(message)) // 需要加密的字符串为 message
	return hex.EncodeToString(h.Sum(nil))
}
