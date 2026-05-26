package segments

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Segment model
type Segment struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	NetworkID       string    `json:"network_id"`
	NetworkType     string    `json:"network_type"`
	PhysicalNetwork string    `json:"physical_network"`
	SegmentationID  int       `json:"segmentation_id"`
	RevisionNumber  int       `json:"revision_number"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// SegmentPage wraps a page of segments.
type SegmentPage struct {
	pagination.LinkedPageBase
}

func (r SegmentPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ExtractSegments(r pagination.Page) ([]Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractSegmentsInto extracts the elements into a slice of Segment structs.
func ExtractSegmentsInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// Segment results
type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Segment, error) { _ = "STUB: not implemented"; return nil, nil }

func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

type GetResult struct {
	commonResult
}

type CreateResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}
