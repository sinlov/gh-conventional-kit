package subcommand_badge

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/utils/git_tools"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
)

const commandName = "badge"

var commandEntry *BadgeCommand

type BadgeCommand struct {
	isDebug     bool
	GitRootPath string
	Remote      string

	NoMarkdown  bool
	BadgeConfig *constant.BadgeConfig
}

func (n *BadgeCommand) Exec() error {
	fistRemoteInfo, err := git_tools.RepositoryFistRemoteInfo(n.GitRootPath, n.Remote)
	if err != nil {
		return err
	}
	branchByPath, err := git_tools.RepositoryNowBranchByPath(n.GitRootPath)
	if err != nil {
		return err
	}
	if n.NoMarkdown {
		err = common_subcommand.PrintBadgeByConfig(
			n.BadgeConfig,
			fistRemoteInfo.User,
			fistRemoteInfo.Repo,
			branchByPath,
		)
		if err != nil {
			return err
		}
		return nil
	}

	err = common_subcommand.PrintBadgeByConfigWithMarkdown(
		n.BadgeConfig,
		fistRemoteInfo.User, fistRemoteInfo.Repo, branchByPath,
	)
	if err != nil {
		return err
	}
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "gitRootFolder",
			Usage: "set add badge target git root folder, defaults is git_tools root path, value ''",
			Value: "",
		},
		&cli.BoolFlag{
			Name:  "no-markdown",
			Usage: "no add markdown badge",
		},
		&cli.StringFlag{
			Name:  "remote",
			Usage: "set git remote name, defaults is origin",
			Value: "origin",
		},
	}
}

func withEntry(c *cli.Context) (*BadgeCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	gitRootFolder := c.String("gitRootFolder")
	if gitRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}
		gitRootFolder = dir
	}
	_, err := git_tools.IsPathGitManagementRoot(gitRootFolder)
	if err != nil {
		return nil, err
	}

	return &BadgeCommand{
		isDebug:     globalEntry.Verbose,
		GitRootPath: gitRootFolder,
		Remote:      c.String("remote"),

		NoMarkdown:  c.Bool("no-markdown"),
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
			Usage:  "add badge at github project, use this command must in git workspace root folder, or set --gitRootFolder",
			Action: action,
			Flags:  urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.BadgeFlags()),
		},
	}
}
