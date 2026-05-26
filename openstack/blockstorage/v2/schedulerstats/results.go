package schedulerstats

import (
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Capabilities represents the information of an individual StoragePool.
type Capabilities struct {
	// The following fields should be present in all storage drivers.
	DriverVersion     string  `json:"driver_version"`
	FreeCapacityGB    float64 `json:"-"`
	StorageProtocol   string  `json:"storage_protocol"`
	TotalCapacityGB   float64 `json:"-"`
	VendorName        string  `json:"vendor_name"`
	VolumeBackendName string  `json:"volume_backend_name"`

	// The following fields are optional and may have empty values depending
	// on the storage driver in use.
	ReservedPercentage       int64   `json:"reserved_percentage"`
	LocationInfo             string  `json:"location_info"`
	QoSSupport               bool    `json:"QoS_support"`
	ProvisionedCapacityGB    float64 `json:"provisioned_capacity_gb"`
	MaxOverSubscriptionRatio string  `json:"-"`
	ThinProvisioningSupport  bool    `json:"thin_provisioning_support"`
	ThickProvisioningSupport bool    `json:"thick_provisioning_support"`
	TotalVolumes             int64   `json:"total_volumes"`
	FilterFunction           string  `json:"filter_function"`
	GoodnessFunction         string  `json:"goodness_function"`
	Multiattach              bool    `json:"multiattach"`
	SparseCopyVolume         bool    `json:"sparse_copy_volume"`
	AllocatedCapacityGB      float64 `json:"-"`
}

// StoragePool represents an individual StoragePool retrieved from the
// schedulerstats API.
type StoragePool struct {
	Name         string       `json:"name"`
	Capabilities Capabilities `json:"capabilities"`
}

func (r *Capabilities) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Generic function to parse a capacity value which may be a numeric
// value, "unknown", or "infinite"

// StoragePoolPage is a single page of all List results.
type StoragePoolPage struct {
	pagination.SinglePageBase
}

// IsEmpty satisfies the IsEmpty method of the Page interface. It returns true
// if a List contains no results.
func (page StoragePoolPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractStoragePools takes a List result and extracts the collection of
// StoragePools returned by the API.
func ExtractStoragePools(p pagination.Page) ([]StoragePool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
