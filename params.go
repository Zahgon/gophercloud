package gophercloud

import (
	"net/url"
	"reflect"
	"time"
)

/*
BuildRequestBody builds a map[string]interface from the given `struct`, or
collection of `structs`. If parent is not an empty string, the final
map[string]interface returned will encapsulate the built one. Parent is
required when passing a list of `structs`.
For example:

	disk := 1
	createOpts := flavors.CreateOpts{
	  ID:         "1",
	  Name:       "m1.tiny",
	  Disk:       &disk,
	  RAM:        512,
	  VCPUs:      1,
	  RxTxFactor: 1.0,
	}

	body, err := gophercloud.BuildRequestBody(createOpts, "flavor")


	opts := []rules.CreateOpts{
		{
			Direction:     "ingress",
			PortRangeMin:  80,
			EtherType:     rules.EtherType4,
			PortRangeMax:  80,
			Protocol:      "tcp",
			SecGroupID:    "a7734e61-b545-452d-a3cd-0189cbd9747a",
		},
		{
			Direction:    "ingress",
			PortRangeMin: 443,
			EtherType:    rules.EtherType4,
			PortRangeMax: 443,
			Protocol:     "tcp",
			SecGroupID:   "a7734e61-b545-452d-a3cd-0189cbd9747a",
		},
	}

	body, err := gophercloud.BuildRequestBody(opts, "security_group_rules")

The above examples can be run as-is, however it is recommended to look at how
BuildRequestBody is used within Gophercloud to more fully understand how it
fits within the request process as a whole rather than use it directly as shown
above.
*/
func BuildRequestBody(opts any, parent string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//fmt.Printf("optsValue.Kind() is a reflect.Struct: %+v\n", optsValue.Kind())

//fmt.Printf("Skipping field: %s...\n", f.Name)

//fmt.Printf("Starting on field: %s...\n", f.Name)

//fmt.Printf("v is zero?: %v\n", zero)

// if the field has a required tag that's set to "true"

//fmt.Printf("Checking required field [%s]:\n\tv: %+v\n\tisZero:%v\n", f.Name, v.Interface(), zero)
// if the field's value is zero, return a missing-argument error

// if the field has a 'required' tag, it can't have a zero-value

//fmt.Printf("Checking `xor` tag for field [%s] with value %+v:\n\txorTag: %s\n", f.Name, v, xorTag)

//fmt.Printf("Checking `or` tag for field with:\n\tname: %+v\n\torTag:%s\n", f.Name, orTag)
//fmt.Printf("field is zero?: %v\n", zero)

//fmt.Printf("value before change: %+v\n", optsValue.Field(i))

//fmt.Printf("value after change: %+v\n", optsValue.Field(i))

//fmt.Printf("Calling BuildRequestBody with:\n\tv: %+v\n\tf.Name:%s\n", v.Interface(), f.Name)

//fmt.Printf("opts: %+v \n", opts)

//fmt.Printf("string(b): %s\n", string(b))

//fmt.Printf("optsMap: %+v\n", optsMap)

//fmt.Printf("optsMap after parent added: %+v\n", optsMap)

// Return an error if we can't work with the underlying type of 'opts'

// EnabledState is a convenience type, mostly used in Create and Update
// operations. Because the zero value of a bool is FALSE, we need to use a
// pointer instead to indicate zero-ness.
type EnabledState *bool

// Convenience vars for EnabledState values.
var (
	iTrue  = true
	iFalse = false

	Enabled  EnabledState = &iTrue
	Disabled EnabledState = &iFalse
)

// IPVersion is a type for the possible IP address versions. Valid instances
// are IPv4 and IPv6
type IPVersion int

const (
	// IPv4 is used for IP version 4 addresses
	IPv4 IPVersion = 4
	// IPv6 is used for IP version 6 addresses
	IPv6 IPVersion = 6
)

// IntToPointer is a function for converting integers into integer pointers.
// This is useful when passing in options to operations.
func IntToPointer(i int) *int {
	_ = "STUB: not implemented"

	/*
	   MaybeString is an internal function to be used by request methods in individual
	   resource packages.

	   It takes a string that might be a zero value and returns either a pointer to its
	   address or nil. This is useful for allowing users to conveniently omit values
	   from an options struct by leaving them zeroed, but still pass nil to the JSON
	   serializer so they'll be omitted from the request body.
	*/return nil
}

func MaybeString(original string) *string { _ = "STUB: not implemented"; return nil }

/*
MaybeInt is an internal function to be used by request methods in individual
resource packages.

Like MaybeString, it accepts an int that may or may not be a zero value, and
returns either a pointer to its address or nil. It's intended to hint that the
JSON serializer should omit its field.
*/
func MaybeInt(original int) *int { _ = "STUB: not implemented"; return nil }

/*
func isUnderlyingStructZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Pointer:
		return isUnderlyingStructZero(v.Elem())
	default:
		return isZero(v)
	}
}
*/

var t time.Time

func isZero(v reflect.Value) bool {
	_ = "STUB: not implemented"
	// fmt.Printf("\n\nchecking isZero for value: %+v\n", v)
	return false
}

// Compare other types directly:

//fmt.Printf("zero type for value: %+v\n\n\n", z)

/*
BuildQueryString is an internal function to be used by request methods in
individual resource packages.

It accepts a tagged structure and expands it into a URL struct. Field names are
converted into query parameters based on a "q" tag. For example:

	type struct Something {
	   Bar string `q:"x_bar"`
	   Baz int    `q:"lorem_ipsum"`
	}

	instance := Something{
	   Bar: "AAA",
	   Baz: "BBB",
	}

will be converted into "?x_bar=AAA&lorem_ipsum=BBB".

The struct's fields may be strings, integers, slices, or boolean values. Fields
left at their type's zero value will be omitted from the query.

Slice are handled in one of two ways:

	type struct Something {
	   Bar []string `q:"bar"` // E.g. ?bar=1&bar=2
	   Baz []int    `q:"baz" format="comma-separated"` // E.g. ?baz=1,2
	}
*/
func BuildQueryString(opts any) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

// if the field has a 'q' tag, it goes in the query string

// if the field is set, add it to the slice of query pieces

// if the field has a 'required' tag, it can't have a zero-value

// Return an error if the underlying type of 'opts' isn't a struct.

/*
BuildHeaders is an internal function to be used by request methods in
individual resource packages.

It accepts an arbitrary tagged structure and produces a string map that's
suitable for use as the HTTP headers of an outgoing request. Field names are
mapped to header names based in "h" tags.

	type struct Something {
	  Bar string `h:"x_bar"`
	  Baz int    `h:"lorem_ipsum"`
	}

	instance := Something{
	  Bar: "AAA",
	  Baz: "BBB",
	}

will be converted into:

	map[string]string{
	  "x_bar": "AAA",
	  "lorem_ipsum": "BBB",
	}

Untagged fields and fields left at their zero values are skipped. Integers,
booleans and string values are supported.
*/
func BuildHeaders(opts any) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

// if the field has a 'h' tag, it goes in the header

// if the field is set, add it to the slice of query pieces

// if the field has a 'required' tag, it can't have a zero-value

// Return an error if the underlying type of 'opts' isn't a struct.

// IDSliceToQueryString takes a slice of elements and converts them into a query
// string. For example, if name=foo and slice=[]int{20, 40, 60}, then the
// result would be `?name=20&name=40&name=60'
func IDSliceToQueryString(name string, ids []int) string { _ = "STUB: not implemented"; return "" }

// IntWithinRange returns TRUE if an integer falls within a defined range, and
// FALSE if not.
func IntWithinRange(val, min, max int) bool { _ = "STUB: not implemented"; return false }
