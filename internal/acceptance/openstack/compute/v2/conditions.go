package v2

import (
	"testing"
)

// RequireGuestAgent will restrict a test to only be run in
// environments that support the QEMU guest agent.
func RequireGuestAgent(t *testing.T) { _ = "STUB: not implemented"; return }

// RequireLiveMigration will restrict a test to only be run in
// environments that support live migration.
func RequireLiveMigration(t *testing.T) { _ = "STUB: not implemented"; return }
