package projects

// InvalidListFilter is returned by the ToUserListQuery method when validation of
// a filter does not pass
type InvalidListFilter struct {
	FilterName string
}

func (e InvalidListFilter) Error() string { _ = "STUB: not implemented"; return "" }
