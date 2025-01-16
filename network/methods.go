package network

import "github.com/avila-r/bitclient/rpc"

const (
	MethodAddNode            rpc.Method = "addnode"            // Method to add a node to the network
	MethodClearBanned        rpc.Method = "clearbanned"        // Method to clear banned nodes
	MethodDisconnectNode     rpc.Method = "disconnectnode"     // Method to disconnect a node
	MethodGetAddedNodeInfo   rpc.Method = "getaddednodeinfo"   // Method to get information about added nodes
	MethodGetConnectionCount rpc.Method = "getconnectioncount" // Method to get the number of connections
	MethodGetNetTotals       rpc.Method = "getnettotals"       // Method to get network totals
	MethodGetNetworkInfo     rpc.Method = "getnetworkinfo"     // Method to get network info
	MethodGetNodeAddresses   rpc.Method = "getnodeaddresses"   // Method to get node addresses
	MethodGetPeerInfo        rpc.Method = "getpeerinfo"        // Method to get peer info
	MethodListBanned         rpc.Method = "listbanned"         // Method to list banned nodes
	MethodPing               rpc.Method = "ping"               // Method to ping the network
	MethodSetBan             rpc.Method = "setban"             // Method to ban/unban a node
	MethodSetNetworkActive   rpc.Method = "setnetworkactive"   // Method to activate/deactivate network
)
