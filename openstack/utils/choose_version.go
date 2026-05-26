package utils

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Version is a supported API version, corresponding to a vN package within the appropriate service.
type Version struct {
	ID       string
	Suffix   string
	Priority int
}

var goodStatus = map[string]bool{
	"current":   true,
	"supported": true,
	"stable":    true,
}

// ChooseVersion queries the base endpoint of an API to choose the identity service version.
// It will pick a version among the recognized, taking into account the priority and avoiding
// experimental alternatives from the published versions. However, if the client specifies a full
// endpoint that is among the recognized versions, it will be used regardless of priority.
// It returns the highest-Priority Version, OR exact match with client endpoint,
// among the alternatives that are provided, as well as its corresponding endpoint.
func ChooseVersion(ctx context.Context, client *gophercloud.ProviderClient, recognized []*Version) (*Version, string, error) {
	_ = "STUB: not implemented"
	// TODO(stephenfin): This could be removed since we can accomplish this with GetServiceVersions now.
	return nil, "", nil
}

// If a full endpoint is specified, check version suffixes for a match first.

// Prefer a version that exactly matches the provided endpoint.

// Otherwise, find the highest-priority version with a whitelisted status.
