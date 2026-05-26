package capsules

import (
	"github.com/gophercloud/gophercloud/v2"
)

type ErrInvalidDataFormat struct {
	gophercloud.BaseError
}

func (e ErrInvalidDataFormat) Error() string { _ = "STUB: not implemented"; return "" }
