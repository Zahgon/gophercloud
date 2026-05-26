package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/sharetypes"
)

// CreateShareType will create a share type with a random name. An
// error will be returned if the share type was unable to be created.
func CreateShareType(t *testing.T, client *gophercloud.ServiceClient) (*sharetypes.ShareType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteShareType will delete a share type. An error will occur if
// the share type was unable to be deleted.
func DeleteShareType(t *testing.T, client *gophercloud.ServiceClient, shareType *sharetypes.ShareType) {
	_ = "STUB: not implemented"
	return
}
