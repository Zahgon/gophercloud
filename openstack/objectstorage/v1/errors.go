package v1

import (
	"github.com/gophercloud/gophercloud/v2"
)

func CheckContainerName(s string) error { _ = "STUB: not implemented"; return nil }

func CheckObjectName(s string) error { _ = "STUB: not implemented"; return nil }

// ErrInvalidContainerName signals a container name containing an illegal
// character.
type ErrInvalidContainerName struct {
	name string
	gophercloud.BaseError
}

func (e ErrInvalidContainerName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrEmptyContainerName signals an empty container name.
type ErrEmptyContainerName struct {
	gophercloud.BaseError
}

func (e ErrEmptyContainerName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrEmptyObjectName signals an empty container name.
type ErrEmptyObjectName struct {
	gophercloud.BaseError
}

func (e ErrEmptyObjectName) Error() string { _ = "STUB: not implemented"; return "" }
