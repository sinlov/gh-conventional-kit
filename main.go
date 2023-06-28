//go:build !test

package main

import (
	_ "embed"
	"fmt"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_badge"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_markdown"
	"github.com/sinlov/gh-conventional-kit/utils/pkgJson"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	os "os"
)

//go:embed package.json
var packageJson string

func main() {
	pkgJson.InitPkgJsonContent(packageJson)
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkgJson.GetPackageJsonVersionGoStyle(false)
	app.Name = pkgJson.GetPackageJsonName()
	if pkgJson.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkgJson.GetPackageJsonHomepage())
	}
	app.Description = pkgJson.GetPackageJsonDescription()

	author := &cli.Author{
		Name:  pkgJson.GetPackageJsonAuthor().Name,
		Email: pkgJson.GetPackageJsonAuthor().Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	flags := urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	app.Flags = flags
	app.Before = command.GlobalBeforeAction
	app.Action = command.GlobalAction
	app.After = command.GlobalAfterAction

	var appCommands []*cli.Command
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_badge.Command())
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_markdown.Command())

	app.Commands = appCommands

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("please see help as: %s --help\n", app.Name)
		os.Exit(2)
	}
	if err := app.Run(args); nil != err {
		fmt.Printf("run err: %v\n", err)
	}
}
