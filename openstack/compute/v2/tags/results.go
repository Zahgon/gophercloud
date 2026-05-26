package tags

import (
	"github.com/gophercloud/gophercloud/v2"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a tags resource.
func (r commonResult) Extract() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type ListResult struct {
	commonResult
}

// CheckResult is the result from the Check operation.
type CheckResult struct {
	gophercloud.Result
}

func (r CheckResult) Extract() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ReplaceAllResult is the result from the ReplaceAll operation.
type ReplaceAllResult struct {
	commonResult
}

// AddResult is the result from the Add operation.
type AddResult struct {
	gophercloud.ErrResult
}

// DeleteResult is the result from the Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}
