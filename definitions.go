package tron_rpc_api

type APIInternalError struct {
	Errors       []string `json:"errors,omitempty"`
	ErrorMessage string   `json:"message,omitempty"`
	ErrorName    string   `json:"name,omitempty"`
}

type AddressInfo struct {
	APIInternalError
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
	HexAddress string `json:"hexAddress"`
}

type AccountInfo struct {
	APIInternalError
	Address               string `json:"address"`
	Balance               string `json:"balance"`
	CreateTime            string `json:"create_time"`
	LatestOprationTime    string `json:"latest_opration_time"`
	LatestConsumeFreeTime string `json:"latest_consume_free_time"`
}
