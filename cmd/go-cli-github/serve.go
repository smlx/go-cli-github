package main

import (
	"fmt"
	"time"

	"github.com/smlx/go-cli-github/internal/server"
)

// ServeCmd represents the `serve` command.
type ServeCmd struct{}

// Run the serve command.
func (*ServeCmd) Run() error {
	fmt.Println(server.New(time.Now).Serve())
	return nil
}
