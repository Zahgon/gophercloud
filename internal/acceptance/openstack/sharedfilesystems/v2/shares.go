package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/shares"
)

// CreateShare will create a share with a name, and a size of 1Gb. An
// error will be returned if the share could not be created
func CreateShare(t *testing.T, client *gophercloud.ServiceClient, optShareType ...string) (*shares.Share, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListShares lists all shares that belong to this tenant's project.
// An error will be returned if the shares could not be listed..
func ListShares(t *testing.T, client *gophercloud.ServiceClient) ([]shares.Share, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GrantAccess will grant access to an existing share. A fatal error will occur if
// this operation fails.
func GrantAccess(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share) (*shares.AccessRight, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RevokeAccess will revoke an exisiting access of a share. A fatal error will occur
// if this operation fails.
func RevokeAccess(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share, accessRight *shares.AccessRight) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAccessRightsSlice will retrieve all access rules assigned to a share.
// A fatal error will occur if this operation fails.
func GetAccessRightsSlice(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share) ([]shares.AccessRight, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteShare will delete a share. A fatal error will occur if the share
// failed to be deleted. This works best when used as a deferred function.
func DeleteShare(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share) {
	_ = "STUB: not implemented"
	return
}

// ExtendShare extends the capacity of an existing share
func ExtendShare(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share, newSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// ShrinkShare shrinks the capacity of an existing share
func ShrinkShare(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share, newSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func PrintMessages(t *testing.T, c *gophercloud.ServiceClient, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForStatus(t *testing.T, c *gophercloud.ServiceClient, id, status string) (*shares.Share, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
