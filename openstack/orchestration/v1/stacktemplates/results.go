package stacktemplates

import (
	"github.com/gophercloud/gophercloud/v2"
)

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

// Extract returns the JSON template and is called after a Get operation.
func (r GetResult) Extract() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ValidatedTemplate represents the parsed object returned from a Validate request.
type ValidatedTemplate struct {
	Description     string         `json:"Description"`
	Parameters      map[string]any `json:"Parameters"`
	ParameterGroups map[string]any `json:"ParameterGroups"`
}

// ValidateResult represents the result of a Validate operation.
type ValidateResult struct {
	gophercloud.Result
}

// Extract returns a pointer to a ValidatedTemplate object and is called after a
// Validate operation.
func (r ValidateResult) Extract() (*ValidatedTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
