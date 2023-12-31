package main

import (
	"fmt"
	"os"

	"github.com/metatechchain/go-metatechchain/cmd/metatechchain/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
