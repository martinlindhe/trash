package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/trasher/osx"
)

var (
	files   = kingpin.Arg("files", "Files to remove.").Strings()
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	recurse = kingpin.Flag("recurse", "Recurse (noop)").Short('r').Bool()
	force   = kingpin.Flag("force", "Force (noop)").Short('f').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	for _, f := range *files {

		err := osx.Trash(f, *verbose)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
