package subcommand_markdown

import (
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
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
	Branch   string

	BadgeConfig *constant.BadgeConfig
}

func (n *MarkdownCommand) Exec() error {

	badgeConfig := n.BadgeConfig
	err := common_subcommand.PrintBadgeByConfigWithMarkdown(badgeConfig, n.UserName, n.RepoName, n.Branch)
	if err != nil {
		return err
	}
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
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
			Name:    "branch",
			Aliases: []string{"b"},
			Value:   "main",
			Usage:   "Repository at git branch",
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
		isDebug:     globalEntry.Verbose,
		GitHost:     c.String("host"),
		UserName:    c.String("user"),
		RepoName:    c.String("repo"),
		Branch:      c.String("branch"),
		BadgeConfig: constant.BindBadgeConfig(c),
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
			Flags:  urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.BadgeFlags()),
		},
	}
}
