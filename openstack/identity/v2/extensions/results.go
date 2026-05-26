package extensions

import (
	common "github.com/gophercloud/gophercloud/v2/openstack/common/extensions"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ExtensionPage is a single page of Extension results.
type ExtensionPage struct {
	common.ExtensionPage
}

// IsEmpty returns true if the current page contains at least one Extension.
func (page ExtensionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractExtensions accepts a Page struct, specifically an ExtensionPage struct, and extracts the
// elements into a slice of Extension structs.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	_ = "STUB: not implemented"
	// Identity v2 adds an intermediate "values" object.
	return nil, nil
}
