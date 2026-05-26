package tasks

import (
	"github.com/gophercloud/gophercloud/v2"
)

const resourcePath = "tasks"

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, taskID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func getURL(c *gophercloud.ServiceClient, taskID string) string {
	_ = "STUB: not implemented"
	return ""
}

func createURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func nextPageURL(endpointURL, requestedNext string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
