package cli

import (
	"fmt"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_badge"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_markdown"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_template"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/utils/pkgJson"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"runtime"
	"time"
)

func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkgJson.GetPackageJsonVersionGoStyle(false)
	app.Name = pkgJson.GetPackageJsonName()
	if pkgJson.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkgJson.GetPackageJsonHomepage())
	}
	app.Description = pkgJson.GetPackageJsonDescription()

	year := time.Now().Year()
	jsonAuthor := pkgJson.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, run on %s %s",
		constant.CopyrightStartYear, year, jsonAuthor.Name, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
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
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_template.Command())

	app.Commands = appCommands

	return app
}