package queues

import (
	"github.com/gophercloud/gophercloud/v2"
)

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func commonURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func createURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func updateURL(client *gophercloud.ServiceClient, queueName string) string {
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

func statURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func shareURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

func purgeURL(client *gophercloud.ServiceClient, queueName string) string {
	_ = "STUB: not implemented"
	return ""
}

// builds next page full url based on service endpoint
func nextPageURL(baseURL string, next string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
