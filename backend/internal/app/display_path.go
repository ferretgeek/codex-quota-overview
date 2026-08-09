package app

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// DisplayPath removes the current user's home directory from UI and diagnostic
// strings. Internal filesystem operations continue to use the original path.
func DisplayPath(path string) string {
	home, err := os.UserHomeDir()
	if err != nil || strings.TrimSpace(home) == "" {
		return filepath.Clean(path)
	}
	replacement := "~"
	if runtime.GOOS == "windows" {
		replacement = "%USERPROFILE%"
	}
	return replacePathRoot(path, home, replacement)
}

func replacePathRoot(path string, root string, replacement string) string {
	cleanPath := filepath.Clean(path)
	cleanRoot := filepath.Clean(root)
	relative, err := filepath.Rel(cleanRoot, cleanPath)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(os.PathSeparator)) {
		return cleanPath
	}
	if relative == "." {
		return replacement
	}
	return filepath.Join(replacement, relative)
}
