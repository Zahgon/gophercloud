package openstack

import (
	"github.com/gophercloud/gophercloud/v2"
)

var nilOptions = gophercloud.AuthOptions{}

/*
AuthOptionsFromEnv fills out an identity.AuthOptions structure with the
settings found on the various OpenStack OS_* environment variables.

The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
OS_PASSWORD and OS_PROJECT_ID.

Of these, OS_USERNAME, OS_PASSWORD, and OS_AUTH_URL must have settings,
or an error will result.  OS_PROJECT_ID, is optional.

OS_TENANT_ID and OS_TENANT_NAME are deprecated forms of OS_PROJECT_ID and
OS_PROJECT_NAME and the latter are expected against a v3 auth api.

If OS_PROJECT_ID and OS_PROJECT_NAME are set, they will still be referred
as "tenant" in Gophercloud.

If OS_PROJECT_NAME is set, it requires OS_DOMAIN_ID or OS_DOMAIN_NAME to be
set as well to handle projects not on the default domain.

To use this function, first set the OS_* environment variables (for example,
by sourcing an `openrc` file), then:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(context.TODO(), opts)
*/
func AuthOptionsFromEnv() (gophercloud.AuthOptions, error) {
	_ = "STUB: not implemented"
	return *new(gophercloud.AuthOptions), nil
}

// If OS_PROJECT_ID is set, overwrite tenantID with the value.

// If OS_PROJECT_NAME is set, overwrite tenantName with the value.

// Empty username and userID could be ignored, when applicationCredentialID and applicationCredentialSecret are set

// silently ignore TOTP passcode warning, since it is not a common auth method
