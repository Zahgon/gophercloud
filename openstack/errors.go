package openstack

import (
	"github.com/gophercloud/gophercloud/v2"
)

// ErrEndpointNotFound is the error when no suitable endpoint can be found
// in the user's catalog
type ErrEndpointNotFound struct{ gophercloud.BaseError }

func (e ErrEndpointNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidAvailabilityProvided is the error when an invalid endpoint
// availability is provided
type ErrInvalidAvailabilityProvided struct{ gophercloud.ErrInvalidInput }

func (e ErrInvalidAvailabilityProvided) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNoAuthURL is the error when the OS_AUTH_URL environment variable is not
// found
type ErrNoAuthURL struct{ gophercloud.ErrInvalidInput }

func (e ErrNoAuthURL) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNoUsername is the error when the OS_USERNAME environment variable is not
// found
type ErrNoUsername struct{ gophercloud.ErrInvalidInput }

func (e ErrNoUsername) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNoPassword is the error when the OS_PASSWORD environment variable is not
// found
type ErrNoPassword struct{ gophercloud.ErrInvalidInput }

func (e ErrNoPassword) Error() string { _ = "STUB: not implemented"; return "" }
