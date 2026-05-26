package client

import (
	"github.com/gophercloud/gophercloud/v2"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"

// ServiceClient returns a generic service client for use in tests.
func ServiceClient(fakeServer th.FakeServer) *gophercloud.ServiceClient {
	_ = "STUB: not implemented"
	return nil
}
