package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/martinlindhe/trasher/osx"
)

func main() {

	forcePtr := flag.Bool("f", false, "force (noop)")
	recursePtr := flag.Bool("r", false, "recurse (noop)")
	verbosePtr := flag.Bool("v", false, "verbose")
	flag.Parse()

	if *forcePtr {
		// NOOP
	}

	if *recursePtr {
		// NOOP, since we just move each argument
		os.Exit(1)
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: trasher <file>\n")
	}

	args := flag.Args()

	trashPath := osx.CurrentTrashPath()

	for _, arg := range args {
		if !Exists(arg) {
			fmt.Printf("File not found: %s\n", arg)
			continue
		}

		cnt := 1
		base := filepath.Base(arg)
		tname := filepath.Join(trashPath, base)

		for Exists(tname) {
			// trashed file exists, come up with a new name
			tname = filepath.Join(trashPath, base+fmt.Sprintf(" (copy %d)", cnt))
			cnt++
		}

		if *verbosePtr {
			fmt.Printf("Trashing %s => %s\n", arg, tname)
		}

		err := os.Rename(arg, tname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
