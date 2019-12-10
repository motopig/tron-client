package tron_rpc_api

import "testing"

func TestClient_CreateAccount(t *testing.T) {
	var c = NewClient()
	addr, err := c.CreateAccount()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(addr)
}
