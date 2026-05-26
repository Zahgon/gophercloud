package manageablevolumes

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
)

type ManageExistingResult struct {
	gophercloud.Result
}

// Extract will get the Volume object out of the ManageExistingResult object.
func (r ManageExistingResult) Extract() (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractInto converts our response data into a volume struct
func (r ManageExistingResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }
