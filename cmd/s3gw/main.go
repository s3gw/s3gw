package main

import (
	"fmt"
	"os"

	"github.com/s3gw/s3gw/internal/cmd"
)

func main() {
	if err := cmd.GetRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
