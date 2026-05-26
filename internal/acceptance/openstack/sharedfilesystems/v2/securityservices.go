package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/securityservices"
)

// CreateSecurityService will create a security service with a random name. An
// error will be returned if the security service was unable to be created.
func CreateSecurityService(t *testing.T, client *gophercloud.ServiceClient) (*securityservices.SecurityService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSecurityService will delete a security service. An error will occur if
// the security service was unable to be deleted.
func DeleteSecurityService(t *testing.T, client *gophercloud.ServiceClient, securityService *securityservices.SecurityService) {
	_ = "STUB: not implemented"
	return
}
