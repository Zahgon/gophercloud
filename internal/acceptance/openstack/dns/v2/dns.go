package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/dns/v2/recordsets"
	transferAccepts "github.com/gophercloud/gophercloud/v2/openstack/dns/v2/transfer/accept"
	transferRequests "github.com/gophercloud/gophercloud/v2/openstack/dns/v2/transfer/request"
	"github.com/gophercloud/gophercloud/v2/openstack/dns/v2/tsigkeys"
	"github.com/gophercloud/gophercloud/v2/openstack/dns/v2/zones"
)

// CreateRecordSet will create a RecordSet with a random name. An error will
// be returned if the zone was unable to be created.
func CreateRecordSet(t *testing.T, client *gophercloud.ServiceClient, zone *zones.Zone) (*recordsets.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateZone will create a Zone with a random name. An error will
// be returned if the zone was unable to be created.
func CreateZone(t *testing.T, client *gophercloud.ServiceClient) (*zones.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecondaryZone will create a Zone with a random name. An error will
// be returned if the zone was unable to be created.
//
// This is only for example purposes as it will try to do a zone transfer.
func CreateSecondaryZone(t *testing.T, client *gophercloud.ServiceClient) (*zones.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateTransferRequest will create a Transfer Request to a spectified Zone. An error will
// be returned if the zone transfer request was unable to be created.
func CreateTransferRequest(t *testing.T, client *gophercloud.ServiceClient, zone *zones.Zone, targetProjectID string) (*transferRequests.TransferRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateTransferAccept will accept a spectified Transfer Request. An error will
// be returned if the zone transfer accept was unable to be created.
func CreateTransferAccept(t *testing.T, client *gophercloud.ServiceClient, zoneTransferRequestID string, key string) (*transferAccepts.TransferAccept, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTransferRequest will delete a specified zone transfer request. A fatal error will occur if
// the transfer request failed to be deleted. This works best when used as a deferred
// function.
func DeleteTransferRequest(t *testing.T, client *gophercloud.ServiceClient, tr *transferRequests.TransferRequest) {
	_ = "STUB: not implemented"
	return
}

// CreateShare will create a zone share. An error will be returned if the
// zone share was unable to be created.
func CreateShare(t *testing.T, client *gophercloud.ServiceClient, zone *zones.Zone, targetProjectID string) (*zones.ZoneShare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnshareZone will unshare a zone. An error will be returned if the
// zone unshare was unable to be created.
func UnshareZone(t *testing.T, client *gophercloud.ServiceClient, share *zones.ZoneShare) {
	_ = "STUB: not implemented"
	return
}

// DeleteRecordSet will delete a specified record set. A fatal error will occur if
// the record set failed to be deleted. This works best when used as a deferred
// function.
func DeleteRecordSet(t *testing.T, client *gophercloud.ServiceClient, rs *recordsets.RecordSet) {
	_ = "STUB: not implemented"
	return
}

// DeleteZone will delete a specified zone. A fatal error will occur if
// the zone failed to be deleted. This works best when used as a deferred
// function.
func DeleteZone(t *testing.T, client *gophercloud.ServiceClient, zone *zones.Zone) {
	_ = "STUB: not implemented"
	return
}

// WaitForRecordSetStatus will poll a record set's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForRecordSetStatus(client *gophercloud.ServiceClient, rs *recordsets.RecordSet, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForTransferRequestStatus will poll a transfer reqeust's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForTransferRequestStatus(client *gophercloud.ServiceClient, tr *transferRequests.TransferRequest, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForTransferAcceptStatus will poll a transfer accept's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForTransferAcceptStatus(client *gophercloud.ServiceClient, ta *transferAccepts.TransferAccept, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForZoneStatus will poll a zone's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForZoneStatus(client *gophercloud.ServiceClient, zone *zones.Zone, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTSIGKey will create a TSIG key with a random name. An error will
// be returned if the TSIG key was unable to be created.
func CreateTSIGKey(t *testing.T, client *gophercloud.ServiceClient) (*tsigkeys.TSIGKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default pool ID from designate/conf/central.py

// DeleteTSIGKey will delete a specified TSIG key. A fatal error will occur if
// the TSIG key failed to be deleted. This works best when used as a deferred
// function.
func DeleteTSIGKey(t *testing.T, client *gophercloud.ServiceClient, tsigkey *tsigkeys.TSIGKey) {
	_ = "STUB: not implemented"
	return
}
