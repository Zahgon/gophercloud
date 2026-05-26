package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/orchestration/v1/stacks"
)

const basicTemplateResourceName = "secgroup_1"
const basicTemplate = `
	{
		"heat_template_version": "2013-05-23",
		"description": "Simple template to test heat commands",
		"resources": {
			"secgroup_1": {
				"type": "OS::Neutron::SecurityGroup",
				"properties": {
					"description": "Gophercloud test",
					"name": "secgroup_1"
				}
			}
		}
	}
`

const validateTemplate = `
	{
		"heat_template_version": "2013-05-23",
		"description": "Simple template to test heat commands",
		"parameters": {
			"flavor": {
				"default": "m1.tiny",
				"type":    "string"
			}
		},
		"resources": {
			"hello_world": {
				"type": "OS::Nova::Server",
				"properties": {
					"key_name": "heat_key",
					"flavor": {
						"get_param": "flavor"
					},
					"image":     "ad091b52-742f-469e-8f3c-fd81cadf0743",
					"user_data": "#!/bin/bash -xv\necho \"hello world\" &gt; /root/hello-world.txt\n"
				}
			}
		}
	}
`

// CreateStack will create a heat stack with a randomly generated name.
// An error will be returned if the stack failed to be created.
func CreateStack(t *testing.T, client *gophercloud.ServiceClient) (*stacks.RetrievedStack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteStack deletes a stack via its ID.
// A fatal error will occur if the stack failed to be deleted. This works
// best when used as a deferred function.
func DeleteStack(t *testing.T, client *gophercloud.ServiceClient, stackName, stackID string) {
	_ = "STUB: not implemented"
	return
}

// WaitForStackStatus will wait until a stack has reached a certain status.
func WaitForStackStatus(client *gophercloud.ServiceClient, stackName, stackID, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateStackWithFile will create a heat stack with a randomly generated name that uses get_file.
// An error will be returned if the stack failed to be created.
func CreateStackWithFile(t *testing.T, client *gophercloud.ServiceClient) (*stacks.RetrievedStack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
