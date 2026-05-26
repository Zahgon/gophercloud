package resourcetypes

import "github.com/gophercloud/gophercloud/v2"

const (
	resTypesPath = "resource_types"
)

func listURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func getSchemaURL(c *gophercloud.ServiceClient, resourceType string) string {
	_ = "STUB: not implemented"
	return ""
}

func generateTemplateURL(c *gophercloud.ServiceClient, resourceType string) string {
	_ = "STUB: not implemented"
	return ""
}
