package main

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
)

var Version = "dev"
var Revision = "dev"

type CLI struct {
	Version VersionFlag `short:"v" help:"Print version information and quit"`
	File    string      `arg:"" name:"file" help:"MP3 file path" type:"path"`
}

type VersionFlag bool

func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func RunCLI(args []string) error {
	cli := CLI{}
	parser, err := kong.New(&cli, kong.Vars{
		"version": fmt.Sprintf("%s (%s)", Version, Revision),
	})
	if err != nil {
		return err
	}
	_, err = parser.Parse(args)
	if err != nil {
		return err
	}
	app := New(&cli)
	err = app.Run(context.Background())
	if err != nil {
		return err
	}
	return nil
}