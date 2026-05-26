package roles

import "github.com/gophercloud/gophercloud/v2"

const (
	ExtPath  = "OS-KSADM"
	RolePath = "roles"
	UserPath = "users"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func userTenantRoleURL(c *gophercloud.ServiceClient, tenantID, userID, roleID string) string {
	_ = "STUB: not implemented"
	return ""
}
