package tron_rpc_api

import "testing"

func TestClient_CreateAccount(t *testing.T) {
	var tc = NewTron()
	addr, err := tc.CreateAccount()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(addr)
}

func TestTron_GetAccount(t *testing.T) {
	var tc = NewTron()
	tc.SetAddress("TQ8zLAj2jmi33799zi8H3ACCfRqXsu5U7w")
	addr, err := tc.GetAccount("")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(addr)
}
