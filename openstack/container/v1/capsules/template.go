package capsules

// Template is a structure that represents OpenStack Zun Capsule templates
type Template struct {
	// Bin stores the contents of the template or environment.
	Bin []byte
	// Parsed contains a parsed version of Bin. Since there are 2 different
	// fields referring to the same value, you must be careful when accessing
	// this filed.
	Parsed map[string]any
}

// Parse will parse the contents and then validate. The contents MUST be either JSON or YAML.
func (t *Template) Parse() error { _ = "STUB: not implemented"; return nil }
