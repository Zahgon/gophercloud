package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
)

// RequireNeutronExtension will restrict a test to be only run in environments
// with the requested Neutron extension present.
func RequireNeutronExtension(t *testing.T, client *gophercloud.ServiceClient, extension string) {
	_ = "STUB: not implemented"
	return
}
