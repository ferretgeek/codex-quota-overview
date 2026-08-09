package app

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReplacePathRootMasksOnlyPathsWithinRoot(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	home := filepath.Join(base, "home")
	inside := filepath.Join(home, "workspace", "accounts")
	outside := filepath.Join(base, "public", "accounts")

	masked := replacePathRoot(inside, home, "~")
	if strings.Contains(masked, home) || masked != filepath.Join("~", "workspace", "accounts") {
		t.Fatalf("inside path was not safely masked: %q", masked)
	}
	if got := replacePathRoot(outside, home, "~"); got != filepath.Clean(outside) {
		t.Fatalf("outside path changed unexpectedly: %q", got)
	}
}
