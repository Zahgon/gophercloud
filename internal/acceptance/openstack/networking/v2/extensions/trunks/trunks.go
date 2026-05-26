package trunks

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/trunks"
)

func CreateTrunk(t *testing.T, client *gophercloud.ServiceClient, parentPortID string, subportIDs ...string) (trunk *trunks.Trunk, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteTrunk(t *testing.T, client *gophercloud.ServiceClient, trunkID string) {
	_ = "STUB: not implemented"
	return
}
