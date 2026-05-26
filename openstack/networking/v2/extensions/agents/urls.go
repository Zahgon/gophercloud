package agents

import "github.com/gophercloud/gophercloud/v2"

const resourcePath = "agents"
const dhcpNetworksResourcePath = "dhcp-networks"
const l3RoutersResourcePath = "l3-routers"
const bgpSpeakersResourcePath = "bgp-drinstances"
const bgpDRAgentSpeakersResourcePath = "bgp-speakers"
const bgpDRAgentAgentResourcePath = "bgp-dragents"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func listURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func getURL(c *gophercloud.ServiceClient, id string) string { _ = "STUB: not implemented"; return "" }

func updateURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func dhcpNetworksURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func l3RoutersURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func listDHCPNetworksURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func listL3RoutersURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func scheduleDHCPNetworkURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func scheduleL3RouterURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func removeDHCPNetworkURL(c *gophercloud.ServiceClient, id string, networkID string) string {
	_ = "STUB: not implemented"
	return ""
}

func removeL3RouterURL(c *gophercloud.ServiceClient, id string, routerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/agents/{agent-id}/bgp-drinstances
func listBGPSpeakersURL(c *gophercloud.ServiceClient, agentID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/agents/{agent-id}/bgp-drinstances
func scheduleBGPSpeakersURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/agents/{agent-id}/bgp-drinstances/{bgp-speaker-id}
func removeBGPSpeakersURL(c *gophercloud.ServiceClient, agentID string, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/bgp-dragents
func listDRAgentHostingBGPSpeakersURL(c *gophercloud.ServiceClient, speakerID string) string {
	_ = "STUB: not implemented"
	return ""
}
