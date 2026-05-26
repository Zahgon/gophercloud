package utils

func parseEndpoint(endpoint string, includeVersion bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BaseEndpoint will return a URL without the /vX.Y
// portion of the URL.
func BaseEndpoint(endpoint string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// BaseVersionedEndpoint will return a URL with the /vX.Y portion of the URL,
// if present, but without a project ID or similar
func BaseVersionedEndpoint(endpoint string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
