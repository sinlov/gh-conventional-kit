package cli

import (
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_action"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_badge"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_markdown"
	"github.com/sinlov/gh-conventional-kit/command/subcommand_template"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/pkg_kit"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli"
	"github.com/urfave/cli/v2"
)

func NewCliApp(bdInfo pkg_kit.BuildInfo) *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = bdInfo.PgkNameString()
	app.Version = bdInfo.VersionString()

	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = "see: " + pkg_kit.GetPackageJsonHomepage()
	}

	app.Description = pkg_kit.GetPackageJsonDescription()
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	app.Copyright = bdInfo.String()
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
