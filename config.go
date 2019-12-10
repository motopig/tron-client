package tron_rpc_api

import (
	"net"
	"net/http"
	"time"
)

const (
	// node
	FullNodeEndpoint     = "https://api.trongrid.io"
	SolidityNodeEndpoint = "https://api.trongrid.io"
	EventServerEndpoint  = "https://api.trongrid.io"

	// status page
	FullNodeStatusPage     = "wallet/getnowblock"
	SolidityNodeStatusPage = "walletsolidity/getnowblock"
	EventServerStatusPage  = "healthcheck"
	ExplorerStatusPage     = "api/system/status"

	// ConnectionExponentFactor backoff exponent factor
	ConnectionExponentFactor float64 = 2.0

	// ConnectionInitialTimeout initial timeout
	ConnectionInitialTimeout = 2 * time.Millisecond

	// ConnectionMaximumJitterInterval jitter interval
	ConnectionMaximumJitterInterval = 2 * time.Millisecond

	// ConnectionMaxTimeout max timeout
	ConnectionMaxTimeout = 10 * time.Millisecond

	// ConnectionRetryCount retry count
	ConnectionRetryCount int = 2

	// ConnectionWithHTTPTimeout with http timeout
	ConnectionWithHTTPTimeout = 10 * time.Second

	// ConnectionTLSHandshakeTimeout tls handshake timeout
	ConnectionTLSHandshakeTimeout = 5 * time.Second

	// ConnectionMaxIdleConnections max idle http connections
	ConnectionMaxIdleConnections int = 10

	// ConnectionIdleTimeout idle connection timeout
	ConnectionIdleTimeout = 20 * time.Second

	// ConnectionExpectContinueTimeout expect continue timeout
	ConnectionExpectContinueTimeout = 3 * time.Second

	// ConnectionDialerTimeout dialer timeout
	ConnectionDialerTimeout = 5 * time.Second

	// ConnectionDialerKeepAlive keep alive
	ConnectionDialerKeepAlive = 20 * time.Second
)

// HTTP and Dialer connection variables
var (
	// _Dialer net dialer for ClientDefaultTransport
	_Dialer = &net.Dialer{
		KeepAlive: ConnectionDialerKeepAlive,
		Timeout:   ConnectionDialerTimeout,
	}

	// ClientDefaultTransport is the default transport struct for the HTTP client
	ClientDefaultTransport = &http.Transport{
		DialContext:           _Dialer.DialContext,
		ExpectContinueTimeout: ConnectionExpectContinueTimeout,
		IdleConnTimeout:       ConnectionIdleTimeout,
		MaxIdleConns:          ConnectionMaxIdleConnections,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   ConnectionTLSHandshakeTimeout,
	}
)
