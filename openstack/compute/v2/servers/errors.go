package servers

import (
	"github.com/gophercloud/gophercloud/v2"
)

// ErrNeitherImageIDNorImageNameProvided is the error when neither the image
// ID nor the image name is provided for a server operation
type ErrNeitherImageIDNorImageNameProvided struct{ gophercloud.ErrMissingInput }

func (e ErrNeitherImageIDNorImageNameProvided) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNeitherFlavorIDNorFlavorNameProvided is the error when neither the flavor
// ID nor the flavor name is provided for a server operation
type ErrNeitherFlavorIDNorFlavorNameProvided struct{ gophercloud.ErrMissingInput }

func (e ErrNeitherFlavorIDNorFlavorNameProvided) Error() string {
	_ = "STUB: not implemented"
	return ""
}

type ErrNoClientProvidedForIDByName struct{ gophercloud.ErrMissingInput }

func (e ErrNoClientProvidedForIDByName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidHowParameterProvided is the error when an unknown value is given
// for the `how` argument
type ErrInvalidHowParameterProvided struct{ gophercloud.ErrInvalidInput }

// ErrNoAdminPassProvided is the error when an administrative password isn't
// provided for a server operation
type ErrNoAdminPassProvided struct{ gophercloud.ErrMissingInput }

// ErrNoImageIDProvided is the error when an image ID isn't provided for a server
// operation
type ErrNoImageIDProvided struct{ gophercloud.ErrMissingInput }

// ErrNoIDProvided is the error when a server ID isn't provided for a server
// operation
type ErrNoIDProvided struct{ gophercloud.ErrMissingInput }

// ErrServer is a generic error type for servers HTTP operations.
type ErrServer struct {
	gophercloud.ErrUnexpectedResponseCode
	ID string
}

func (se ErrServer) Error() string { _ = "STUB: not implemented"; return "" }

// Error404 overrides the generic 404 error message.
func (se ErrServer) Error404(e gophercloud.ErrUnexpectedResponseCode) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrServerNotFound is the error when a 404 is received during server HTTP
// operations.
type ErrServerNotFound struct {
	ErrServer
}

func (e ErrServerNotFound) Error() string { _ = "STUB: not implemented"; return "" }
