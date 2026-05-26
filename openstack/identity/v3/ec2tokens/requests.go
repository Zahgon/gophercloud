package ec2tokens

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

const (
	// EC2CredentialsAwsRequestV4 is a constant, used to generate AWS
	// Credential V4.
	EC2CredentialsAwsRequestV4 = "aws4_request"
	// EC2CredentialsHmacSha1V2 is a HMAC SHA1 signature method. Used to
	// generate AWS Credential V2.
	EC2CredentialsHmacSha1V2 = "HmacSHA1"
	// EC2CredentialsHmacSha256V2 is a HMAC SHA256 signature method. Used
	// to generate AWS Credential V2.
	EC2CredentialsHmacSha256V2 = "HmacSHA256"
	// EC2CredentialsAwsHmacV4 is an AWS signature V4 signing method.
	// More details:
	// https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
	EC2CredentialsAwsHmacV4 = "AWS4-HMAC-SHA256"
	// EC2CredentialsTimestampFormatV4 is an AWS signature V4 timestamp
	// format.
	EC2CredentialsTimestampFormatV4 = "20060102T150405Z"
	// EC2CredentialsDateFormatV4 is an AWS signature V4 date format.
	EC2CredentialsDateFormatV4 = "20060102"
)

// AuthOptions represents options for authenticating a user using EC2 credentials.
type AuthOptions struct {
	// Access is the EC2 Credential Access ID.
	Access string `json:"access" required:"true"`
	// Secret is the EC2 Credential Secret, used to calculate signature.
	// Not used, when a Signature is is.
	Secret string `json:"-"`
	// Host is a HTTP request Host header. Used to calculate an AWS
	// signature V2. For signature V4 set the Host inside Headers map.
	// Optional.
	Host string `json:"host"`
	// Path is a HTTP request path. Optional.
	Path string `json:"path"`
	// Verb is a HTTP request method. Optional.
	Verb string `json:"verb"`
	// Headers is a map of HTTP request headers. Optional.
	Headers map[string]string `json:"headers"`
	// Region is a region name to calculate an AWS signature V4. Optional.
	Region string `json:"-"`
	// Service is a service name to calculate an AWS signature V4. Optional.
	Service string `json:"-"`
	// Params is a map of GET method parameters. Optional.
	Params map[string]string `json:"params"`
	// AllowReauth allows Gophercloud to re-authenticate automatically
	// if/when your token expires.
	AllowReauth bool `json:"-"`
	// Signature can be either a []byte (encoded to base64 automatically) or
	// a string. You can set the singature explicitly, when you already know
	// it. In this case default Params won't be automatically set. Optional.
	Signature any `json:"signature"`
	// BodyHash is a HTTP request body sha256 hash. When nil and Signature
	// is not set, a random hash is generated. Optional.
	BodyHash *string `json:"body_hash"`
	// Timestamp is a timestamp to calculate a V4 signature. Optional.
	Timestamp *time.Time `json:"-"`
	// Token is a []byte string (encoded to base64 automatically) which was
	// signed by an EC2 secret key. Used by S3 tokens for validation only.
	// Token must be set with a Signature. If a Signature is not provided,
	// a Token will be generated automatically along with a Signature.
	Token []byte `json:"token,omitempty"`
}

// EC2CredentialsBuildCanonicalQueryStringV2 builds a canonical query string
// for an AWS signature V2.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L133
func EC2CredentialsBuildCanonicalQueryStringV2(params map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// EC2CredentialsBuildStringToSignV2 builds a string to sign an AWS signature
// V2.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L148
func EC2CredentialsBuildStringToSignV2(opts AuthOptions) []byte {
	_ = "STUB: not implemented"
	return nil
}

// EC2CredentialsBuildCanonicalQueryStringV2 builds a canonical query string
// for an AWS signature V4.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L244
func EC2CredentialsBuildCanonicalQueryStringV4(verb string, params map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// EC2CredentialsBuildCanonicalHeadersV4 builds a canonical string based on
// "headers" map and "signedHeaders" string parameters.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L216
func EC2CredentialsBuildCanonicalHeadersV4(headers map[string]string, signedHeaders string) string {
	_ = "STUB: not implemented"
	return ""
}

// EC2CredentialsBuildSignatureKeyV4 builds a HMAC 256 signature key based on
// input parameters.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L169
func EC2CredentialsBuildSignatureKeyV4(secret, region, service string, date time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// EC2CredentialsBuildStringToSignV4 builds an AWS v4 signature string to sign
// based on input parameters.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L251
func EC2CredentialsBuildStringToSignV4(opts AuthOptions, signedHeaders string, bodyHash string, date time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// EC2CredentialsBuildSignatureV4 builds an AWS v4 signature based on input
// parameters.
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L285..L286
func EC2CredentialsBuildSignatureV4(key []byte, stringToSign []byte) string {
	_ = "STUB: not implemented"
	return ""
}

// EC2CredentialsBuildAuthorizationHeaderV4 builds an AWS v4 Authorization
// header based on auth parameters, date and signature
func EC2CredentialsBuildAuthorizationHeaderV4(opts AuthOptions, signedHeaders string, signature string, date time.Time) string {
	_ = "STUB: not implemented"
	return ""
}

// ToTokenV3ScopeMap is a dummy method to satisfy tokens.AuthOptionsBuilder
// interface.
func (opts *AuthOptions) ToTokenV3ScopeMap() (map[string]any, error) {
	_ = "STUB: not implemented"

	// ToTokenV3HeadersMap allows AuthOptions to satisfy the AuthOptionsBuilder
	// interface in the v3 tokens package.
	return nil, nil
}

func (opts *AuthOptions) ToTokenV3HeadersMap(map[string]any) (map[string]string, error) {
	_ = "STUB: not implemented"

	// CanReauth is a method method to satisfy tokens.AuthOptionsBuilder interface
	return nil, nil
}

func (opts *AuthOptions) CanReauth() bool { _ = "STUB: not implemented"; return false }

// ToTokenV3CreateMap formats an AuthOptions into a create request.
func (opts *AuthOptions) ToTokenV3CreateMap(map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calculate signature, when it is not set

// detect and process a signature v2

// params is a map of strings

// keystone uses this method only when HmacSHA256 is not available on the server side
// https://github.com/openstack/python-keystoneclient/blob/stable/train/keystoneclient/contrib/ec2/utils.py#L151..L156

// it is not a signature v2, but a signature v4

// when body_hash is not set, generate a random one

// token is only used for S3 tokens validation and will be removed when using EC2 validation

// Create authenticates and either generates a new token from EC2 credentials
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts tokens.AuthOptionsBuilder) (r tokens.CreateResult) {
	_ = "STUB: not implemented"
	return *new(tokens.CreateResult)
}

// delete "token" element, since it is used in s3tokens

// ValidateS3Token authenticates an S3 request using EC2 credentials. Doesn't
// generate a new token ID, but returns a tokens.CreateResult.
func ValidateS3Token(ctx context.Context, c *gophercloud.ServiceClient, opts tokens.AuthOptionsBuilder) (r tokens.CreateResult) {
	_ = "STUB: not implemented"
	return *new(tokens.CreateResult)
}

// delete unused element, since it is used in ec2tokens only

// The following are small helper functions used to help build the signature.

// sumHMAC1 is a func to implement the HMAC SHA1 signature method.
func sumHMAC1(key []byte, data []byte) []byte { _ = "STUB: not implemented"; return nil }

// sumHMAC256 is a func to implement the HMAC SHA256 signature method.
func sumHMAC256(key []byte, data []byte) []byte { _ = "STUB: not implemented"; return nil }

// randomBodyHash is a func to generate a random sha256 hexdigest.
func randomBodyHash() (string, error) { _ = "STUB: not implemented"; return "", nil }

// interfaceToMap is a func used to represent a "credentials" map element as a
// "map[string]string"
func interfaceToMap(c map[string]any, key string) map[string]string {
	_ = "STUB: not implemented"
	// convert map[string]any to map[string]string
	return nil
}

// deleteBodyElements deletes map body elements
func deleteBodyElements(b map[string]any, elements ...string) { _ = "STUB: not implemented"; return }
