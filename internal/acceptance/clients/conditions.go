package clients

import (
	"testing"
)

// RequiredSystemScope will restrict a test to only be run by system scope.
func RequiredSystemScope(t *testing.T) { _ = "STUB: not implemented"; return }

// RequireAdmin will restrict a test to only be run by admin users.
func RequireAdmin(t *testing.T) { _ = "STUB: not implemented"; return }

// RequireNonAdmin will restrict a test to only be run by non-admin users.
func RequireNonAdmin(t *testing.T) { _ = "STUB: not implemented"; return }

// RequireLong will ensure long-running tests can run.
func RequireLong(t *testing.T) { _ = "STUB: not implemented"; return }

func getReleaseFromEnv(t *testing.T) string { _ = "STUB: not implemented"; return "" }

// SkipRelease will have the test be skipped on a certain release.
// Releases are named such as 'stable/dalmatian', master, etc.
func SkipRelease(t *testing.T, release string) { _ = "STUB: not implemented"; return }

// SkipReleasesBelow will have the test be skipped on releases below a certain one.
// Releases are named such as 'stable/dalmatian', master, etc.
func SkipReleasesBelow(t *testing.T, release string) { _ = "STUB: not implemented"; return }

// SkipReleasesAbove will have the test be skipped on releases above a certain one.
// The test is always skipped on master release.
// Releases are named such as 'stable/dalmatian', master, etc.
func SkipReleasesAbove(t *testing.T, release string) { _ = "STUB: not implemented"; return }

func isReleaseNumeral(release string) bool { _ = "STUB: not implemented"; return false }

// IsCurrentAbove will return true on releases above a certain one.
// The result is always true on master release.
// Releases are named such as 'stable/dalmatian', master, etc.
func IsCurrentAbove(t *testing.T, release string) bool { _ = "STUB: not implemented"; return false }

// Assume master is always too new

// Numeral releases are always newer than non-numeral ones

// IsCurrentBelow will return true on releases below a certain one.
// The result is always false on master release.
// Releases are named such as 'stable/dalmatian', master, etc.
func IsCurrentBelow(t *testing.T, release string) bool { _ = "STUB: not implemented"; return false }

// Assume master is always too new

// Numeral releases are always newer than non-numeral ones
