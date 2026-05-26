package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/keymanager/v1/containers"
	"github.com/gophercloud/gophercloud/v2/openstack/keymanager/v1/orders"
	"github.com/gophercloud/gophercloud/v2/openstack/keymanager/v1/secrets"
)

// CreateAsymmetric Order will create a random asymmetric order.
// An error will be returned if the order could not be created.
func CreateAsymmetricOrder(t *testing.T, client *gophercloud.ServiceClient) (*orders.Order, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateCertificateContainer will create a random certificate container.
// An error will be returned if the container could not be created.
func CreateCertificateContainer(t *testing.T, client *gophercloud.ServiceClient, passphrase, private, certificate *secrets.Secret) (*containers.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateKeyOrder will create a random key order.
// An error will be returned if the order could not be created.
func CreateKeyOrder(t *testing.T, client *gophercloud.ServiceClient) (*orders.Order, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRSAContainer will create a random RSA container.
// An error will be returned if the container could not be created.
func CreateRSAContainer(t *testing.T, client *gophercloud.ServiceClient, passphrase, private, public *secrets.Secret) (*containers.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateCertificateSecret will create a random certificate secret. An error
// will be returned if the secret could not be created.
func CreateCertificateSecret(t *testing.T, client *gophercloud.ServiceClient, cert []byte) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateEmptySecret will create a random secret with no payload. An error will
// be returned if the secret could not be created.
func CreateEmptySecret(t *testing.T, client *gophercloud.ServiceClient) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateGenericContainer will create a random generic container with a
// specified secret. An error will be returned if the container could not
// be created.
func CreateGenericContainer(t *testing.T, client *gophercloud.ServiceClient, secret *secrets.Secret) (*containers.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReplaceGenericContainerSecretRef will replace the container old secret
// reference with a new one. An error will be returned if the reference could
// not be replaced.
func ReplaceGenericContainerSecretRef(t *testing.T, client *gophercloud.ServiceClient, container *containers.Container, secretOld *secrets.Secret, secretNew *secrets.Secret) error {
	_ = "STUB: not implemented"
	return nil
}

// CreatePassphraseSecret will create a random passphrase secret.
// An error will be returned if the secret could not be created.
func CreatePassphraseSecret(t *testing.T, client *gophercloud.ServiceClient, passphrase string) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePublicSecret will create a random public secret. An error
// will be returned if the secret could not be created.
func CreatePublicSecret(t *testing.T, client *gophercloud.ServiceClient, pub []byte) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePrivateSecret will create a random private secret. An error
// will be returned if the secret could not be created.
func CreatePrivateSecret(t *testing.T, client *gophercloud.ServiceClient, priv []byte) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecretWithPayload will create a random secret with a given payload.
// An error will be returned if the secret could not be created.
func CreateSecretWithPayload(t *testing.T, client *gophercloud.ServiceClient, payload string) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSymmetricSecret will create a random symmetric secret. An error
// will be returned if the secret could not be created.
func CreateSymmetricSecret(t *testing.T, client *gophercloud.ServiceClient) (*secrets.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteContainer will delete a container. A fatal error will occur if the
// container could not be deleted. This works best when used as a deferred
// function.
func DeleteContainer(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

// DeleteOrder will delete an order. A fatal error will occur if the
// order could not be deleted. This works best when used as a deferred
// function.
func DeleteOrder(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

// DeleteSecret will delete a secret. A fatal error will occur if the secret
// could not be deleted. This works best when used as a deferred function.
func DeleteSecret(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

func ParseID(ref string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CreateCertificate will create a random certificate. A fatal error will
// be returned if creation failed.
// https://golang.org/src/crypto/tls/generate_cert.go
func CreateCertificate(t *testing.T, passphrase string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateRSAKeyPair will create a random RSA key pair. An error will be
// returned if the pair could not be created.
func CreateRSAKeyPair(t *testing.T, passphrase string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func WaitForOrder(client *gophercloud.ServiceClient, orderID string) error {
	_ = "STUB: not implemented"
	return nil
}
