package main

import (
	"fmt"
	"os"

	"github.com/MugamboBC/go-mugambo/cmd/mugambo/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
