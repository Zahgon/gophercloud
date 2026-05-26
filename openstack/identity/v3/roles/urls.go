package roles

import "github.com/gophercloud/gophercloud/v2"

const (
	rolePath = "roles"
)

func listURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func getURL(client *gophercloud.ServiceClient, roleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func createURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func updateURL(client *gophercloud.ServiceClient, roleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(client *gophercloud.ServiceClient, roleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"
	return ""
}

func listAssignmentsOnResourceURL(client *gophercloud.ServiceClient, targetType, targetID, actorType, actorID string) string {
	_ = "STUB: not implemented"
	return ""
}

func assignURL(client *gophercloud.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func createRoleInferenceRuleURL(client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getRoleInferenceRuleURL(client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func listRoleInferenceRulesURL(client *gophercloud.ServiceClient) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteRoleInferenceRuleURL(client *gophercloud.ServiceClient, priorRoleID, impliedRoleID string) string {
	_ = "STUB: not implemented"
	return ""
}
