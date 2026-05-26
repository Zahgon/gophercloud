package diagnostics

import (
	"github.com/gophercloud/gophercloud/v2"
)

type serverDiagnosticsResult struct {
	gophercloud.Result
}

// Extract interprets any diagnostic response as a map
func (r serverDiagnosticsResult) Extract() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
