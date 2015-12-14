package osx

import (
	"os"
	"path/filepath"
)

// CurrentTrashPath returns the path of the current user's trash
func CurrentTrashPath() string {
	return filepath.Join(os.Getenv("HOME"), ".Trash")
}
