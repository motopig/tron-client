package tron_rpc_api

import "testing"

func TestAddress2HexString(t *testing.T) {
	addr := "TQ8zLAj2jmi33799zi8H3ACCfRqXsu5U7w"
	hex, _ := Address2HexString(addr, 0, 3, true)
	t.Log(hex)
}

func TestHexString2Address(t *testing.T) {
	addr := "41b7d3e037f2cac232daeafb86852ae29c0415ef0187ee5d15"
	t.Log(HexString2Address(addr))
}

func TestSubstr(t *testing.T) {
	str := "abcdefghigk"
	t.Log(Substr(str, 0, len(str)-3))
}
