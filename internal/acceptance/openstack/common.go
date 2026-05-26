// Package openstack contains common functions that can be used
// across all OpenStack components for acceptance testing.
package openstack

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/common/extensions"
)

// PrintExtension prints an extension and all of its attributes.
func PrintExtension(t *testing.T, extension *extensions.Extension) {
	_ = "STUB: not implemented"
	return
}
