package segments

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/segments"
)

func CreateSegment(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*segments.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteSegment(t *testing.T, client *gophercloud.ServiceClient, segmentID string) {
	_ = "STUB: not implemented"
	return
}
