package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/shareaccessrules"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/shares"
)

func ShareAccessRuleGet(t *testing.T, client *gophercloud.ServiceClient, accessID string) (*shareaccessrules.ShareAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccessRightToShareAccess is a helper function that converts
// shares.AccessRight into shareaccessrules.ShareAccess struct.
func AccessRightToShareAccess(accessRight *shares.AccessRight) *shareaccessrules.ShareAccess {
	_ = "STUB: not implemented"
	return nil
}

func WaitForShareAccessRule(t *testing.T, client *gophercloud.ServiceClient, accessRule *shareaccessrules.ShareAccess, status string) error {
	_ = "STUB: not implemented"
	return nil
}

func ShareAccessRuleList(t *testing.T, client *gophercloud.ServiceClient, shareID string) ([]shareaccessrules.ShareAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
