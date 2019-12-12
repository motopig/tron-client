package tron_rpc_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *Tron) GetTransaction(transactionId string) (transaction *Transaction, err error) {
	// Create the request
	var resp string

	resp, err = t.Client.Request("wallet/gettransactionbyid", http.MethodPost, []byte(fmt.Sprintf(`{"value":"%s"}`, transactionId)))
	if err != nil {
		return
	}

	// Process the response
	transaction = new(Transaction)
	if err = json.Unmarshal([]byte(resp), transaction); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "transaction")
		return
	}

	return
}

func (t *Tron) GetTransactionInfo(transactionId string) (transactionInfo *TransactionInfo, err error) {
	// Create the request
	var resp string

	resp, err = t.Client.Request("wallet/gettransactionbyid", http.MethodPost, []byte(fmt.Sprintf(`{"value":"%s"}`, transactionId)))
	if err != nil {
		return
	}

	// Process the response
	transactionInfo = new(TransactionInfo)
	if err = json.Unmarshal([]byte(resp), transactionInfo); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "transactionInfo")
		return
	}

	return
}

func getTransactionsRelated(t *Tron, confirm bool, unconfirm bool, address string, direction string, limit int, offset int) (transactionRet *TransactionRet, err error) {
	if direction != "to" && direction != "from" {
		err = fmt.Errorf("error: %s", "direction")
		return transactionRet, err
	}

	// Create the request
	var resp string
	var from, to bool
	if direction == "to" {
		from = false
		to = true
	} else {
		to = false
		from = true
	}
	raw := `{"only_confirmed":%v,"only_unconfirmed":%v,"only_from":%v,"only_to":%v,"limit":%d,"offset":%d}`

	payload := []byte(fmt.Sprintf(raw, confirm, unconfirm, from, to, limit, offset))
	fmt.Println(string(payload))
	resp, err = t.Client.Request(fmt.Sprintf("/accounts/%s/transactions", address), http.MethodGet, payload)
	if err != nil {
		return
	}

	// Process the response
	transactionRet = new(TransactionRet)
	if err = json.Unmarshal([]byte(resp), transactionRet); err != nil {
		return
	}

	// Error from request?
	if t.Client.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", "transactions")
		return
	}

	return
}

func (t *Tron) GetTransactionsToAddress(address string, confirm bool, unconfirm bool, limit int, offset int) *TransactionRet {
	if address == "" {
		address = t.Address["hex"]
	} else {
		address = ToHex(address)
	}
	transactionRet, _ := getTransactionsRelated(t, confirm, unconfirm, address, "to", limit, offset)
	return transactionRet
}

func (t *Tron) GetTransactionsFromAddress(address string, confirm bool, unconfirm bool, limit int, offset int) *TransactionRet {
	if address == "" {
		address = t.Address["hex"]
	} else {
		address = ToHex(address)
	}
	transactionRet, _ := getTransactionsRelated(t, confirm, unconfirm, address, "from", limit, offset)
	return transactionRet
}
