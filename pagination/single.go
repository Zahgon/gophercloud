package pagination

// SinglePageBase may be embedded in a Page that contains all of the results from an operation at once.
type SinglePageBase PageResult

// NextPageURL always returns "" to indicate that there are no more pages to return.
func (current SinglePageBase) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"

	// IsEmpty satisifies the IsEmpty method of the Page interface
	return "", nil
}

func (current SinglePageBase) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GetBody returns the single page's body. This method is needed to satisfy the
// Page interface.
func (current SinglePageBase) GetBody() any { _ = "STUB: not implemented"; return *new(any) }
