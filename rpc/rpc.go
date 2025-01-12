package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/avila-r/env"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/logger"
)

type RPCClient struct {
	URL            string
	Authentication Authentication
	client         *http.Client
}

type Request struct {
	ID      ID      `json:"id"`
	Version Version `json:"jsonrpc"`
	Method  Method  `json:"method"`
	Params  Params  `json:"params"`
}

type Response struct {
	ID     ID              `json:"id"`
	Error  any             `json:"error"`
	Result json.RawMessage `json:"result"`
}

type Json map[string]interface{}

func (j Json) ToString() string {
	data, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		logger.Debugf("Failed to serialize json as string: %v", err)
	}

	return string(data)
}

var (
	Client = func() *RPCClient {
		rpcURL := env.Get("RPC_URL")
		rpcAuthType := env.Get("RPC_AUTH_TYPE")
		rpcAuthLabel := env.Get("RPC_AUTH_LABEL")

		if rpcURL == "" || rpcAuthType == "" || rpcAuthLabel == "" {
			logger.Warnf("unable to initialize a default rpc.Client (RPC_URL, RPC_AUTH_TYPE and RPC_AUTH_LABEL must be provided)")
			return nil
		}

		return &RPCClient{
			client: &http.Client{},
			URL:    rpcURL,
			Authentication: Authentication{
				Type:  AuthenticationType(rpcAuthType),
				Label: rpcAuthLabel,
			},
		}
	}()
)

func New(uri string, authentication Authentication) (*RPCClient, error) {
	// Validate URL
	if uri == "" {
		return nil, errs.Of("URL cannot be empty")
	}

	parsed, err := url.Parse(uri)
	if err != nil || !strings.HasPrefix(parsed.Scheme, "http") {
		return nil, errs.Of("invalid URL: must be a valid HTTP/HTTPS URL")
	}

	if err := authentication.Validate(); err != nil {
		return nil, err
	}

	return &RPCClient{
		URL:            uri,
		Authentication: authentication,
		client:         &http.Client{},
	}, nil
}

func (c *RPCClient) Do(request Request) (*Json, error) {
	// Serialize the request to JSON
	body, err := json.Marshal(request)
	if err != nil {
		logger.Debugf("Error serializing request: %v", err)
		return nil, errs.Of("failed to serialize request: %v", err.Error())
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(body))
	if err != nil {
		logger.Debugf("Error creating HTTP request: %v", err)
		return nil, errs.Of("failed to set up http request: %v", err.Error())
	}

	// Setup authentication
	if err := c.Authentication.Setup(req); err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set(ContentTypeHeaderLabel, string(ContentTypeApplicationJson))

	resp, err := c.client.Do(req)
	if err != nil {
		logger.Debugf("Error sending request: %v", err)
		return nil, errs.Of("failed to send http request: %v", err.Error())
	}
	defer resp.Body.Close()

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Debugf("Error reading response: %v", err)
		return nil, errs.Of("failed to read http response: %v", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		logger.Debugf("Server response error: %s", payload)
		return nil, errs.Of("server responded with status code %d: %s", resp.StatusCode, payload)
	}

	response := Response{}
	if err := json.Unmarshal(payload, &response); err != nil {
		logger.Debugf("Error deserializing response: %v", err)
		return nil, errs.Of("failed to deserialize response: %v", err.Error())
	}

	if response.Error != nil {
		logger.Debugf("RPC call error: %v", response.Error)
		return nil, errs.Of("%v", response.Error)
	}

	// Display the result
	result := Json{}
	if err := json.Unmarshal(response.Result, &result); err != nil {
		logger.Debugf("Error processing result: %v", err)
		return nil, errs.Of("failed to process result: %v", err.Error())
	}

	return &result, nil
}
