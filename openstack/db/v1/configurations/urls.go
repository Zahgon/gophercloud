package configurations

import "github.com/gophercloud/gophercloud/v2"

func baseURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, configID string) string {
	_ = "STUB: not implemented"
	return ""
}

func instancesURL(c *gophercloud.ServiceClient, configID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listDSParamsURL(c *gophercloud.ServiceClient, datastoreID, versionID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getDSParamURL(c *gophercloud.ServiceClient, datastoreID, versionID, paramID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listGlobalParamsURL(c *gophercloud.ServiceClient, versionID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getGlobalParamURL(c *gophercloud.ServiceClient, versionID, paramID string) string {
	_ = "STUB: not implemented"
	return ""
}
