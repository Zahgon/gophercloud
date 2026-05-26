package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/messages"
)

// DeleteMessage will delete a message. An error will occur if
// the message was unable to be deleted.
func DeleteMessage(t *testing.T, client *gophercloud.ServiceClient, message *messages.Message) {
	_ = "STUB: not implemented"
	return
}
