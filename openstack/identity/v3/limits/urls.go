package limits

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath             = "limits"
	enforcementModelPath = "model"
)

func enforcementModelURL(client *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"
	return ""
}

func rootURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(client *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
