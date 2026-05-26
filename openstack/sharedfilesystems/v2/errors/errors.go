package errors

type ManilaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type ErrorDetails map[string]ManilaError

// error types from provider_client.go
func ExtractErrorInto(rawError error, errorDetails *ErrorDetails) (err error) {
	_ = "STUB: not implemented"
	return nil
}
