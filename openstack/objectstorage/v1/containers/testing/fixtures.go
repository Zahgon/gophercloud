package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/objectstorage/v1/containers"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

type handlerOptions struct {
	path string
}

type option func(*handlerOptions)

func WithPath(s string) option { _ = "STUB: not implemented"; return *new(option) }

// ExpectedListInfo is the result expected from a call to `List` when full
// info is requested.
var ExpectedListInfo = []containers.Container{
	{
		Count: 0,
		Bytes: 0,
		Name:  "janeausten",
	},
	{
		Count: 1,
		Bytes: 14,
		Name:  "marktwain",
	},
}

// ExpectedListNames is the result expected from a call to `List` when just
// container names are requested.
var ExpectedListNames = []string{"janeausten", "marktwain"}

// HandleListContainerInfoSuccessfully creates an HTTP handler at `/` on the test handler mux that
// responds with a `List` response when full info is requested.
func HandleListContainerInfoSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

// HandleListZeroContainerNames204 creates an HTTP handler at `/` on the test handler mux that
// responds with "204 No Content" when container names are requested. This happens on some, but not all,
// objectstorage instances. This case is peculiar in that the server sends no `content-type` header.
func HandleListZeroContainerNames204(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

// HandleCreateContainerSuccessfully creates an HTTP handler at `/testContainer` on the test handler mux that
// responds with a `Create` response.
func HandleCreateContainerSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

// HandleDeleteContainerSuccessfully creates an HTTP handler at `/testContainer` on the test handler mux that
// responds with a `Delete` response.
func HandleDeleteContainerSuccessfully(t *testing.T, fakeServer th.FakeServer, options ...option) {
	_ = "STUB: not implemented"
	return
}

const bulkDeleteResponse = `
{
    "Response Status": "foo",
    "Response Body": "bar",
    "Errors": [],
    "Number Deleted": 2,
    "Number Not Found": 0
}
`

// HandleBulkDeleteSuccessfully creates an HTTP handler at `/` on the test
// handler mux that responds with a `Delete` response.
func HandleBulkDeleteSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	_ = "STUB: not implemented"
	return
}

// HandleUpdateContainerSuccessfully creates an HTTP handler at `/testContainer` on the test handler mux that
// responds with a `Update` response.
func HandleUpdateContainerSuccessfully(t *testing.T, fakeServer th.FakeServer, options ...option) {
	_ = "STUB: not implemented"
	return
}

// HandleUpdateContainerVersioningOn creates an HTTP handler at `/testVersioning` on the test handler mux that
// responds with a `Update` response.
func HandleUpdateContainerVersioningOn(t *testing.T, fakeServer th.FakeServer, options ...option) {
	_ = "STUB: not implemented"
	return
}

// HandleUpdateContainerVersioningOff creates an HTTP handler at `/testVersioning` on the test handler mux that
// responds with a `Update` response.
func HandleUpdateContainerVersioningOff(t *testing.T, fakeServer th.FakeServer, options ...option) {
	_ = "STUB: not implemented"
	return
}

// HandleGetContainerSuccessfully creates an HTTP handler at `/testContainer` on the test handler mux that
// responds with a `Get` response.
func HandleGetContainerSuccessfully(t *testing.T, fakeServer th.FakeServer, options ...option) {
	_ = "STUB: not implemented"
	return
}
