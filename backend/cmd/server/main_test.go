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

func TestLoopbackListenAddress(t *testing.T) {
	for _, address := range []string{"127.0.0.1:8787", "127.1.2.3:8787", "localhost:8787", "[::1]:8787"} {
		if !isLoopbackListenAddress(address) {
			t.Fatalf("expected %q to be accepted", address)
		}
	}
	for _, address := range []string{"0.0.0.0:8787", ":8787", "192.0.2.10:8787", "example.com:8787", "bad"} {
		if isLoopbackListenAddress(address) {
			t.Fatalf("expected %q to be rejected", address)
		}
	}
}
