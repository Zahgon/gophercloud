package taas

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/taas/tapmirrors"
)

// CreateTapMirror will create a Tap Mirror with the specified portID and remoteIP. An error
// will be returned if the Tap Mirror could not be created.
func CreateTapMirror(t *testing.T, client *gophercloud.ServiceClient, portID string, remoteIP string) (*tapmirrors.TapMirror, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTapMirror will delete a Tap Mirror with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteTapMirror(t *testing.T, client *gophercloud.ServiceClient, mirrorID string) {
	_ = "STUB: not implemented"
	return
}
