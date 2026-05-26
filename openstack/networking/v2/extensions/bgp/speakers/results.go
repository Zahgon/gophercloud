package speakers

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

const jroot = "bgp_speaker"

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a bgp speaker resource.
func (r commonResult) Extract() (*BGPSpeaker, error) { _ = "STUB: not implemented"; return nil, nil }

func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// BGPSpeaker BGP Speaker
type BGPSpeaker struct {
	// UUID for the bgp speaker
	ID string `json:"id"`

	// Human-readable name for the bgp speaker. Might not be unique.
	Name string `json:"name"`

	// TenantID is the project owner of the bgp speaker.
	TenantID string `json:"tenant_id"`

	// ProjectID is the project owner of the bgp speaker.
	ProjectID string `json:"project_id"`

	// If the speaker would advertise floating ip host routes
	AdvertiseFloatingIPHostRoutes bool `json:"advertise_floating_ip_host_routes"`

	// If the speaker would advertise tenant networks
	AdvertiseTenantNetworks bool `json:"advertise_tenant_networks"`

	// IP version
	IPVersion int `json:"ip_version"`

	// Local Autonomous System
	LocalAS int `json:"local_as"`

	// The uuid of the Networks configured with this speaker
	Networks []string `json:"networks"`

	// The uuid of the BGP Peer Configured with this speaker
	Peers []string `json:"peers"`
}

// BGPSpeakerPage is the page returned by a pager when traversing over a
// collection of bgp speakers.
type BGPSpeakerPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether a BGPSpeakerPage struct is empty.
func (r BGPSpeakerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractBGPSpeakers accepts a Page struct, specifically a BGPSpeakerPage struct,
// and extracts the elements into a slice of BGPSpeaker structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBGPSpeakers(r pagination.Page) ([]BGPSpeaker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractBGPSpeakersInto accepts a Page struct and an any. The former contains
// a list of BGPSpeaker and the later should be used to store the result that would be
// extracted from the former.
func ExtractBGPSpeakersInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a BGPSpeaker.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a BGPSpeaker.
type CreateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a BGPSpeaker.
type UpdateResult struct {
	commonResult
}

// AddBGPPeerResult represent the response of the PUT /v2.0/bgp-speakers/{bgp-speaker-id}/add-bgp-peer
type AddBGPPeerResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a AddBGPPeerResult resource
func (r AddBGPPeerResult) Extract() (*AddBGPPeerOpts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AddBGPPeerResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// RemoveBGPPeerResult represent the response of the PUT /v2.0/bgp-speakers/{bgp-speaker-id}/remove-bgp-peer
// There is no body content for the response of a successful DELETE request.
type RemoveBGPPeerResult struct {
	gophercloud.ErrResult
}

// AdvertisedRoute represents an advertised route
type AdvertisedRoute struct {
	// NextHop IP address
	NextHop string `json:"next_hop"`

	// Destination Network
	Destination string `json:"destination"`
}

// AdvertisedRoutePage is the page returned by a pager when you call
type AdvertisedRoutePage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether a AdvertisedRoutePage struct is empty.
func (r AdvertisedRoutePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractAdvertisedRoutes accepts a Page struct, a.k.a. AdvertisedRoutePage struct,
// and extracts the elements into a slice of AdvertisedRoute structs.
func ExtractAdvertisedRoutes(r pagination.Page) ([]AdvertisedRoute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractAdvertisedRoutesInto extract the advertised routes from the first param into the 2nd
func ExtractAdvertisedRoutesInto(r pagination.Page, v any) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGatewayNetworkResult represents the data that would be PUT to
// /v2.0/bgp-speakers/{bgp-speaker-id}/add_gateway_network
type AddGatewayNetworkResult struct {
	gophercloud.Result
}

func (r AddGatewayNetworkResult) Extract() (*AddGatewayNetworkOpts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveGatewayNetworkResult represents the data that would be PUT to
// /v2.0/bgp-speakers/{bgp-speaker-id}/remove_gateway_network
type RemoveGatewayNetworkResult struct {
	gophercloud.ErrResult
}
