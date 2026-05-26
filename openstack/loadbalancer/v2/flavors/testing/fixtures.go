package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/flavors"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

const FlavorsListBody = `
{
	"flavors": [
        {
            "id": "4c82a610-8c7f-4a72-8cca-42f584e3f6d1",
            "name": "Basic",
            "description": "A basic standalone Octavia load balancer.",
            "enabled": true,
            "flavor_profile_id": "bdba88c7-beab-4fc9-a5dd-3635de59185b"
        },
		{
            "id": "0af3b9cc-9284-44c2-9494-0ec337fa31bb",
            "name": "Advance",
            "description": "A advance standalone Octavia load balancer.",
            "enabled": false,
            "flavor_profile_id": "c221abc6-a845-45a0-925c-27110c9d7bdc"
        }
    ]
}
`

const SingleFlavorBody = `
{
	"flavor": {
		"id": "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		"name": "Basic",
		"description": "A basic standalone Octavia load balancer.",
		"enabled": true,
		"flavor_profile_id": "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1"
	}
}
`

const SingleFlavorDisabledBody = `
{
	"flavor": {
		"id": "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		"name": "Basic",
		"description": "A basic standalone Octavia load balancer.",
		"enabled": false,
		"flavor_profile_id": "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1"
	}
}
`

const PostUpdateFlavorBody = `
{
	"flavor": {
		"id": "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		"name": "Basic v2",
		"description": "Rename flavor",
		"enabled": false,
		"flavor_profile_id": "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1"
	}
}
`

var (
	FlavorBasic = flavors.Flavor{
		ID:              "4c82a610-8c7f-4a72-8cca-42f584e3f6d1",
		Name:            "Basic",
		Description:     "A basic standalone Octavia load balancer.",
		Enabled:         true,
		FlavorProfileID: "bdba88c7-beab-4fc9-a5dd-3635de59185b",
	}

	FlavorAdvance = flavors.Flavor{
		ID:              "0af3b9cc-9284-44c2-9494-0ec337fa31bb",
		Name:            "Advance",
		Description:     "A advance standalone Octavia load balancer.",
		Enabled:         false,
		FlavorProfileID: "c221abc6-a845-45a0-925c-27110c9d7bdc",
	}

	FlavorDb = flavors.Flavor{
		ID:              "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		Name:            "Basic",
		Description:     "A basic standalone Octavia load balancer.",
		Enabled:         true,
		FlavorProfileID: "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1",
	}

	FlavorDisabled = flavors.Flavor{
		ID:              "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		Name:            "Basic",
		Description:     "A basic standalone Octavia load balancer.",
		Enabled:         false,
		FlavorProfileID: "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1",
	}

	FlavorUpdated = flavors.Flavor{
		ID:              "5548c807-e6e8-43d7-9ea4-b38d34dd74a0",
		Name:            "Basic v2",
		Description:     "Rename flavor",
		Enabled:         false,
		FlavorProfileID: "9daa2768-74e7-4d13-bf5d-1b8e0dc239e1",
	}
)

func HandleFlavorListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorCreationSuccessfullyDisabled(t *testing.T, fakeServer th.FakeServer, response string) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorDeletionSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}
