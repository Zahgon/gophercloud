package pagination

// MarkerPage is a stricter Page interface that describes additional functionality required for use with NewMarkerPager.
// For convenience, embed the MarkedPageBase struct.
type MarkerPage interface {
	Page

	// LastMarker returns the last "marker" value on this page.
	LastMarker() (string, error)
}

// MarkerPageBase is a page in a collection that's paginated by "limit" and "marker" query parameters.
type MarkerPageBase struct {
	PageResult

	// Owner is a reference to the embedding struct.
	Owner MarkerPage
}

// NextPageURL generates the URL for the page of results after this one.
func (current MarkerPageBase) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty satisifies the IsEmpty method of the Page interface
func (current MarkerPageBase) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GetBody returns the linked page's body. This method is needed to satisfy the
// Page interface.
func (current MarkerPageBase) GetBody() any { _ = "STUB: not implemented"; return *new(any) }
