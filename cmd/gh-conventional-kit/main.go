//go:build !test

package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov/gh-conventional-kit"
	"github.com/sinlov/gh-conventional-kit/cmd/cli"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"github.com/sinlov/gh-conventional-kit/internal/pkgJson"
	os "os"
)

func main() {
	pkgJson.InitPkgJsonContent(gh_conventional_kit.PackageJson)
	embed_source.RegisterSettings(embed_source.DefaultFunctions)
	app := cli.NewCliApp()

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(2)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
