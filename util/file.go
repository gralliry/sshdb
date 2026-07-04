package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteFile writes data to a file, creating parent directories as needed.
func WriteFile(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create directory %s: %w", dir, err)
	}
	return os.WriteFile(path, data, perm)
}
