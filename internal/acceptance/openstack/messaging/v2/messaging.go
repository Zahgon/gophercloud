package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/messaging/v2/claims"
	"github.com/gophercloud/gophercloud/v2/openstack/messaging/v2/messages"
	"github.com/gophercloud/gophercloud/v2/openstack/messaging/v2/queues"
)

func CreateQueue(t *testing.T, client *gophercloud.ServiceClient) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func DeleteQueue(t *testing.T, client *gophercloud.ServiceClient, queueName string) {
	_ = "STUB: not implemented"
	return
}

func GetQueue(t *testing.T, client *gophercloud.ServiceClient, queueName string) (queues.QueueDetails, error) {
	_ = "STUB: not implemented"
	return *new(queues.QueueDetails), nil
}

func CreateShare(t *testing.T, client *gophercloud.ServiceClient, queueName string) (queues.QueueShare, error) {
	_ = "STUB: not implemented"
	return *new(queues.QueueShare), nil
}

func CreateMessage(t *testing.T, client *gophercloud.ServiceClient, queueName string) (messages.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(messages.ResourceList), nil
}

func ListMessages(t *testing.T, client *gophercloud.ServiceClient, queueName string) ([]messages.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string) ([]claims.Messages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string, claimID string) (*claims.Claim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string, claimID string) error {
	_ = "STUB: not implemented"
	return nil
}

func ExtractIDs(claim []claims.Messages) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
