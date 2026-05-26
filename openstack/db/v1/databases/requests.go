package databases

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder builds create options
type CreateOptsBuilder interface {
	ToDBCreateMap() (map[string]any, error)
}

// CreateOpts is the struct responsible for configuring a database; often in
// the context of an instance.
type CreateOpts struct {
	// Specifies the name of the database. Valid names can be composed
	// of the following characters: letters (either case); numbers; these
	// characters '@', '?', '#', ' ' but NEVER beginning a name string; '_' is
	// permitted anywhere. Prohibited characters that are forbidden include:
	// single quotes, double quotes, back quotes, semicolons, commas, backslashes,
	// and forward slashes.
	Name string `json:"name" required:"true"`
	// Set of symbols and encodings. The default character set is
	// "utf8". See http://dev.mysql.com/doc/refman/5.1/en/charset-mysql.html for
	// supported character sets.
	CharSet string `json:"character_set,omitempty"`
	// Set of rules for comparing characters in a character set. The
	// default value for collate is "utf8_general_ci". See
	// http://dev.mysql.com/doc/refman/5.1/en/charset-mysql.html for supported
	// collations.
	Collate string `json:"collate,omitempty"`
}

// ToMap is a helper function to convert individual DB create opt structures
// into sub-maps.
func (opts CreateOpts) ToMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// BatchCreateOpts allows for multiple databases to created and modified.
type BatchCreateOpts []CreateOpts

// ToDBCreateMap renders a JSON map for creating DBs.
func (opts BatchCreateOpts) ToDBCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new database within the specified instance. If the
// specified instance does not exist, a 404 error will be returned.
func Create(ctx context.Context, client *gophercloud.ServiceClient, instanceID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// List will list all of the databases for a specified instance. Note: this
// operation will only return user-defined databases; it will exclude system
// databases like "mysql", "information_schema", "lost+found" etc.
func List(client *gophercloud.ServiceClient, instanceID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Delete will permanently delete the database within a specified instance.
// All contained data inside the database will also be permanently deleted.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, instanceID, dbName string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
