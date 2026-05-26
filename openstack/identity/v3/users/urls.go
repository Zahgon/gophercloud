package users

import "github.com/gophercloud/gophercloud/v2"

func listURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func getURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func createURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func updateURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func changePasswordURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listGroupsURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func addToGroupURL(client *gophercloud.ServiceClient, groupID, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func isMemberOfGroupURL(client *gophercloud.ServiceClient, groupID, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func removeFromGroupURL(client *gophercloud.ServiceClient, groupID, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listProjectsURL(client *gophercloud.ServiceClient, userID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listInGroupURL(client *gophercloud.ServiceClient, groupID string) string {
	_ = "STUB: not implemented"
	return ""
}
