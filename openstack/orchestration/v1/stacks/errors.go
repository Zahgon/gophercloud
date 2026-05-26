package stacks

import (
	"github.com/gophercloud/gophercloud/v2"
)

type ErrInvalidEnvironment struct {
	gophercloud.BaseError
	Section string
}

func (e ErrInvalidEnvironment) Error() string { _ = "STUB: not implemented"; return "" }

type ErrInvalidDataFormat struct {
	gophercloud.BaseError
}

func (e ErrInvalidDataFormat) Error() string { _ = "STUB: not implemented"; return "" }

type ErrInvalidTemplateFormatVersion struct {
	gophercloud.BaseError
	Version string
}

func (e ErrInvalidTemplateFormatVersion) Error() string { _ = "STUB: not implemented"; return "" }

type ErrTemplateRequired struct {
	gophercloud.BaseError
}

func (e ErrTemplateRequired) Error() string { _ = "STUB: not implemented"; return "" }
