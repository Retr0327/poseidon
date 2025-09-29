package path

import (
	"os"
	"path/filepath"
)

// ProjectRootDir returns the root directory of the Go project by locating the go.mod file.
func ProjectRootDir() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
