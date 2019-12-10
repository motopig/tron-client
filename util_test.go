package tron_rpc_api

import "testing"

func TestAddress2HexString(t *testing.T) {
	addr := "TSjCUZDs94iXqGX16qFohv1bpxtcKdBuik"
	t.Log(Address2HexString(addr))
}

func TestHexString2Address(t *testing.T) {
	addr := "41b7d3e037f2cac232daeafb86852ae29c0415ef0187ee5d15"
	t.Log(HexString2Address(addr))
}
