package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/availabilityzoneprofiles"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

const AvailabilityZoneProfilesListBody = `
{
	"availability_zone_profiles": [
        {
            "id": "1d334061-d807-4997-8f34-9fe428ba37df",
            "name": "availability-zone-profile-first",
            "provider_name": "amphora",
            "availability_zone_data": "{\"compute_zone\": \"nova\"}"
        },
		{
            "id": "56f45d00-86e4-4bea-8525-19e835776c4e",
            "name": "availability-zone-profile-second",
            "provider_name": "amphora",
            "availability_zone_data": "{\"compute_zone\": \"nova\"}"
        }
    ]
}
`

const SingleAvailabilityZoneProfileBody = `
{
	"availability_zone_profile": {
		"id": "13be083b-f502-426e-8500-07600f98b91b",
		"name": "availability-zone-profile",
		"provider_name": "amphora",
		"availability_zone_data": "{\"compute_zone\": \"nova\"}"
	}
}
`

const PostUpdateAvailabilityZoneFlavorBody = `
{
	"availability_zone_profile": {
		"id": "13be083b-f502-426e-8500-07600f98b91b",
		"name": "availability-zone-profile-updated",
		"provider_name": "amphora",
		"availability_zone_data": "{\"compute_zone\": \"nova\"}"
	}
}
`

var (
	AvailabilityZoneProfileSingle = availabilityzoneprofiles.AvailabilityZoneProfile{
		ID:                   "1d334061-d807-4997-8f34-9fe428ba37df",
		Name:                 "availability-zone-profile-first",
		ProviderName:         "amphora",
		AvailabilityZoneData: "{\"compute_zone\": \"nova\"}",
	}

	AvailabilityZoneProfileAct = availabilityzoneprofiles.AvailabilityZoneProfile{
		ID:                   "56f45d00-86e4-4bea-8525-19e835776c4e",
		Name:                 "availability-zone-profile-second",
		ProviderName:         "amphora",
		AvailabilityZoneData: "{\"compute_zone\": \"nova\"}",
	}

	AvailabilityZoneProfileDb = availabilityzoneprofiles.AvailabilityZoneProfile{
		ID:                   "13be083b-f502-426e-8500-07600f98b91b",
		Name:                 "availability-zone-profile",
		ProviderName:         "amphora",
		AvailabilityZoneData: "{\"compute_zone\": \"nova\"}",
	}

	AvailabilityZoneProfileUpdated = availabilityzoneprofiles.AvailabilityZoneProfile{
		ID:                   "13be083b-f502-426e-8500-07600f98b91b",
		Name:                 "availability-zone-profile-updated",
		ProviderName:         "amphora",
		AvailabilityZoneData: "{\"compute_zone\": \"nova\"}",
	}
)

func HandleAvailabilityZoneProfileListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleAvailabilityZoneProfileCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	_ = "STUB: not implemented"
	return
}

func HandleAvailabilityZoneProfileGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleAvailabilityZoneProfileDeletionSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

func HandleAvailabilityZoneProfileUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}
