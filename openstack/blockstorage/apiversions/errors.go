package apiversions

// ErrVersionNotFound is the error when the requested API version
// could not be found.
type ErrVersionNotFound struct{}

func (e ErrVersionNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMultipleVersionsFound is the error when a request for an API
// version returns multiple results.
type ErrMultipleVersionsFound struct {
	Count int
}

func (e ErrMultipleVersionsFound) Error() string { _ = "STUB: not implemented"; return "" }
