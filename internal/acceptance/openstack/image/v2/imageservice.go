// Package v2 contains common functions for creating image resources
// for use in acceptance tests. See the `*_test.go` files for example usages.
package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/imageimport"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/tasks"
)

// CreateEmptyImage will create an image, but with no actual image data.
// An error will be returned if an image was unable to be created.
func CreateEmptyImage(t *testing.T, client *gophercloud.ServiceClient) (*images.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteImage deletes an image.
// A fatal error will occur if the image failed to delete. This works best when
// used as a deferred function.
func DeleteImage(t *testing.T, client *gophercloud.ServiceClient, image *images.Image) {
	_ = "STUB: not implemented"
	return
}

// ImportImageURL contains an URL of a test image that can be imported.
const ImportImageURL = "http://download.cirros-cloud.net/0.4.0/cirros-0.4.0-x86_64-disk.img"

// CreateTask will create a task to import the CirrOS image.
// An error will be returned if a task couldn't be created.
func CreateTask(t *testing.T, client *gophercloud.ServiceClient, imageURL string) (*tasks.Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetImportInfo will retrieve Import API information.
func GetImportInfo(t *testing.T, client *gophercloud.ServiceClient) (*imageimport.ImportInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StageImage will stage local image file to the referenced remote queued image.
func StageImage(t *testing.T, client *gophercloud.ServiceClient, filepath, imageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// DownloadImageFileFromURL will download an image from the specified URL and
// place it into the specified path.
func DownloadImageFileFromURL(t *testing.T, url, filepath string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteImageFile will delete local image file.
func DeleteImageFile(t *testing.T, filepath string) { _ = "STUB: not implemented"; return }

// ImportImage will import image data from the remote source to the Image service.
func ImportImage(t *testing.T, client *gophercloud.ServiceClient, imageID string) error {
	_ = "STUB: not implemented"
	return nil
}
