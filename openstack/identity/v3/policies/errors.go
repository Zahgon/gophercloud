package policies

// InvalidListFilter is returned by the ToPolicyListQuery method when
// validation of a filter does not pass
type InvalidListFilter struct {
	FilterName string
}

func (e InvalidListFilter) Error() string { _ = "STUB: not implemented"; return "" }

// StringFieldLengthExceedsLimit is returned by the
// ToPolicyCreateMap/ToPolicyUpdateMap methods when validation of
// a type does not pass
type StringFieldLengthExceedsLimit struct {
	Field string
	Limit int
}

func (e StringFieldLengthExceedsLimit) Error() string { _ = "STUB: not implemented"; return "" }
