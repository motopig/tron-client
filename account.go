package tron_rpc_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *Tron) CreateAccount() (addressInfo *AddressInfo, err error) {
	// Create the request
	var resp string

	resp, err = t.Client.Request("wallet/generateaddress", http.MethodPost, nil)
	if err != nil {
		return
	}

	// Process the response
	addressInfo = new(AddressInfo)
	if err = json.Unmarshal([]byte(resp), addressInfo); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "addressInfo") // todo
		return
	}

	return
}

func (t *Tron) GetAccount(address string) (accountInfo *AccountInfo, err error) {
	// Create the request
	var resp string
	if address == "" {
		address = t.Address["hex"]
	} else {
		address = ToHex(address)
	}

	resp, err = t.Client.Request("walletsolidity/getaccount", http.MethodPost, []byte(fmt.Sprintf(`{"address":"%s"}`, address)))
	if err != nil {
		return
	}

	// Process the response
	accountInfo = new(AccountInfo)
	if err = json.Unmarshal([]byte(resp), accountInfo); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "accountInfo")
		return
	}

	return
}

func (t *Tron) GetBalance(address string) (accountBalance *AccountBalance, err error) {
	accountInfo, err := t.GetAccount(address)
	if err != nil {
		err = fmt.Errorf("error: %s", "accountInfo")
		return
	}
	accountBalance = new(AccountBalance)
	accountBalance.Balance = accountInfo.Balance
	return
}

func (t *Tron) GetAccountResources(address string) (accountResources *AccountResources, err error) {

	// Create the request
	var resp string
	if address == "" {
		address = t.Address["hex"]
	} else {
		address = ToHex(address)
	}

	resp, err = t.Client.Request("/wallet/getaccountresource", http.MethodPost, []byte(fmt.Sprintf(`{"address":"%s"}`, address)))
	if err != nil {
		return
	}

	// Process the response
	accountResources = new(AccountResources)
	if err = json.Unmarshal([]byte(resp), accountResources); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "accountResources")
		return
	}

	return
}
