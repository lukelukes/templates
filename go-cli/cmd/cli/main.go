package main

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Add AddCmd `cmd:"" help:"example cmd"`

	Path string `name:"path" short:"c" help:"example arg"`
}

type AddCmd = string

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("replaceme"),
		kong.Description("replaceme"),
		kong.UsageOnError(),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
