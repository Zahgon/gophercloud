package users

import "github.com/gophercloud/gophercloud/v2"

const (
	tenantPath = "tenants"
	userPath   = "users"
	rolePath   = "roles"
)

func ResourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func listRolesURL(c *gophercloud.ServiceClient, tenantID, userID string) string {
	_ = "STUB: not implemented"
	return ""
}
