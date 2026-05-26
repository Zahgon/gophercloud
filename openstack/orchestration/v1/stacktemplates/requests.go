package stacktemplates

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retreives data for the given stack template.
func Get(ctx context.Context, c *gophercloud.ServiceClient, stackName, stackID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ValidateOptsBuilder describes struct types that can be accepted by the Validate call.
// The ValidateOpts struct in this package does.
type ValidateOptsBuilder interface {
	ToStackTemplateValidateMap() (map[string]any, error)
}

// ValidateOpts specifies the template validation parameters.
type ValidateOpts struct {
	Template    string `json:"template" or:"TemplateURL"`
	TemplateURL string `json:"template_url" or:"Template"`
}

// ToStackTemplateValidateMap assembles a request body based on the contents of a ValidateOpts.
func (opts ValidateOpts) ToStackTemplateValidateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate validates the given stack template.
func Validate(ctx context.Context, c *gophercloud.ServiceClient, opts ValidateOptsBuilder) (r ValidateResult) {
	_ = "STUB: not implemented"
	return *new(ValidateResult)
}
