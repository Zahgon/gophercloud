package peers

import "github.com/gophercloud/gophercloud/v2"

const urlBase = "bgp-peers"

// return /v2.0/bgp-peers/{bgp-peer-id}
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-peers
func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

// return /v2.0/bgp-peers/{bgp-peer-id}
func getURL(c *gophercloud.ServiceClient, id string) string { _ = "STUB: not implemented"; return "" }

// return /v2.0/bgp-peers
func listURL(c *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"

	// return /v2.0/bgp-peers
	return ""
}

func createURL(c *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"

	// return /v2.0/bgp-peers/{bgp-peer-id}
	return ""
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""

	// return /v2.0/bgp-peers/{bgp-peer-id}
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
