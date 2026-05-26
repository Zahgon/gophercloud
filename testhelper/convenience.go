package testhelper

import (
	"reflect"
	"testing"
)

const (
	logBodyFmt = "\033[1;31m%s %s\033[0m"
	greenCode  = "\033[0m\033[1;32m"
	yellowCode = "\033[0m\033[1;33m"
	resetCode  = "\033[0m\033[1;31m"
)

func prefix(depth int) string { _ = "STUB: not implemented"; return "" }

func green(str any) string { _ = "STUB: not implemented"; return "" }

func yellow(str any) string { _ = "STUB: not implemented"; return "" }

func logFatal(t *testing.T, str string) { _ = "STUB: not implemented"; return }

func logError(t *testing.T, str string) { _ = "STUB: not implemented"; return }

type diffLogger func([]string, any, any)

type visit struct {
	a1  uintptr
	a2  uintptr
	typ reflect.Type
}

// Recursively visits the structures of "expected" and "actual". The diffLogger function will be
// invoked with each different value encountered, including the reference path that was followed
// to get there.
func deepDiffEqual(expected, actual reflect.Value, visited map[visit]bool, path []string, logDifference diffLogger) {
	_ = "STUB: not implemented"

	// Fall back to the regular reflect.DeepEquals function.
	return
}

// References are identical. We can short-circuit

// Already visited.

// Remember this visit for later.

func deepDiff(expected, actual any, logDifference diffLogger) { _ = "STUB: not implemented"; return }

// AssertEquals compares two arbitrary values and performs a comparison. If the
// comparison fails, a fatal error is raised that will fail the test
func AssertEquals(t *testing.T, expected, actual any) { _ = "STUB: not implemented"; return }

// CheckEquals is similar to AssertEquals, except with a non-fatal error
func CheckEquals(t *testing.T, expected, actual any) { _ = "STUB: not implemented"; return }

// AssertDeepEquals - like Equals - performs a comparison - but on more complex
// structures that requires deeper inspection
func AssertTypeEquals(t *testing.T, expected, actual any) { _ = "STUB: not implemented"; return }

// AssertDeepEquals - like Equals - performs a comparison - but on more complex
// structures that requires deeper inspection
func AssertDeepEquals(t *testing.T, expected, actual any) { _ = "STUB: not implemented"; return }

// CheckDeepEquals is similar to AssertDeepEquals, except with a non-fatal error
func CheckDeepEquals(t *testing.T, expected, actual any) { _ = "STUB: not implemented"; return }

func isByteArrayEquals(expectedBytes []byte, actualBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// AssertByteArrayEquals a convenience function for checking whether two byte arrays are equal
func AssertByteArrayEquals(t *testing.T, expectedBytes []byte, actualBytes []byte) {
	_ = "STUB: not implemented"
	return
}

// CheckByteArrayEquals a convenience function for silent checking whether two byte arrays are equal
func CheckByteArrayEquals(t *testing.T, expectedBytes []byte, actualBytes []byte) {
	_ = "STUB: not implemented"
	return
}

// isJSONEquals is a utility function that implements JSON comparison for AssertJSONEquals and
// CheckJSONEquals.
func isJSONEquals(t *testing.T, expectedJSON string, actual any) bool {
	_ = "STUB: not implemented"
	return false
}

// We can't use green() here because %#v prints prettyExpected as a byte array literal, which
// is... unhelpful. Converting it to a string first leaves "\n" uninterpreted for some reason.

// We can't use yellow() for the same reason.

// AssertJSONEquals serializes a value as JSON, parses an expected string as JSON, and ensures that
// both are consistent. If they aren't, the expected and actual structures are pretty-printed and
// shown for comparison.
//
// This is useful for comparing structures that are built as nested map[string]any values,
// which are a pain to construct as literals.
func AssertJSONEquals(t *testing.T, expectedJSON string, actual any) {
	_ = "STUB: not implemented"
	return
}

// CheckJSONEquals is similar to AssertJSONEquals, but nonfatal.
func CheckJSONEquals(t *testing.T, expectedJSON string, actual any) {
	_ = "STUB: not implemented"
	return
}

// AssertNoErr is a convenience function for checking whether an error value is
// an actual error
func AssertNoErr(t *testing.T, e error) { _ = "STUB: not implemented"; return }

// AssertErr is a convenience function for checking whether an error value is
// nil
func AssertErr(t *testing.T, e error) { _ = "STUB: not implemented"; return }

// AssertErrIs is a convenience function for checking whether an error value is
// target one
func AssertErrIs(t *testing.T, e error, target error) { _ = "STUB: not implemented"; return }

// CheckNoErr is similar to AssertNoErr, except with a non-fatal error
func CheckNoErr(t *testing.T, e error) { _ = "STUB: not implemented"; return }

// CheckErr is similar to AssertErr, except with a non-fatal error. If expected
// errors are passed, this function also checks that an error in e's tree is
// assignable to one of them. The tree consists of e itself, followed by the
// errors obtained by repeatedly calling Unwrap.
//
// CheckErr panics if expected contains anything other than non-nil pointers to
// either a type that implements error, or to any interface type.
func CheckErr(t *testing.T, e error, expected ...any) { _ = "STUB: not implemented"; return }

// AssertIntLesserOrEqual verifies that first value is lesser or equal than second values
func AssertIntLesserOrEqual(t *testing.T, v1 int, v2 int) { _ = "STUB: not implemented"; return }

// AssertIntGreaterOrEqual verifies that first value is greater or equal than second values
func AssertIntGreaterOrEqual(t *testing.T, v1 int, v2 int) { _ = "STUB: not implemented"; return }
