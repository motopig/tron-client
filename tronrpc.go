package tron_rpc_api

import (
	"bytes"
	"fmt"
	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	// HTTPClient carries out the POST operations
	HTTPClient heimdall.Client
	// LastRequest is the raw information from the last request
	LastRequest *LastRequest
}

// LastRequest is used to track what was submitted to the Request()
type LastRequest struct {

	// Method is either POST or GET
	Method string

	// PostData is the post data submitted if POST request
	PostData string

	// StatusCode is the last code from the request
	StatusCode int

	// URL is the url used for the request
	URL string
}

func NewClient() *Client {
	c := new(Client)
	backOff := heimdall.NewExponentialBackoff(
		ConnectionInitialTimeout,
		ConnectionMaxTimeout,
		ConnectionExponentFactor,
		ConnectionMaximumJitterInterval,
	)

	c.HTTPClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(ConnectionWithHTTPTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(ConnectionRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: ClientDefaultTransport,
			Timeout:   ConnectionWithHTTPTimeout,
		}),
	)

	// Create a last request struct
	c.LastRequest = new(LastRequest)

	return c
}

func (c *Client) Request(endpoint string, method string, payload []byte) (response string, err error) {
	// Set reader
	var bodyReader io.Reader

	// Add the network value todo check request type
	endpoint = fmt.Sprintf("%s/%s", FullNodeEndpoint, endpoint)

	// Switch on Methods
	switch method {
	case http.MethodPost, http.MethodPut:
		{
			bodyReader = bytes.NewBuffer(payload)
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.URL = endpoint

	// Start the request
	var request *http.Request
	if request, err = http.NewRequest(method, endpoint, bodyReader); err != nil {
		return
	}

	// Set the content type on Method
	if method == http.MethodPost || method == http.MethodPut {
		request.Header.Set("Content-Type", "application/json")
	}

	// Fire the http request
	var resp *http.Response
	if resp, err = c.HTTPClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		if bodyErr := resp.Body.Close(); bodyErr != nil {
			log.Printf("error closing response body: %s", bodyErr.Error())
		}
	}()

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Parse the response
	response = string(body)

	// Done
	return
}
