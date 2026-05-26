package speakers

import "github.com/gophercloud/gophercloud/v2"

const urlBase = "bgp-speakers"

// return /v2.0/bgp-speakers/{bgp-speaker-id}
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers
func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

// return /v2.0/bgp-speakers/{bgp-speaker-id}
func getURL(c *gophercloud.ServiceClient, id string) string { _ = "STUB: not implemented"; return "" }

// return /v2.0/bgp-speakers
func listURL(c *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"

	// return /v2.0/bgp-speakers
	return ""
}

func createURL(c *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"

	// return /v2.0/bgp-speakers/{bgp-peer-id}
	return ""
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""

	// return /v2.0/bgp-speakers/{bgp-peer-id}
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""

	// return /v2.0/bgp-speakers/{bgp-speaker-id}/add_bgp_peer
}

func addBGPPeerURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/remove_bgp_peer
func removeBGPPeerURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/get_advertised_routes
func getAdvertisedRoutesURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/add_gateway_network
func addGatewayNetworkURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/remove_gateway_network
func removeGatewayNetworkURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}
