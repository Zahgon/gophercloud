package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/flavorprofiles"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

const FlavorProfilesListBody = `
{
	"flavorprofiles": [
        {
            "id": "c55d080d-af45-47ee-b48c-4caa5e87724f",
            "name": "amphora-single",
            "provider_name": "amphora",
            "flavor_data": "{\"loadbalancer_topology\": \"SINGLE\"}"
        },
		{
            "id": "f78d2815-3714-4b6e-91d8-cf821ba01017",
            "name": "amphora-act-stdby",
            "provider_name": "amphora",
            "flavor_data": "{\"loadbalancer_topology\": \"ACTIVE_STANDBY\"}"
        }
    ]
}
`

const SingleFlavorProfileBody = `
{
	"flavorprofile": {
		"id": "dcd65be5-f117-4260-ab3d-b32cc5bd1272",
		"name": "amphora-test",
		"provider_name": "amphora",
		"flavor_data": "{\"loadbalancer_topology\": \"ACTIVE_STANDBY\"}"
	}
}
`

const PostUpdateFlavorBody = `
{
	"flavorprofile": {
		"id": "dcd65be5-f117-4260-ab3d-b32cc5bd1272",
		"name": "amphora-test-updated",
		"provider_name": "amphora",
		"flavor_data": "{\"loadbalancer_topology\": \"SINGLE\"}"
	}
}
`

var (
	FlavorProfileSingle = flavorprofiles.FlavorProfile{
		ID:           "c55d080d-af45-47ee-b48c-4caa5e87724f",
		Name:         "amphora-single",
		ProviderName: "amphora",
		FlavorData:   "{\"loadbalancer_topology\": \"SINGLE\"}",
	}

	FlavorProfileAct = flavorprofiles.FlavorProfile{
		ID:           "f78d2815-3714-4b6e-91d8-cf821ba01017",
		Name:         "amphora-act-stdby",
		ProviderName: "amphora",
		FlavorData:   "{\"loadbalancer_topology\": \"ACTIVE_STANDBY\"}",
	}

	FlavorDb = flavorprofiles.FlavorProfile{
		ID:           "dcd65be5-f117-4260-ab3d-b32cc5bd1272",
		Name:         "amphora-test",
		ProviderName: "amphora",
		FlavorData:   "{\"loadbalancer_topology\": \"ACTIVE_STANDBY\"}",
	}

	FlavorUpdated = flavorprofiles.FlavorProfile{
		ID:           "dcd65be5-f117-4260-ab3d-b32cc5bd1272",
		Name:         "amphora-test-updated",
		ProviderName: "amphora",
		FlavorData:   "{\"loadbalancer_topology\": \"SINGLE\"}",
	}
)

func HandleFlavorProfileListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorProfileCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorProfileGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorProfileDeletionSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleFlavorProfileUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}
