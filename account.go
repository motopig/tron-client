package tron_rpc_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateAccount() (addressInfo *AddressInfo, err error) {
	// Create the request
	var resp string

	resp, err = c.Request("wallet/generateaddress", http.MethodPost, nil)
	if err != nil {
		return
	}

	// Process the response
	addressInfo = new(AddressInfo)
	if err = json.Unmarshal([]byte(resp), addressInfo); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", addressInfo.ErrorMessage)
		return
	}

	return
}
