package ruletypes

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*RuleType, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a RuleType.
type GetResult struct {
	commonResult
}

// RuleType represents a single QoS rule type.
type RuleType struct {
	Type    string   `json:"type"`
	Drivers []Driver `json:"drivers"`
}

// Driver represents a single QoS driver.
type Driver struct {
	Name                string               `json:"name"`
	SupportedParameters []SupportedParameter `json:"supported_parameters"`
}

// SupportedParameter represents a single set of supported parameters for a some QoS driver's .
type SupportedParameter struct {
	ParameterName   string `json:"parameter_name"`
	ParameterType   string `json:"parameter_type"`
	ParameterValues any    `json:"parameter_values"`
}

type ListRuleTypesPage struct {
	pagination.SinglePageBase
}

func (r ListRuleTypesPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ExtractRuleTypes(r pagination.Page) ([]RuleType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
