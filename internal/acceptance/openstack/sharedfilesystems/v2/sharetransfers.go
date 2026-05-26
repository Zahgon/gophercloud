package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/shares"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/sharetransfers"
)

func CreateTransferRequest(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share, name string) (*sharetransfers.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AcceptTransfer(t *testing.T, client *gophercloud.ServiceClient, transferRequest *sharetransfers.Transfer) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteTransferRequest(t *testing.T, client *gophercloud.ServiceClient, transfer *sharetransfers.Transfer) {
	_ = "STUB: not implemented"
	return
}
