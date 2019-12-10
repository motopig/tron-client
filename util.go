package tron_rpc_api

import (
	"encoding/hex"
	"github.com/mr-tron/base58"
	"strings"
)

func Address2HexString(str string) (string, error) {
	if len(str) == 42 && strings.HasPrefix(str, "41") {
		return str, nil
	}
	addr, err := base58.Decode(str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(addr), nil
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
