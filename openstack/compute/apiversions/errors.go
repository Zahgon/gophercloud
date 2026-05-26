package apiversions

// ErrVersionNotFound is the error when the requested API version
// could not be found.
type ErrVersionNotFound struct{}

func (e ErrVersionNotFound) Error() string { _ = "STUB: not implemented"; return "" }
