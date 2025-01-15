package rpc

const (
	MethodGetMemoryInfo Method = "getmemoryinfo" // Method to get memory usage info
	MethodGetRpcInfo    Method = "getrpcinfo"    // Method to get RPC connection information
	MethodHelp          Method = "help"          // Method to get help information for RPC methods
	MethodLogging       Method = "logging"       // Method to get or set logging information
)
