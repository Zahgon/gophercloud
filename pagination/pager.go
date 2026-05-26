package pagination

import (
	"context"
	"errors"

	"github.com/gophercloud/gophercloud/v2"
)

var (
	// ErrPageNotAvailable is returned from a Pager when a next or previous page is requested, but does not exist.
	ErrPageNotAvailable = errors.New("the requested page does not exist")
)

// Page must be satisfied by the result type of any resource collection.
// It allows clients to interact with the resource uniformly, regardless of whether or not or how it's paginated.
// Generally, rather than implementing this interface directly, implementors should embed one of the concrete PageBase structs,
// instead.
// Depending on the pagination strategy of a particular resource, there may be an additional subinterface that the result type
// will need to implement.
type Page interface {
	// NextPageURL generates the URL for the page of data that follows this collection.
	// Return "" if no such page exists.
	NextPageURL(endpointURL string) (string, error)

	// IsEmpty returns true if this Page has no items in it.
	IsEmpty() (bool, error)

	// GetBody returns the Page Body. This is used in the `AllPages` method.
	GetBody() any
}

// Pager knows how to advance through a specific resource collection, one page at a time.
type Pager struct {
	client *gophercloud.ServiceClient

	initialURL string

	createPage func(r PageResult) Page

	firstPage Page

	Err error

	// Headers supplies additional HTTP headers to populate on each paged request.
	Headers map[string]string
}

// NewPager constructs a manually-configured pager.
// Supply the URL for the first page, a function that requests a specific page given a URL, and a function that counts a page.
func NewPager(client *gophercloud.ServiceClient, initialURL string, createPage func(r PageResult) Page) Pager {
	_ = "STUB: not implemented"
	return *new(Pager)
}

// WithPageCreator returns a new Pager that substitutes a different page creation function. This is
// useful for overriding List functions in delegation.
func (p Pager) WithPageCreator(createPage func(r PageResult) Page) Pager {
	_ = "STUB: not implemented"
	return *new(Pager)
}

func (p Pager) fetchNextPage(ctx context.Context, url string) (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

// EachPage iterates over each page returned by a Pager, yielding one at a time
// to a handler function. Return "false" from the handler to prematurely stop
// iterating.
func (p Pager) EachPage(ctx context.Context, handler func(context.Context, Page) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// if first page has already been fetched, no need to fetch it again

// AllPages returns all the pages from a `List` operation in a single page,
// allowing the user to retrieve all the pages at once.
func (p Pager) AllPages(ctx context.Context) (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

// pagesSlice holds all the pages until they get converted into as Page Body.

// body will contain the final concatenated Page body.

// Grab a first page to ascertain the page body type.

// Store the page type so we can use reflection to create a new mega-page of
// that type.

// if it's a single page, just return the firstPage (first page)

// store the first page to avoid getting it twice

// Switch on the page body type. Recognized types are `map[string]any`,
// `[]byte`, and `[]any`.

// key is the map key for the page body if the body type is `map[string]any`.

// Iterate over the pages to concatenate the bodies.

// If it's a linked page, we don't want the `links`, we want the other one.

// check the field's type. we only want []any (which is really []map[string]any)

// Set body to value of type `map[string]any`

// Iterate over the pages to concatenate the bodies.

// seperate pages with a comma

// Remove the trailing comma.

// Combine the slice of slices in to a single slice.

// Set body to value of type `bytes`.

// Iterate over the pages to concatenate the bodies.

// Set body to value of type `[]any`

// Each `Extract*` function is expecting a specific type of page coming back,
// otherwise the type assertion in those functions will fail. pageType is needed
// to create a type in this method that has the same type that the `Extract*`
// function is expecting and set the Body of that object to the concatenated
// pages.

// Set the page body to be the concatenated pages.

// Set any additional headers that were pass along. The `objectstorage` pacakge,
// for example, passes a Content-Type header.

// Type assert the page to a Page interface so that the type assertion in the
// `Extract*` methods will work.
