package gophercloud

import (
	"net/http"
)

// BaseError is an error type that all other error types embed.
type BaseError struct {
	DefaultErrString string
	Info             string
}

func (e BaseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e BaseError) choseErrString() string { _ = "STUB: not implemented"; return "" }

// ErrMissingInput is the error when input is required in a particular
// situation but not provided by the user
type ErrMissingInput struct {
	BaseError
	Argument string
}

func (e ErrMissingInput) Error() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidInput is an error type used for most non-HTTP Gophercloud errors.
type ErrInvalidInput struct {
	ErrMissingInput
	Value any
}

func (e ErrInvalidInput) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMissingEnvironmentVariable is the error when environment variable is required
// in a particular situation but not provided by the user
type ErrMissingEnvironmentVariable struct {
	BaseError
	EnvironmentVariable string
}

func (e ErrMissingEnvironmentVariable) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMissingAnyoneOfEnvironmentVariables is the error when anyone of the environment variables
// is required in a particular situation but not provided by the user
type ErrMissingAnyoneOfEnvironmentVariables struct {
	BaseError
	EnvironmentVariables []string
}

func (e ErrMissingAnyoneOfEnvironmentVariables) Error() string {
	_ = "STUB: not implemented"
	return ""
}

// ErrUnexpectedResponseCode is returned by the Request method when a response code other than
// those listed in OkCodes is encountered.
type ErrUnexpectedResponseCode struct {
	BaseError
	URL            string
	Method         string
	Expected       []int
	Actual         int
	Body           []byte
	ResponseHeader http.Header
}

func (e ErrUnexpectedResponseCode) Error() string { _ = "STUB: not implemented"; return "" }

// GetStatusCode returns the actual status code of the error.
func (e ErrUnexpectedResponseCode) GetStatusCode() int {
	_ = "STUB: not implemented"

	// ResponseCodeIs returns true if this error is or contains an ErrUnexpectedResponseCode reporting
	// that the request failed with the given response code. For example, this checks if a request
	// failed because of a 404 error:
	//
	//	allServers, err := servers.List(client, servers.ListOpts{})
	//	if gophercloud.ResponseCodeIs(err, http.StatusNotFound) {
	//		handleNotFound()
	//	}
	//
	// It is safe to pass a nil error, in which case this function always returns false.
	return 0
}

func ResponseCodeIs(err error, status int) bool { _ = "STUB: not implemented"; return false }

// ErrTimeOut is the error type returned when an operations times out.
type ErrTimeOut struct {
	BaseError
}

func (e ErrTimeOut) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUnableToReauthenticate is the error type returned when reauthentication fails.
type ErrUnableToReauthenticate struct {
	BaseError
	ErrOriginal error
	ErrReauth   error
}

func (e ErrUnableToReauthenticate) Error() string { _ = "STUB: not implemented"; return "" }

// ErrErrorAfterReauthentication is the error type returned when reauthentication
// succeeds, but an error occurs afterword (usually an HTTP error).
type ErrErrorAfterReauthentication struct {
	BaseError
	ErrOriginal error
}

func (e ErrErrorAfterReauthentication) Error() string { _ = "STUB: not implemented"; return "" }

// ErrServiceNotFound is returned when no service in a service catalog matches
// the provided EndpointOpts. This is generally returned by provider service
// factory methods like "NewComputeV2()" and can mean that a service is not
// enabled for your account.
type ErrServiceNotFound struct {
	BaseError
}

func (e ErrServiceNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrEndpointNotFound is returned when no available endpoints match the
// provided EndpointOpts. This is also generally returned by provider service
// factory methods, and usually indicates that a region was specified
// incorrectly.
type ErrEndpointNotFound struct {
	BaseError
}

func (e ErrEndpointNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrResourceNotFound is the error when trying to retrieve a resource's
// ID by name and the resource doesn't exist.
type ErrResourceNotFound struct {
	BaseError
	Name         string
	ResourceType string
}

func (e ErrResourceNotFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMultipleResourcesFound is the error when trying to retrieve a resource's
// ID by name and multiple resources have the user-provided name.
type ErrMultipleResourcesFound struct {
	BaseError
	Name         string
	Count        int
	ResourceType string
}

func (e ErrMultipleResourcesFound) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUnexpectedType is the error when an unexpected type is encountered
type ErrUnexpectedType struct {
	BaseError
	Expected string
	Actual   string
}

func (e ErrUnexpectedType) Error() string { _ = "STUB: not implemented"; return "" }

func unacceptedAttributeErr(attribute string) string { _ = "STUB: not implemented"; return "" }

func redundantWithTokenErr(attribute string) string { _ = "STUB: not implemented"; return "" }

func redundantWithUserID(attribute string) string { _ = "STUB: not implemented"; return "" }

// ErrAPIKeyProvided indicates that an APIKey was provided but can't be used.
type ErrAPIKeyProvided struct{ BaseError }

func (e ErrAPIKeyProvided) Error() string { _ = "STUB: not implemented"; return "" }

// ErrTenantIDProvided indicates that a TenantID was provided but can't be used.
type ErrTenantIDProvided struct{ BaseError }

func (e ErrTenantIDProvided) Error() string { _ = "STUB: not implemented"; return "" }

// ErrTenantNameProvided indicates that a TenantName was provided but can't be used.
type ErrTenantNameProvided struct{ BaseError }

func (e ErrTenantNameProvided) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUsernameWithToken indicates that a Username was provided, but token authentication is being used instead.
type ErrUsernameWithToken struct{ BaseError }

func (e ErrUsernameWithToken) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUserIDWithToken indicates that a UserID was provided, but token authentication is being used instead.
type ErrUserIDWithToken struct{ BaseError }

func (e ErrUserIDWithToken) Error() string { _ = "STUB: not implemented"; return "" }

// ErrDomainIDWithToken indicates that a DomainID was provided, but token authentication is being used instead.
type ErrDomainIDWithToken struct{ BaseError }

func (e ErrDomainIDWithToken) Error() string { _ = "STUB: not implemented"; return "" }

// ErrDomainNameWithToken indicates that a DomainName was provided, but token authentication is being used instead.s
type ErrDomainNameWithToken struct{ BaseError }

func (e ErrDomainNameWithToken) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUsernameOrUserID indicates that neither username nor userID are specified, or both are at once.
type ErrUsernameOrUserID struct{ BaseError }

func (e ErrUsernameOrUserID) Error() string { _ = "STUB: not implemented"; return "" }

// ErrDomainIDWithUserID indicates that a DomainID was provided, but unnecessary because a UserID is being used.
type ErrDomainIDWithUserID struct{ BaseError }

func (e ErrDomainIDWithUserID) Error() string { _ = "STUB: not implemented"; return "" }

// ErrDomainNameWithUserID indicates that a DomainName was provided, but unnecessary because a UserID is being used.
type ErrDomainNameWithUserID struct{ BaseError }

func (e ErrDomainNameWithUserID) Error() string { _ = "STUB: not implemented"; return "" }

// ErrDomainIDOrDomainName indicates that a username was provided, but no domain to scope it.
// It may also indicate that both a DomainID and a DomainName were provided at once.
type ErrDomainIDOrDomainName struct{ BaseError }

func (e ErrDomainIDOrDomainName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMissingPassword indicates that no password was provided and no token is available.
type ErrMissingPassword struct{ BaseError }

func (e ErrMissingPassword) Error() string { _ = "STUB: not implemented"; return "" }

// ErrScopeDomainIDOrDomainName indicates that a domain ID or Name was required in a Scope, but not present.
type ErrScopeDomainIDOrDomainName struct{ BaseError }

func (e ErrScopeDomainIDOrDomainName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrScopeProjectIDOrProjectName indicates that both a ProjectID and a ProjectName were provided in a Scope.
type ErrScopeProjectIDOrProjectName struct{ BaseError }

func (e ErrScopeProjectIDOrProjectName) Error() string { _ = "STUB: not implemented"; return "" }

// ErrScopeProjectIDAlone indicates that a ProjectID was provided with other constraints in a Scope.
type ErrScopeProjectIDAlone struct{ BaseError }

func (e ErrScopeProjectIDAlone) Error() string { _ = "STUB: not implemented"; return "" }

// ErrScopeEmpty indicates that no credentials were provided in a Scope.
type ErrScopeEmpty struct{ BaseError }

func (e ErrScopeEmpty) Error() string { _ = "STUB: not implemented"; return "" }

// ErrAppCredMissingSecret indicates that no Application Credential Secret was provided with Application Credential ID or Name
type ErrAppCredMissingSecret struct{ BaseError }

func (e ErrAppCredMissingSecret) Error() string { _ = "STUB: not implemented"; return "" }
