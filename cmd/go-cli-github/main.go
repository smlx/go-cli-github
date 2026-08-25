// Package main implements the command-line interface of a server.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

// CLI represents the command-line interface.
type CLI struct {
	Version kong.VersionFlag `kong:"help='Print version information'"`
	Serve   ServeCmd         `kong:"cmd,default=1,help='(default) Example serve command'"`
}

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	verJSON, err := formatVersionJSON()
	if err != nil {
		panic(err)
	}
	// parse CLI config
	cli := CLI{}
	kctx := kong.Parse(&cli,
		kong.UsageOnError(),
		kong.BindFor(ctx),
		kong.Vars{
			"version": verJSON,
		},
	)
	// execute CLI
	kctx.FatalIfErrorf(kctx.Run())
}
