package swauth

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// AuthOptsBuilder describes struct types that can be accepted by the Auth call.
type AuthOptsBuilder interface {
	ToAuthOptsMap() (map[string]string, error)
}

// AuthOpts specifies an authentication request.
type AuthOpts struct {
	// User is an Swauth-based username in username:tenant format.
	User string `h:"X-Auth-User" required:"true"`

	// Key is a secret/password to authenticate the User with.
	Key string `h:"X-Auth-Key" required:"true"`
}

// ToAuthOptsMap formats an AuthOpts structure into a request body.
func (opts AuthOpts) ToAuthOptsMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Auth performs an authentication request for a Swauth-based user.
func Auth(ctx context.Context, c *gophercloud.ProviderClient, opts AuthOptsBuilder) (r GetAuthResult) {
	_ = "STUB: not implemented"
	return *new(GetAuthResult)
}

// NewObjectStorageV1 creates a Swauth-authenticated *gophercloud.ServiceClient
// client that can issue ObjectStorage-based API calls.
func NewObjectStorageV1(ctx context.Context, pc *gophercloud.ProviderClient, authOpts AuthOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
