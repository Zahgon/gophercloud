package zones

import "github.com/gophercloud/gophercloud/v2"

// baseURL returns the base URL for zones.
func baseURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

// zoneURL returns the URL for a specific zone.
func zoneURL(c *gophercloud.ServiceClient, zoneID string) string {
	_ = "STUB: not implemented"
	return ""
}

// sharesBaseURL returns the URL for shared zones.
func sharesBaseURL(c *gophercloud.ServiceClient, zoneID string) string {
	_ = "STUB: not implemented"
	return ""
}

// shareURL returns the URL for a shared zone.
func shareURL(c *gophercloud.ServiceClient, zoneID, sharedZoneID string) string {
	_ = "STUB: not implemented"
	return ""
}
