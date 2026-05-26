package pagination

// LinkedPageBase may be embedded to implement a page that provides navigational "Next" and "Previous" links within its result.
type LinkedPageBase struct {
	PageResult

	// LinkPath lists the keys that should be traversed within a response to arrive at the "next" pointer.
	// If any link along the path is missing, an empty URL will be returned.
	// If any link results in an unexpected value type, an error will be returned.
	// When left as "nil", []string{"links", "next"} will be used as a default.
	LinkPath []string
}

// NextPageURL extracts the pagination structure from a JSON response and returns the "next" link, if one is present.
// It assumes that the links are available in a "links" element of the top-level response object.
// If this is not the case, override NextPageURL on your result type.
func (current LinkedPageBase) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Actual null element.

// IsEmpty satisifies the IsEmpty method of the Page interface
func (current LinkedPageBase) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// GetBody returns the linked page's body. This method is needed to satisfy the
// Page interface.
func (current LinkedPageBase) GetBody() any { _ = "STUB: not implemented"; return *new(any) }
