package messages

import (
	"github.com/gophercloud/gophercloud/v2"
)

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func createURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func getURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func DeleteMessageURL(client *gophercloud.ServiceClient, queueName string, messageID string) string {
	_ = "STUB: not implemented"
	return ""
}

func messageURL(client *gophercloud.ServiceClient, queueName string, messageID string) string {
	_ = "STUB: not implemented"
	return ""
}

// builds next page full url based on service endpoint
func nextPageURL(endpointURL, next string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
