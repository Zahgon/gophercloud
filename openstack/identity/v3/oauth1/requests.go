package oauth1

import (
	"context"
	"net/url"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Type SignatureMethod is a OAuth1 SignatureMethod type.
type SignatureMethod string

const (
	// HMACSHA1 is a recommended OAuth1 signature method.
	HMACSHA1 SignatureMethod = "HMAC-SHA1"

	// PLAINTEXT signature method is not recommended to be used in
	// production environment.
	PLAINTEXT SignatureMethod = "PLAINTEXT"

	// OAuth1TokenContentType is a supported content type for an OAuth1
	// token.
	OAuth1TokenContentType = "application/x-www-form-urlencoded"
)

// AuthOptions represents options for authenticating a user using OAuth1 tokens.
type AuthOptions struct {
	// OAuthConsumerKey is the OAuth1 Consumer Key.
	OAuthConsumerKey string `q:"oauth_consumer_key" required:"true"`

	// OAuthConsumerSecret is the OAuth1 Consumer Secret. Used to generate
	// an OAuth1 request signature.
	OAuthConsumerSecret string `required:"true"`

	// OAuthToken is the OAuth1 Request Token.
	OAuthToken string `q:"oauth_token" required:"true"`

	// OAuthTokenSecret is the OAuth1 Request Token Secret. Used to generate
	// an OAuth1 request signature.
	OAuthTokenSecret string `required:"true"`

	// OAuthSignatureMethod is the OAuth1 signature method the Consumer used
	// to sign the request. Supported values are "HMAC-SHA1" or "PLAINTEXT".
	// "PLAINTEXT" is not recommended for production usage.
	OAuthSignatureMethod SignatureMethod `q:"oauth_signature_method" required:"true"`

	// OAuthTimestamp is an OAuth1 request timestamp. If nil, current Unix
	// timestamp will be used.
	OAuthTimestamp *time.Time

	// OAuthNonce is an OAuth1 request nonce. Nonce must be a random string,
	// uniquely generated for each request. Will be generated automatically
	// when it is not set.
	OAuthNonce string `q:"oauth_nonce"`

	// AllowReauth allows Gophercloud to re-authenticate automatically
	// if/when your token expires.
	AllowReauth bool
}

// ToTokenV3HeadersMap builds the headers required for an OAuth1-based create
// request.
func (opts AuthOptions) ToTokenV3HeadersMap(headerOpts map[string]any) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToTokenV3ScopeMap allows AuthOptions to satisfy the tokens.AuthOptionsBuilder
// interface.
func (opts AuthOptions) ToTokenV3ScopeMap() (map[string]any, error) {
	_ = "STUB: not implemented"

	// CanReauth allows AuthOptions to satisfy the tokens.AuthOptionsBuilder
	// interface.
	return nil, nil
}

func (opts AuthOptions) CanReauth() bool { _ = "STUB: not implemented"; return false }

// ToTokenV3CreateMap builds a create request body.
func (opts AuthOptions) ToTokenV3CreateMap(map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	// identityReq defines the "identity" portion of an OAuth1-based authentication
	// create request body.
	return nil, nil
}

// authReq defines the "auth" portion of an OAuth1-based authentication
// create request body.

// oauth1Request defines how  an OAuth1-based authentication create
// request body looks.

// Create authenticates and either generates a new OpenStack token
// from an OAuth1 token.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts tokens.AuthOptionsBuilder) (r tokens.CreateResult) {
	_ = "STUB: not implemented"
	return *new(tokens.CreateResult)
}

// CreateConsumerOptsBuilder allows extensions to add additional parameters to
// the CreateConsumer request.
type CreateConsumerOptsBuilder interface {
	ToOAuth1CreateConsumerMap() (map[string]any, error)
}

// CreateConsumerOpts provides options used to create a new Consumer.
type CreateConsumerOpts struct {
	// Description is the consumer description.
	Description string `json:"description"`
}

// ToOAuth1CreateConsumerMap formats a CreateConsumerOpts into a create request.
func (opts CreateConsumerOpts) ToOAuth1CreateConsumerMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateConsumer creates a new Consumer.
func CreateConsumer(ctx context.Context, client *gophercloud.ServiceClient, opts CreateConsumerOptsBuilder) (r CreateConsumerResult) {
	_ = "STUB: not implemented"
	return *new(CreateConsumerResult)
}

// DeleteConsumer deletes a Consumer.
func DeleteConsumer(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteConsumerResult) {
	_ = "STUB: not implemented"
	return *new(DeleteConsumerResult)
}

// List enumerates Consumers.
func ListConsumers(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetConsumer retrieves details on a single Consumer by ID.
func GetConsumer(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetConsumerResult) {
	_ = "STUB: not implemented"
	return *new(GetConsumerResult)
}

// UpdateConsumerOptsBuilder allows extensions to add additional parameters to the
// UpdateConsumer request.
type UpdateConsumerOptsBuilder interface {
	ToOAuth1UpdateConsumerMap() (map[string]any, error)
}

// UpdateConsumerOpts provides options used to update a consumer.
type UpdateConsumerOpts struct {
	// Description is the consumer description.
	Description string `json:"description"`
}

// ToOAuth1UpdateConsumerMap formats an UpdateConsumerOpts into a consumer update
// request.
func (opts UpdateConsumerOpts) ToOAuth1UpdateConsumerMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateConsumer updates an existing Consumer.
func UpdateConsumer(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateConsumerOptsBuilder) (r UpdateConsumerResult) {
	_ = "STUB: not implemented"
	return *new(UpdateConsumerResult)
}

// RequestTokenOptsBuilder allows extensions to add additional parameters to the
// RequestToken request.
type RequestTokenOptsBuilder interface {
	ToOAuth1RequestTokenHeaders(string, string) (map[string]string, error)
}

// RequestTokenOpts provides options used to get a consumer unauthorized
// request token.
type RequestTokenOpts struct {
	// OAuthConsumerKey is the OAuth1 Consumer Key.
	OAuthConsumerKey string `q:"oauth_consumer_key" required:"true"`

	// OAuthConsumerSecret is the OAuth1 Consumer Secret. Used to generate
	// an OAuth1 request signature.
	OAuthConsumerSecret string `required:"true"`

	// OAuthSignatureMethod is the OAuth1 signature method the Consumer used
	// to sign the request. Supported values are "HMAC-SHA1" or "PLAINTEXT".
	// "PLAINTEXT" is not recommended for production usage.
	OAuthSignatureMethod SignatureMethod `q:"oauth_signature_method" required:"true"`

	// OAuthTimestamp is an OAuth1 request timestamp. If nil, current Unix
	// timestamp will be used.
	OAuthTimestamp *time.Time

	// OAuthNonce is an OAuth1 request nonce. Nonce must be a random string,
	// uniquely generated for each request. Will be generated automatically
	// when it is not set.
	OAuthNonce string `q:"oauth_nonce"`

	// RequestedProjectID is a Project ID a consumer user requested an
	// access to.
	RequestedProjectID string `h:"Requested-Project-Id"`
}

// ToOAuth1RequestTokenHeaders formats a RequestTokenOpts into a map of request
// headers.
func (opts RequestTokenOpts) ToOAuth1RequestTokenHeaders(method, u string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RequestToken requests an unauthorized OAuth1 Token.
func RequestToken(ctx context.Context, client *gophercloud.ServiceClient, opts RequestTokenOptsBuilder) (r TokenResult) {
	_ = "STUB: not implemented"
	return *new(TokenResult)
}

// AuthorizeTokenOptsBuilder allows extensions to add additional parameters to
// the AuthorizeToken request.
type AuthorizeTokenOptsBuilder interface {
	ToOAuth1AuthorizeTokenMap() (map[string]any, error)
}

// AuthorizeTokenOpts provides options used to authorize a request token.
type AuthorizeTokenOpts struct {
	Roles []Role `json:"roles"`
}

// Role is a struct representing a role object in a AuthorizeTokenOpts struct.
type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ToOAuth1AuthorizeTokenMap formats an AuthorizeTokenOpts into an authorize token
// request.
func (opts AuthorizeTokenOpts) ToOAuth1AuthorizeTokenMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AuthorizeToken authorizes an unauthorized consumer token.
func AuthorizeToken(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AuthorizeTokenOptsBuilder) (r AuthorizeTokenResult) {
	_ = "STUB: not implemented"
	return *new(AuthorizeTokenResult)
}

// CreateAccessTokenOptsBuilder allows extensions to add additional parameters
// to the CreateAccessToken request.
type CreateAccessTokenOptsBuilder interface {
	ToOAuth1CreateAccessTokenHeaders(string, string) (map[string]string, error)
}

// CreateAccessTokenOpts provides options used to create an OAuth1 token.
type CreateAccessTokenOpts struct {
	// OAuthConsumerKey is the OAuth1 Consumer Key.
	OAuthConsumerKey string `q:"oauth_consumer_key" required:"true"`

	// OAuthConsumerSecret is the OAuth1 Consumer Secret. Used to generate
	// an OAuth1 request signature.
	OAuthConsumerSecret string `required:"true"`

	// OAuthToken is the OAuth1 Request Token.
	OAuthToken string `q:"oauth_token" required:"true"`

	// OAuthTokenSecret is the OAuth1 Request Token Secret. Used to generate
	// an OAuth1 request signature.
	OAuthTokenSecret string `required:"true"`

	// OAuthVerifier is the OAuth1 verification code.
	OAuthVerifier string `q:"oauth_verifier" required:"true"`

	// OAuthSignatureMethod is the OAuth1 signature method the Consumer used
	// to sign the request. Supported values are "HMAC-SHA1" or "PLAINTEXT".
	// "PLAINTEXT" is not recommended for production usage.
	OAuthSignatureMethod SignatureMethod `q:"oauth_signature_method" required:"true"`

	// OAuthTimestamp is an OAuth1 request timestamp. If nil, current Unix
	// timestamp will be used.
	OAuthTimestamp *time.Time

	// OAuthNonce is an OAuth1 request nonce. Nonce must be a random string,
	// uniquely generated for each request. Will be generated automatically
	// when it is not set.
	OAuthNonce string `q:"oauth_nonce"`
}

// ToOAuth1CreateAccessTokenHeaders formats a CreateAccessTokenOpts into a map of
// request headers.
func (opts CreateAccessTokenOpts) ToOAuth1CreateAccessTokenHeaders(method, u string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateAccessToken creates a new OAuth1 Access Token
func CreateAccessToken(ctx context.Context, client *gophercloud.ServiceClient, opts CreateAccessTokenOptsBuilder) (r TokenResult) {
	_ = "STUB: not implemented"
	return *new(TokenResult)
}

// GetAccessToken retrieves details on a single OAuth1 access token by an ID.
func GetAccessToken(ctx context.Context, client *gophercloud.ServiceClient, userID string, id string) (r GetAccessTokenResult) {
	_ = "STUB: not implemented"
	return *new(GetAccessTokenResult)
}

// RevokeAccessToken revokes an OAuth1 access token.
func RevokeAccessToken(ctx context.Context, client *gophercloud.ServiceClient, userID string, id string) (r RevokeAccessTokenResult) {
	_ = "STUB: not implemented"
	return *new(RevokeAccessTokenResult)
}

// ListAccessTokens enumerates authorized access tokens.
func ListAccessTokens(client *gophercloud.ServiceClient, userID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListAccessTokenRoles enumerates authorized access token roles.
func ListAccessTokenRoles(client *gophercloud.ServiceClient, userID string, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetAccessTokenRole retrieves details on a single OAuth1 access token role by
// an ID.
func GetAccessTokenRole(ctx context.Context, client *gophercloud.ServiceClient, userID string, id string, roleID string) (r GetAccessTokenRoleResult) {
	_ = "STUB: not implemented"
	return *new(GetAccessTokenRoleResult)
}

// The following are small helper functions used to help build the signature.

// buildOAuth1QueryString builds a URLEncoded parameters string specific for
// OAuth1-based requests.
func buildOAuth1QueryString(opts any, timestamp *time.Time, callback string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use provided timestamp

// use current timestamp

// when nonce is not set, generate a random one

// buildStringToSign builds a string to be signed.
func buildStringToSign(method string, u string, query url.Values) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Default scheme port must be stripped

// Ensure that URL doesn't contain queries

// signString signs a string using an OAuth1 signature method.
func signString(signatureMethod SignatureMethod, strToSign []byte, signatureKeys []string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildAuthHeader generates an OAuth1 Authorization header with a signature
// calculated using an OAuth1 signature method.
func buildAuthHeader(query url.Values, signature string) string {
	_ = "STUB: not implemented"
	return ""
}
