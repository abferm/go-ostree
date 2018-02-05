package main

import (
	"fmt"
	"os"

	"github.com/abferm/go-ostree/ostree"
)

func main() {
	commit, err := ostree.GetBootedCommit()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	fmt.Printf("Booted commit hash: %s\n", commit)
	version, err := ostree.GetVersionTag(commit)
	if err != nil {
		fmt.Println("Commit has no version tag, using the first 6 characters of the commit hash instead.")
		version = commit[:6]
	}
	fmt.Printf("Version: %s\n", version)
}
