package lib

import (
	"fmt"
)

// Trash moves a file to the trash folder
func Trash(fileName string, verbose bool) error {
	if !fileExists(fileName) {
		return fmt.Errorf("file not found: %s", fileName)
	}
	if verbose {
		fmt.Printf("Trashing %s\n", fileName)
	}
	return runInteractiveCommand("gio", "trash", fileName)
}
