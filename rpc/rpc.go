package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/avila-r/env"

	"github.com/avila-r/bitclient/failure"
	"github.com/avila-r/bitclient/logger"
)

// RPCClient struct represents a client for making RPC calls.
type RPCClient struct {
	URL            string         // The URL of the RPC server
	Authentication Authentication // Authentication method used to access the RPC server
	client         *http.Client   // HTTP client used to send requests
}

// Request struct represents the structure of an RPC request.
type Request struct {
	ID      ID      `json:"id"`      // ID of the request
	Version Version `json:"jsonrpc"` // JSON-RPC version
	Method  Method  `json:"method"`  // Method name to be called
	Params  Params  `json:"params"`  // Parameters to be passed to the method
}

// Response struct represents the structure of an RPC response.
type Response struct {
	ID     ID              `json:"id"`     // ID of the response, matches the request ID
	Error  any             `json:"error"`  // Error field, if any error occurred
	Result json.RawMessage `json:"result"` // Raw response data
}

var (
	// Client initializes the default RPCClient based on environment variables.
	Client = func() *RPCClient {
		rpcURL := env.Get("RPC_URL")              // Get RPC URL from environment
		rpcAuthType := env.Get("RPC_AUTH_TYPE")   // Get RPC authentication type
		rpcAuthLabel := env.Get("RPC_AUTH_LABEL") // Get RPC authentication label

		// If any of the required environment variables are missing, log a warning and return nil
		if rpcURL == "" || rpcAuthType == "" || rpcAuthLabel == "" {
			logger.Warnf("unable to initialize a default rpc.Client (RPC_URL, RPC_AUTH_TYPE and RPC_AUTH_LABEL must be provided)")
			return nil
		}

		// Return a new RPCClient initialized with environment values
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

// New creates and returns a new RPCClient. It validates the URL and authentication parameters.
func New(uri string, authentication Authentication) (*RPCClient, error) {
	// Validate URL
	if uri == "" {
		return nil, failure.Of("URL cannot be empty")
	}

	parsed, err := url.Parse(uri) // Parse the URI
	if err != nil || !strings.HasPrefix(parsed.Scheme, "http") {
		return nil, failure.Of("invalid URL: must be a valid HTTP/HTTPS URL")
	}

	// Validate the authentication details
	if err := authentication.Validate(); err != nil {
		return nil, err
	}

	// Return a new RPCClient instance if all validations pass
	return &RPCClient{
		URL:            uri,
		Authentication: authentication,
		client:         &http.Client{},
	}, nil
}

// Do sends an RPC request and returns the corresponding response or an error.
func (c *RPCClient) Do(request Request) (*Response, error) {
	// Serialize the request to JSON
	body, err := json.Marshal(request)
	if err != nil {
		logger.Debugf("Error serializing request: %v", err)
		return nil, failure.Of("failed to serialize request: %v", err.Error())
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(body))
	if err != nil {
		logger.Debugf("Error creating HTTP request: %v", err)
		return nil, failure.Of("failed to set up http request: %v", err.Error())
	}

	// Setup authentication headers
	if err := c.Authentication.Setup(req); err != nil {
		return nil, err
	}

	// Set the Content-Type header
	req.Header.Set(ContentTypeHeaderLabel, string(ContentTypeApplicationJson))

	// Send the HTTP request
	resp, err := c.client.Do(req)
	if err != nil {
		logger.Debugf("Error sending request: %v", err)
		return nil, failure.Of("failed to send http request: %v", err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Debugf("Error reading response: %v", err)
		return nil, failure.Of("failed to read http response: %v", err.Error())
	}

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		logger.Debugf("Server response error: %s", payload)
		return nil, failure.Of("server responded with status code %d: %s", resp.StatusCode, payload)
	}

	// Unmarshal the response payload into the Response struct
	response := Response{}
	if err := json.Unmarshal(payload, &response); err != nil {
		logger.Debugf("Error deserializing response: %v", err)
		return nil, failure.Of("failed to deserialize response: %v", err.Error())
	}

	// If the response contains an error, return it
	if response.Error != nil {
		logger.Debugf("RPC call error: %v", response.Error)
		return nil, failure.Of("%v", response.Error)
	}

	// Return the successfully unmarshaled response
	return &response, nil
}

// GetMemoryInfo retrieves memory usage information from the Bitcoin client.
//
// This function sends a JSON-RPC request using the "getmemoryinfo" procedure call.
// The response contains details about the node's memory usage, including allocation statistics
// and optionally malloc information.
//
// Parameters:
// - mode (optional, string): Can be one of the following:
//   - "stats" (default): Returns general memory usage statistics.
//   - "mallocinfo": Returns low-level malloc implementation details as an XML string.
//
// Returns:
// - *Json: The JSON-RPC response containing memory usage data.
// - error: An error if the mode is invalid or if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient rpc memory
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getmemoryinfo "stats"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getmemoryinfo", "params": ["stats"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getmemoryinfo",
//	  "params": ["stats"]
//	}
//
// JSON Response:
//
//	{
//	  "locked": {
//	    "used": 512,
//	    "free": 1024,
//	    "total": 1536,
//	    "locked": 768,
//	    "chunks_used": 3,
//	    "chunks_free": 5
//	  }
//	}
//
// Notes:
// - Ensure the Bitcoin node is running to process the RPC request.
// - The "mallocinfo" mode is useful for debugging memory allocation at a lower level.
func GetMemoryInfo(mode ...string) (*Json, error) {
	params := Params{}
	if len(mode) > 0 && (mode[0] == "stats" || mode[0] == "mallocinfo") {
		params = append(params, mode[0])
	}

	request := Request{
		ID:      Identifier,
		Version: Version2,
		Method:  MethodGetMemoryInfo,
		Params:  params,
	}

	return JsonResult(Client.Do(request))
}

// GetInfo retrieves general information about the Bitcoin client.
//
// This function sends a JSON-RPC request using the "getrpcinfo" procedure call.
// The response contains information about the node, including its version, protocol, and network.
//
// Returns:
// - *Json: The JSON-RPC response containing general node information.
// - error: An error if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient rpc info
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getrpcinfo
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getrpcinfo", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getrpcinfo",
//	  "params": []
//	}
//
// JSON Response:
//
//	{
//	  "active_commands": [
//	    {
//	      "method": "getinfo",
//	      "duration": 3
//	    }
//	  ],
//	  "logpath": "/path/to/debug.log"
//	}
//
// Notes:
// - Useful for debugging and monitoring RPC-related commands and logs.
func GetInfo() (*Json, error) {
	request := Request{
		ID:      Identifier,
		Version: Version2,
		Method:  MethodGetRpcInfo,
		Params:  NoParams,
	}

	return JsonResult(Client.Do(request))
}

// Help retrieves help information for a specific RPC command or a list of all commands.
//
// This function sends a JSON-RPC request using the "help" procedure call.
// If a command is specified, the response contains detailed help for that command.
//
// Parameters:
// - command (optional, string): The name of the RPC command for which to retrieve help.
//
// Returns:
// - string: A string containing help information for the requested command.
// - error: An error if the request fails or the command is invalid.
//
// Example Usage (Assuming "getinfo" is the target):
//
//   - Using Bitclient:
//     $ bitclient rpc help getinfo
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli help "getinfo"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "help", "params": ["getinfo"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "help",
//	  "params": ["getinfo"]
//	}
//
// Response Example:
//
//	"Returns a list of RPC commands or help for a single command."
//
// Notes:
// - The help information may vary depending on the version of the Bitcoin client.
func Help(command ...string) (string, error) {
	params := Params{}
	if len(command) > 0 {
		params = append(params, command[0])
	}

	request := Request{
		ID:      Identifier,
		Version: Version2,
		Method:  MethodHelp,
		Params:  params,
	}

	response, err := Client.Do(request)
	if response == nil || err != nil {
		return "", err
	}

	return string(response.Result), nil
}

// LoggingProcedure configures logging categories for the Bitcoin client.
//
// This function sends a JSON-RPC request using the "logging" procedure call.
// The response modifies the enabled and disabled logging categories.
//
// Parameters:
// - include ([]string): Categories to enable.
// - exclude ([]string): Categories to disable.
//
// Returns:
// - *Json: The JSON-RPC response containing the updated logging state.
// - error: An error if the request fails or the categories are invalid.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient rpc logging ["net", "http"] ["rpc"]
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli logging ["net", "http"] ["rpc"]
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "logging", "params": [["net", "http"], ["rpc"]]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "logging",
//	  "params": [["net", "http"], ["rpc"]]
//	}
//
// JSON Response:
//
//	{
//	  "active": ["net", "http"],
//	  "inactive": ["rpc", "db"]
//	}
//
// Notes:
// - Categories must be valid logging categories supported by the Bitcoin client.
var LoggingProcedure = func(include []string, exclude []string) (*Json, error) {
	params := Params{}
	if len(include) > 0 {
		params = append(params, include)
	}
	if len(exclude) > 0 {
		params = append(params, exclude)
	}

	request := Request{
		ID:      Identifier,
		Version: Version2,
		Method:  MethodLogging,
		Params:  params,
	}

	return JsonResult(Client.Do(request))
}

// GetLogging retrieves the current active and inactive logging categories from the Bitcoin client.
//
// Returns:
//   - A JSON object with "active" and "inactive" logging categories.
//   - Error: If the request to the Bitcoin client fails.
func GetLogging() (*Json, error) {
	return LoggingProcedure(nil, nil)
}

// SetLogging updates the logging configuration of the Bitcoin client.
//
// Parameters:
//   - logging (LoggingConfig): Includes categories to enable and exclude categories to disable.
//
// Returns:
//   - A JSON object reflecting the updated logging configuration.
//   - Error: If the request to the Bitcoin client fails or the parameters are invalid.
func SetLogging(logging LoggingConfig) (*Json, error) {
	return LoggingProcedure(logging.Include, logging.Exclude)
}
