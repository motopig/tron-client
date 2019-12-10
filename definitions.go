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
