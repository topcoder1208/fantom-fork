package main

import (
	"fmt"
	"os"

	"github.com/topcoder1208/fantom-fork/cmd/mugambo/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
