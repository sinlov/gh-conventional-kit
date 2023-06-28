package subcommand_markdown

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/badges/shields_helper"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
)

const commandName = "markdown"

var commandEntry *MarkdownCommand

type MarkdownCommand struct {
	isDebug bool

	GitHost  string
	UserName string
	RepoName string

	noCommonBadges bool
}

func (n *MarkdownCommand) Exec() error {

	if !n.noCommonBadges {
		if n.UserName == "" || n.RepoName == "" {
			return fmt.Errorf("need set user and repo")
		}

		slog.Info("\nCommon badges:")
		fmt.Println(shields_helper.GithubTagLatestMarkdown(n.UserName, n.RepoName))
		fmt.Println(shields_helper.GitHubReleaseMarkdown(n.UserName, n.RepoName))
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-common",
			Value: false,
			Usage: "no badges common for this repo",
		},

		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Repository at git user name",
		},
		&cli.StringFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Repository at git repo",
		},

		&cli.StringFlag{
			Hidden: true,
			Name:   "host",
			Usage:  "Repository at git host",
			Value:  "github.com",
		},
	}
}

func withEntry(c *cli.Context) (*MarkdownCommand, error) {
	globalEntry := command.CmdGlobalEntry()
	return &MarkdownCommand{
		isDebug:  globalEntry.Verbose,
		GitHost:  c.String("host"),
		UserName: c.String("user"),
		RepoName: c.String("repo"),

		noCommonBadges: c.Bool("no-common"),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "generate markdown text for git project",
			Action: action,
			Flags:  flag(),
		},
	}
}
