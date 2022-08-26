package main

import (
	"os"

	"gitlab.com/tokene/nonce-auth-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
