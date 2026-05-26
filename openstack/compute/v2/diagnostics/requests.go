package diagnostics

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Diagnostics
func Get(ctx context.Context, client *gophercloud.ServiceClient, serverId string) (r serverDiagnosticsResult) {
	_ = "STUB: not implemented"
	return *new(serverDiagnosticsResult)
}
