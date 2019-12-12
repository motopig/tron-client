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
	t.Log(tc.GetAddress())
	addr, err := tc.GetAccount("")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(addr)
}

func TestTron_GetBalance(t *testing.T) {
	var tc = NewTron()
	tc.SetAddress("TQ8zLAj2jmi33799zi8H3ACCfRqXsu5U7w")

	balance, err := tc.GetBalance("")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(balance)
}

func TestTron_GetTransactionsFromAddress(t *testing.T) {
	var tc = NewTron()
	tc.SetAddress("TQ8zLAj2jmi33799zi8H3ACCfRqXsu5U7w")
	tr := tc.GetTransactionsFromAddress("", true, false, 100, 0)
	t.Log(tr)
}

func TestTron_GetTransactionsToAddress(t *testing.T) {
	var tc = NewTron()
	tc.SetAddress("TQ8zLAj2jmi33799zi8H3ACCfRqXsu5U7w")
	tr := tc.GetTransactionsToAddress("", true, false, 100, 0)
	t.Log(tr)
}
