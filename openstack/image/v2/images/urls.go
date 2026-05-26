package images

import (
	"github.com/gophercloud/gophercloud/v2"
)

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of images in the service `c`.
func listURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func createURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

// `imageURL(c,i)` is the URL for the image identified by ID `i` in
// the service `c`.
func imageURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

// `getURL(c,i)` is a URL for which a GET request will respond with
// information about the image identified by ID `i` in the service
// `c`.
func getURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

func updateURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

// builds next page full url based on current url
func nextPageURL(endpointURL, requestedNext string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
