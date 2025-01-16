package network

import (
	"strconv"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/rpc"
)

// ConnectToNode attempts to establish a one-time connection to a specified node.
//
// This function sends a JSON-RPC request using the "addnode" procedure call with the "onetry" command.
// The node specified will be attempted for a one-time connection.
//
// Parameters:
// - node (string): The address of the node to connect to (host:port).
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes connect "192.168.0.6:8333"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli addnode "192.168.0.6:8333" "onetry"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "addnode", "params": ["192.168.0.6:8333", "onetry"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "addnode",
//	  "params": ["192.168.0.6:8333", "onetry"]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
// - This method is used for attempting to connect to a node once, and is often used for troubleshooting or specific network scenarios.
func ConnectToNode(node string) error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodAddNode,
		Params:  rpc.Params{node, "onetry"},
	}

	_, err := rpc.Client.Do(request)

	return err
}

// AddNode attempts to add a node to the "addnode" list, ensuring that it is not disconnected due to DoS protections.
//
// This function sends a JSON-RPC request using the "addnode" procedure call with the "add" command,
// which adds the node to the addnode list for permanent connection.
//
// Parameters:
// - node (string): The address of the node to add to the list (host:port).
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes add "192.168.0.6:8333"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli addnode "192.168.0.6:8333" "add"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "addnode", "params": ["192.168.0.6:8333", "add"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "addnode",
//	  "params": ["192.168.0.6:8333", "add"]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
// - The node added using this method will be protected from DoS disconnection and can be used for long-term connections.
func AddNode(node string) error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodAddNode,
		Params:  rpc.Params{node, "add"},
	}

	_, err := rpc.Client.Do(request)

	return err
}

// RemoveNode removes a node from the "addnode" list.
//
// This function sends a JSON-RPC request using the "addnode" procedure call with the "remove" command,
// which removes the specified node from the addnode list.
//
// Parameters:
// - node (string): The address of the node to remove (host:port).
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes remove "192.168.0.6:8333"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli addnode "192.168.0.6:8333" "remove"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "addnode", "params": ["192.168.0.6:8333", "remove"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "addnode",
//	  "params": ["192.168.0.6:8333", "remove"]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
// - The node will be removed from the list and may be disconnected from the network.
func RemoveNode(node string) error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodAddNode,
		Params:  rpc.Params{node, "remove"},
	}

	_, err := rpc.Client.Do(request)

	return err
}

// ClearBanned clears all banned IPs from the node's banned list.
//
// This function sends a JSON-RPC request using the "clearbanned" procedure call to remove all IPs
// from the node's banned list.
//
// Parameters:
// - None.
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes unban
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli clearbanned
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "clearbanned", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "clearbanned",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
// - This method removes all banned IP addresses from the list, allowing those IPs to reconnect.
func ClearBanned() error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodClearBanned,
		Params:  rpc.NoParams,
	}

	_, err := rpc.Client.Do(request)

	return err
}

// DisconnectNode disconnects from the specified peer node using either its address or node ID.
//
// This function sends a JSON-RPC request using the "disconnectnode" procedure call to disconnect
// from a peer node based on either its address or node ID.
//
// Parameters:
//   - node (string): The IP address/port of the node or the node ID to disconnect from.
//     If the argument is a valid numeric node ID, it will be used as the node ID.
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes disconnect "192.168.0.6:8333"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli disconnectnode "192.168.0.6:8333"
//     $ bitcoin-cli disconnectnode "" 1
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "disconnectnode", "params": ["192.168.0.6:8333"]}' \
//     -H 'content-type: text/plain;' {url}
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "disconnectnode", "params": ["", 1]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "disconnectnode",
//	  "params": ["192.168.0.6:8333"]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
//   - Strictly one of 'address' or 'nodeid' must be provided to identify the node.
//     If both are provided, only the valid argument will be used.
func DisconnectNode(node string) error {
	params := rpc.Params{}
	if _, err := strconv.Atoi(node); err != nil {
		// If 'node' is not a numeric ID, it is treated as an address.
		params = append(params, node)
	} else {
		// Otherwise, it is treated as a node ID.
		params = rpc.Params{"", node}
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodDisconnectNode,
		Params:  params,
	}

	_, err := rpc.Client.Do(request)

	return err
}

// InspectAddedNodes retrieves information about the added nodes in the Bitcoin network.
//
// This function sends a JSON-RPC request using the "getaddednodeinfo" procedure to return information
// about one or all added nodes (excluding one-time nodes).
//
// Parameters:
//   - node (string...): An optional parameter. If provided, the function will return information for the
//     specific node. If omitted, information for all added nodes will be returned.
//
// Returns:
// - *rpc.Array: The array of node information if the request is successful, or an error if the request fails.
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient nodes info "192.168.0.201"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getaddednodeinfo "192.168.0.201"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getaddednodeinfo", "params": ["192.168.0.201"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getaddednodeinfo",
//	  "params": ["192.168.0.201"]
//	}
//
// JSON Response Example:
//
//	[
//	  {
//	    "addednode": "192.168.0.201",
//	    "connected": true,
//	    "addresses": [
//	      {
//	        "address": "192.168.0.201:8333",
//	        "connected": "inbound"
//	      }
//	    ]
//	  }
//	]
//
// Notes:
//   - If no 'node' argument is provided, all added nodes are returned. If a 'node' is provided, only information
//     for that specific node is returned.
func InspectAddedNodes(node ...string) (*rpc.Array, error) {
	params := rpc.Params{}
	if len(node) > 0 {
		// If a node argument is provided, append it to the params.
		params = append(params, node[0])
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetAddedNodeInfo,
		Params:  params,
	}

	return rpc.ArrayResult(rpc.Client.Do(request))
}

// GetConnectionCount retrieves the number of connections to other nodes in the Bitcoin network.
//
// This function sends a JSON-RPC request using the "getconnectioncount" procedure to return the number of
// active connections to other nodes.
//
// Returns:
//   - *rpc.Response: The response containing the connection count if the request is successful, or an error
//     if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network connections
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getconnectioncount
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getconnectioncount", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getconnectioncount",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	{
//	  "result": 8, // The number of active connections to other nodes
//	  "error": null,
//	  "id": "curltest"
//	}
//
// Notes:
// - This method returns the total number of connections to other nodes.
func GetConnectionCount() (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetConnectionCount,
		Params:  rpc.NoParams,
	}

	return rpc.Client.Do(request)
}

// InspectTraffic retrieves the network traffic statistics including total bytes received,
// total bytes sent, and current time information.
//
// This function sends a JSON-RPC request using the "getnettotals" procedure to return information
// about the network traffic, including the amount of data transferred and the time-related metrics.
//
// Returns:
//   - *rpc.Json: The response containing network traffic information if the request is successful, or an error
//     if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network traffic
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getnettotals
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnettotals", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getnettotals",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	{
//	  "result": {
//	    "totalbytesrecv": 1234567890,
//	    "totalbytessent": 987654321,
//	    "timemillis": 1612204200000,
//	    "uploadtarget": {
//	      "timeframe": 3600,
//	      "target": 1000000000,
//	      "target_reached": true,
//	      "serve_historical_blocks": true,
//	      "bytes_left_in_cycle": 12345,
//	      "time_left_in_cycle": 60
//	    }
//	  },
//	  "error": null,
//	  "id": "curltest"
//	}
//
// Notes:
// - This method provides total bytes sent and received, as well as data about the upload target and remaining cycle.
func InspectTraffic() (*rpc.Json, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetNetTotals,
		Params:  rpc.NoParams,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}

// GetNetworkInfo retrieves various state information regarding P2P networking.
//
// This function sends a JSON-RPC request using the "getnetworkinfo" procedure call.
// The response contains details about the node's network status, connections, fees, and warnings.
//
// Returns:
// - *rpc.Json: The JSON-RPC response containing the network information.
// - error: An error if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network info
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getnetworkinfo
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnetworkinfo", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getnetworkinfo",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	{
//	  "version": 230000,
//	  "subversion": "/Satoshi:23.0.0/",
//	  "protocolversion": 70015,
//	  "localservices": "0000000000000409",
//	  "localservicesnames": ["NETWORK", "WITNESS"],
//	  "localrelay": true,
//	  "timeoffset": 0,
//	  "connections": 8,
//	  "connections_in": 2,
//	  "connections_out": 6,
//	  "networkactive": true,
//	  "networks": [
//	    {
//	      "name": "ipv4",
//	      "limited": false,
//	      "reachable": true,
//	      "proxy": "",
//	      "proxy_randomize_credentials": false
//	    },
//	    {
//	      "name": "ipv6",
//	      "limited": false,
//	      "reachable": true,
//	      "proxy": "",
//	      "proxy_randomize_credentials": false
//	    }
//	  ],
//	  "relayfee": 0.00001000,
//	  "incrementalfee": 0.00001000,
//	  "localaddresses": [
//	    {
//	      "address": "192.0.2.1",
//	      "port": 8333,
//	      "score": 1
//	    }
//	  ],
//	  "warnings": ""
//	}
//
// Notes:
// - This command is useful for monitoring the network state, including connections and fees.
// - Check the "warnings" field for any network or blockchain-related alerts.
func GetNetworkInfo() (*rpc.Json, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetNetworkInfo,
		Params:  rpc.NoParams,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}

// FindAddresses retrieves known addresses that can potentially be used to find new nodes in the network.
// It returns up to a specified number of addresses, or all known addresses if 0 is specified.
//
// Arguments:
// - max: Optional. The maximum number of addresses to return. If not provided, defaults to 1.
//
// Returns:
// - *rpc.Array: A JSON array of addresses, including their last seen time, services, address, and port.
//
// Example Usage:
//
//   - Using Bitcoin CLI:
//     $ bitcoin-cli nodes find
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getnodeaddresses", "params": [8]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getnodeaddresses",
//	  "params": [8]
//	}
//
// JSON Response Example:
//
//	{
//	  "result": [
//	    {
//	      "time": 1612204200000,
//	      "services": 1,
//	      "address": "192.168.0.1",
//	      "port": 8333
//	    },
//	    {
//	      "time": 1612204250000,
//	      "services": 1,
//	      "address": "192.168.0.2",
//	      "port": 8333
//	    }
//	  ],
//	  "error": null,
//	  "id": "curltest"
//	}
//
// Notes:
// - Use `max` to limit the number of addresses returned. If `max` is 0, all known addresses will be returned.
func FindAddresses(max ...int) (*rpc.Array, error) {
	params := rpc.Params{}
	if len(max) > 0 {
		params = append(params, max[0])
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetNodeAddresses,
		Params:  params,
	}

	return rpc.ArrayResult(rpc.Client.Do(request))
}

// GetPeers retrieves data about each connected network node.
//
// This function sends a JSON-RPC request using the "getpeerinfo" procedure call.
// The response is a JSON array of objects, each containing detailed information
// about a peer connected to the node.
//
// Returns:
// - *rpc.Array: The JSON-RPC response containing information about connected peers.
// - error: An error if the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network peers
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getpeerinfo
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getpeerinfo", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getpeerinfo",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	[
//	  {
//	    "id": 1,
//	    "addr": "192.0.2.1:8333",
//	    "addrbind": "192.0.2.2:8333",
//	    "addrlocal": "192.0.2.3:8333",
//	    "network": "ipv4",
//	    "mapped_as": 12345,
//	    "services": "0000000000000409",
//	    "servicesnames": ["NETWORK", "WITNESS"],
//	    "relaytxes": true,
//	    "lastsend": 1673539200,
//	    "lastrecv": 1673539201,
//	    "last_transaction": 1673539202,
//	    "last_block": 1673539203,
//	    "bytessent": 123456,
//	    "bytesrecv": 654321,
//	    "conntime": 1673539100,
//	    "timeoffset": 0,
//	    "pingtime": 0.001,
//	    "minping": 0.0005,
//	    "pingwait": 0.002,
//	    "version": 70015,
//	    "subver": "/Satoshi:23.0.0/",
//	    "inbound": false,
//	    "connection_type": "outbound-full-relay",
//	    "startingheight": 750000,
//	    "banscore": 0,
//	    "synced_headers": 750000,
//	    "synced_blocks": 750000,
//	    "inflight": [],
//	    "permissions": ["relay", "download"],
//	    "minfeefilter": 0.00001000,
//	    "bytessent_per_msg": {
//	      "inv": 12345,
//	      "tx": 67890
//	    },
//	    "bytesrecv_per_msg": {
//	      "block": 54321,
//	      "*other*": 98765
//	    }
//	  }
//	]
//
// Notes:
//   - This command is useful for analyzing the node's peer connections and behaviors.
//   - Deprecated fields such as "banscore", "whitelisted", and "addnode" may require
//     additional configuration options to be included in the response.
func GetPeers() (*rpc.Array, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetPeerInfo,
		Params:  rpc.NoParams,
	}

	return rpc.ArrayResult(rpc.Client.Do(request))
}

// ListBanned retrieves all manually banned IPs and subnets, including the time until the address is banned and when the ban was created.
//
// Returns:
// - *rpc.Array: A JSON array containing the banned addresses, including their banned time and creation time.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network blacklist
//
//   - Using Bitcoin CLI:
//     $ bitcoin-cli listbanned
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "listbanned", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "listbanned",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	{
//	  "result": [
//	    {
//	      "address": "192.168.0.1",
//	      "banned_until": 1612204200,
//	      "ban_created": 1612204100
//	    },
//	    {
//	      "address": "192.168.0.2",
//	      "banned_until": 1612204250,
//	      "ban_created": 1612204150
//	    }
//	  ],
//	  "error": null,
//	  "id": "curltest"
//	}
//
// Notes:
// - The `banned_until` field is the UNIX epoch time indicating when the ban will expire.
// - The `ban_created` field indicates the time the ban was created.
func ListBanned() (*rpc.Array, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodListBanned,
		Params:  rpc.NoParams,
	}

	return rpc.ArrayResult(rpc.Client.Do(request))
}

// Ping requests that a ping be sent to all other nodes to measure the ping time.
//
// This function sends a JSON-RPC request using the "ping" procedure call. The response is
// null, indicating the ping was successfully processed. The results of the ping are available
// in the `pingtime` and `pingwait` fields in the "getpeerinfo" command.
//
// Returns:
// - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient ping
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli ping
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "ping", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "ping",
//	  "params": []
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
// - The ping command measures processing backlog, not just network ping.
// - The results are available in the `pingtime` and `pingwait` fields of the `getpeerinfo` response.
func Ping() error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodPing,
		Params:  rpc.NoParams,
	}

	_, err := rpc.Client.Do(request)

	return err
}

// Health checks if the node is responsive by sending a ping and verifying no error occurs.
//
// Returns:
// - bool: True if the ping was successful (node is healthy), false otherwise.
func Health() bool {
	return Ping() == nil
}

// SetBan attempts to add a subnet/IP to the banned list.
//
// This function sends a JSON-RPC request using the "setban" procedure call. The ban operation
// can add a specified IP or subnet to the banned list. The ban can be temporary or permanent,
// depending on the provided `bantime` and `absolute` fields.
//
// Arguments:
//   - ban.Subnet (string): The IP or subnet to be banned. The subnet can be specified
//     as an IP address (e.g., "192.168.0.6") or a network with a subnet mask (e.g., "192.168.0.0/24").
//   - ban.Time (int): The duration (in seconds) for which the IP/subnet is banned. A value of 0
//     means to use the default ban time of 24 hours. If `Absolute` is set to true, this field should
//     be a UNIX timestamp representing the absolute ban time.
//   - ban.Absolute (bool): If true, the `Time` field is treated as an absolute timestamp (UNIX epoch time).
//     If false, `Time` is treated as a relative ban duration.
//
// Returns:
//   - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network ban "192.168.0.6" 86400
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli setban "192.168.0.6" "add" 86400
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "setban", "params": ["192.168.0.6", "add", 86400]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "setban",
//	  "params": ["192.168.0.6", "add", 86400]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
//   - A subnet can be specified in the form of an IP address with a subnet mask (e.g., "192.168.0.0/24").
//   - If `bantime` is set to 0, the default ban time of 24 hours will be used, unless overridden by
//     the `-bantime` argument during startup.
//   - If `absolute` is set to true, the `bantime` should be a UNIX timestamp indicating the absolute
//     time the ban should end.
func SetBan(ban Ban) error {
	if ban.Subnet == "" {
		return errs.Of("ban's subnet must be provided")
	}

	params := rpc.Params{
		ban.Subnet,
		"add",
		ban.Time,
		ban.Absolute,
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodSetBan,
		Params:  params,
	}

	_, err := rpc.Client.Do(request)

	return err
}

// Unban attempts to remove a subnet/IP from the banned list.
//
// This function sends a JSON-RPC request using the "setban" procedure call. The ban operation
// can remove a specified IP or subnet from the banned list. The ban can be temporary or permanent,
// depending on the provided `bantime` and `absolute` fields.
//
// Arguments:
//   - subnet (string): The IP or subnet to be banned. The subnet can be specified as an IP
//     address (e.g., "192.168.0.6") or a network with a subnet mask (e.g., "192.168.0.0/24").
//
// Returns:
//   - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network unban "192.168.0.6"
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli setban "192.168.0.6" "remove"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "setban", "params": ["192.168.0.6", "remove"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "setban",
//	  "params": ["192.168.0.6", "remove"]
//	}
//
// JSON Response Example:
//
//	null
//
// Notes:
//   - A subnet can be specified in the form of an IP address with a subnet mask (e.g., "192.168.0.0/24").
func Unban(subnet string) error {
	if subnet == "" {
		return errs.Of("ban's subnet must be provided")
	}

	params := rpc.Params{
		subnet,
		"remove",
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodSetBan,
		Params:  params,
	}

	_, err := rpc.Client.Do(request)

	return err
}

// SetNetworkActive enables or disables all P2P network activity.
//
// This function sends a JSON-RPC request using the "setnetworkactive" procedure call. It allows
// you to enable or disable the P2P network activity of the node. This is useful when you want to
// temporarily stop the node from connecting to peers or performing network-related tasks.
//
// Arguments:
//   - status (bool): A boolean value indicating whether to enable or disable network activity.
//     Pass `true` to enable networking, or `false` to disable it.
//
// Returns:
//   - error: An error if the request fails, otherwise nil.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient network activate|deactivate
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli setnetworkactive true
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "setnetworkactive", "params": [true]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "setnetworkactive",
//	  "params": [true]
//	}
//
// Notes:
//   - This command can be used to temporarily stop the node from making outbound connections or
//     responding to incoming connections.
func SetNetworkActive(status bool) error {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodSetNetworkActive,
		Params:  rpc.Params{status},
	}

	_, err := rpc.Client.Do(request)

	return err
}
