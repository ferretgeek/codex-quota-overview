package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveAppDirFromRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "web"), 0o755); err != nil {
		t.Fatal(err)
	}
	if got := resolveAppDir(root); got != root {
		t.Fatalf("resolveAppDir(%q) = %q, want %q", root, got, root)
	}
}

func TestResolveAppDirFromBackendDirectory(t *testing.T) {
	root := t.TempDir()
	backend := filepath.Join(root, "backend")
	for _, path := range []string{backend, filepath.Join(root, "web")} {
		if err := os.Mkdir(path, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if got := resolveAppDir(backend); got != root {
		t.Fatalf("resolveAppDir(%q) = %q, want %q", backend, got, root)
	}
}
