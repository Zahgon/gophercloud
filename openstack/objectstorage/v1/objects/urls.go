package objects

import (
	"github.com/gophercloud/gophercloud/v2"
)

// tempURL returns an unescaped virtual path to generate the HMAC signature.
// Names must not be URL-encoded in this case.
//
// See: https://docs.openstack.org/swift/latest/api/temporary_url_middleware.html#hmac-signature-for-temporary-urls
func tempURL(c *gophercloud.ServiceClient, container, object string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(c *gophercloud.ServiceClient, container string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func copyURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func deleteURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func downloadURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func updateURL(c *gophercloud.ServiceClient, container, object string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func bulkDeleteURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }
