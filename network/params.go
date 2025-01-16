package network

// Ban struct represents a ban operation on an IP or Subnet, specifying the subnet, the operation command,
// the duration of the ban, and whether the ban time is absolute (in UNIX epoch time).
type Ban struct {
	// Subnet defines the IP or Subnet to be banned or unbanned.
	// Can be in the form of a single IP or a network with a subnet mask (e.g., 192.168.0.0/24).
	Subnet string

	// Time specifies the ban duration in seconds. A value of 0 means the default time of 24 hours.
	// If the 'absolute' field is set to true, this value should be a UNIX timestamp.
	Time int

	// Absolute specifies whether the 'Time' field represents an absolute UNIX timestamp.
	// If set to true, 'Time' should be an absolute timestamp rather than a relative duration.
	Absolute bool
}
