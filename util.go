package tron_rpc_api

import (
	"encoding/hex"
	"github.com/mr-tron/base58"
	"strings"
)

func ToHex(str string) string {
	if strings.HasPrefix(str, "T") && len(str) == 34 {
		hex, _ := Address2HexString(str, 0, 3, true)
		return hex
	}
	return "" // todo
}

func Address2HexString(str string, removeLeadingBytes int, removeTrailingBytes int, removeCompression bool) (string, error) {
	if len(str) == 42 && strings.HasPrefix(str, "41") {
		return str, nil
	}
	addr, err := base58.Decode(str)
	if err != nil {
		return "", err
	}
	raw := hex.EncodeToString(addr)

	if removeLeadingBytes > 0 {
		raw = Substr(raw, 0, removeLeadingBytes*2)
	}
	if removeTrailingBytes > 0 {
		raw = Substr(raw, 0, (len(raw) - removeTrailingBytes*2))
	}

	if removeCompression {
		raw = Substr(raw, 0, (len(raw) - 2))
	}
	return raw, nil
}

func HexString2Address(str string) string {
	if !ctypeXdigit(str) {
		return str
	}
	if len(str) < 2 || (len(str)&1) != 0 {
		return ""
	}
	addr, _ := hex.DecodeString(str)
	return base58.Encode(addr)
}

func ctypeXdigit(str string) bool {
	dat := []byte(str)
	isHex := true
	for _, v := range dat {
		if v < 48 || v > 57 && v < 65 || v > 70 && v < 97 || v > 102 {
			isHex = false
			break
		}
	}
	return isHex
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
