package members

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

/*
Create member for specific image

# Preconditions

  - The specified images must exist.
  - You can only add a new member to an image which 'visibility' attribute is
    private.
  - You must be the owner of the specified image.

# Synchronous Postconditions

With correct permissions, you can see the member status of the image as
pending through API calls.

More details here:
http://developer.openstack.org/api-ref-image-v2.html#createImageMember-v2
*/
func Create(ctx context.Context, client *gophercloud.ServiceClient, id string, member string) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// List members returns list of members for specifed image id.
func List(client *gophercloud.ServiceClient, id string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get image member details.
func Get(ctx context.Context, client *gophercloud.ServiceClient, imageID string, memberID string) (r DetailsResult) {
	_ = "STUB: not implemented"
	return *new(DetailsResult)
}

// Delete membership for given image. Callee should be image owner.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, imageID string, memberID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToImageMemberUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options to an Update request.
type UpdateOpts struct {
	Status string
}

// ToMemberUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToImageMemberUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update function updates member.
func Update(ctx context.Context, client *gophercloud.ServiceClient, imageID string, memberID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
