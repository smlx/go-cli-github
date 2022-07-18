package main

import (
	"fmt"

	"github.com/smlx/go-cli-github/internal/server"
)

// ServeCmd represents the `serve` command.
type ServeCmd struct{}

// Run the serve command.
func (*ServeCmd) Run() error {
	fmt.Println(server.Serve())
	return nil
}
