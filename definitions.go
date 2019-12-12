package tron_rpc_api

type APIInternalError struct {
	Errors       []string `json:"errors,omitempty"`
	ErrorMessage string   `json:"message,omitempty"`
	ErrorName    string   `json:"name,omitempty"`
}

type AddressInfo struct {
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
	HexAddress string `json:"hexAddress"`
}

type AccountInfo struct {
	Address               string  `json:"address"`
	Balance               float64 `json:"balance"`
	CreateTime            int64   `json:"create_time"`
	LatestOprationTime    int64   `json:"latest_opration_time"`
	LatestConsumeFreeTime int64   `json:"latest_consume_free_time"`
}

type AccountBalance struct {
	Balance float64 `json:"balance"`
}

type AccountResources struct {
}

type TransactionRet struct {
	Data    []Transaction `json:"data"`
	Success bool          `json:"success"`
}

type Transaction struct {
	BlockTimeStamp float64  `json:"block_timestamp"`
	RawDataHex     string   `json:"raw_data_hex"`
	TxID           string   `json:"txID"`
	Signature      []string `json:"signature"`
}

type TransactionInfo struct {
}
