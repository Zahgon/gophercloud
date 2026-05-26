package stackresources

import "github.com/gophercloud/gophercloud/v2"

func findURL(c *gophercloud.ServiceClient, stackName string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(c *gophercloud.ServiceClient, stackName, stackID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getURL(c *gophercloud.ServiceClient, stackName, stackID, resourceName string) string {
	_ = "STUB: not implemented"
	return ""
}

func metadataURL(c *gophercloud.ServiceClient, stackName, stackID, resourceName string) string {
	_ = "STUB: not implemented"
	return ""
}

func listTypesURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func schemaURL(c *gophercloud.ServiceClient, typeName string) string {
	_ = "STUB: not implemented"
	return ""
}

func templateURL(c *gophercloud.ServiceClient, typeName string) string {
	_ = "STUB: not implemented"
	return ""
}

func markUnhealthyURL(c *gophercloud.ServiceClient, stackName, stackID, resourceName string) string {
	_ = "STUB: not implemented"
	return ""
}
