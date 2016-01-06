package osx

import (
	"fmt"
	"os"
	"path/filepath"
)

// trashPath returns the path of the current user's trash
func trashPath() string {
	return filepath.Join(os.Getenv("HOME"), ".Trash")
}

// Trash moves a file to the trash folder
func Trash(fileName string, verbose bool) error {

	if !fileExists(fileName) {
		return fmt.Errorf("File not found: %s\n", fileName)
	}

	trashPath := trashPath()
	baseName := filepath.Base(fileName)
	dstName := filepath.Join(trashPath, baseName)

	i := 1
	for fileExists(dstName) {
		// come up with a new name
		dstName = filepath.Join(trashPath, baseName+fmt.Sprintf(" (copy %d)", i))
		i++
	}

	if verbose {
		fmt.Printf("Trashing %s => %s\n", fileName, dstName)
	}

	return os.Rename(fileName, dstName)
}

// reports whether the named file exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
