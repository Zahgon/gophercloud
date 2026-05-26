package tools

import (
	"context"
	"testing"
	"time"
)

// WaitFor uses WaitForTimeout to poll a predicate function once per second to
// wait for a certain state to arrive, with a default timeout of 600 seconds.
func WaitFor(predicate func(context.Context) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForTimeout polls a predicate function once per second to wait for a
// certain state to arrive, or until the given timeout is reached.
func WaitForTimeout(predicate func(context.Context) (bool, error), timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// MakeNewPassword generates a new string that's guaranteed to be different than the given one.
func MakeNewPassword(oldPass string) string { _ = "STUB: not implemented"; return "" }

// RandomString generates a string of given length, but random content.
// All content will be within the ASCII graphic character set.
func RandomString(prefix string, n int) string { _ = "STUB: not implemented"; return "" }

// RandomFunnyString returns a random string of the given length filled with
// funny Unicode code points.
func RandomFunnyString(length int) string { _ = "STUB: not implemented"; return "" }

// RandomFunnyStringNoSlash returns a random string of the given length filled with
// funny Unicode code points, but no forward slash.
func RandomFunnyStringNoSlash(length int) string { _ = "STUB: not implemented"; return "" }

func randomString(charset []rune, length int) string { _ = "STUB: not implemented"; return "" }

// RandomInt will return a random integer between a specified range.
func RandomInt(min, max int) int { _ = "STUB: not implemented"; return 0 }

func RandomUUID() string { _ = "STUB: not implemented"; return "" }

// Elide returns the first bit of its input string with a suffix of "..." if it's longer than
// a comfortable 40 characters.
func Elide(value string) string { _ = "STUB: not implemented"; return "" }

// PrintResource returns a resource as a readable structure
func PrintResource(t *testing.T, resource any) { _ = "STUB: not implemented"; return }
