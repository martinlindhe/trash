package main

import (
	"flag"
	"fmt"
	"os"

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

	for _, arg := range args {

		err := osx.Trash(arg, *verbosePtr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
