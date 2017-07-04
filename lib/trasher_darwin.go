package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

// Trash moves a file to the trash folder
func Trash(fileName string, verbose bool) error {
	if !fileExists(fileName) {
		return fmt.Errorf("file not found: %s", fileName)
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
		fmt.Printf("Trashing %s\n", fileName)
	}
	return os.Rename(fileName, dstName)
}

// trashPath returns the path of the current user's trash
func trashPath() string {
	return filepath.Join(os.Getenv("HOME"), ".Trash")
}
