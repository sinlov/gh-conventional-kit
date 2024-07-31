package cli

import (
	"fmt"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_action"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_badge"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_markdown"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_template"
	"github.com/sinlov/gh-conventional-kit/internal/pkgJson"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
	"runtime"
	"time"
)

const (
	copyrightStartYear = "2023"
	defaultExitCode    = 1
)

func NewCliApp(buildId string) *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)
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
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, build id: %s, run on %s %s",
		copyrightStartYear, year, jsonAuthor.Name, runtime.Version(), buildId, runtime.GOOS, runtime.GOARCH)
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
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_action.Command())

	app.Commands = appCommands

	return app
}
