package subcommand_badge

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/utils/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

const commandName = "badge"

var commandEntry *BadgeCommand

type BadgeCommand struct {
	isDebug            bool
	GitRootPath        string
	Remote             string
	LocalGitRemoteInfo *git_info.GitRemoteInfo
	LocalGitBranch     string

	TargetFile  string
	NoMarkdown  bool
	BadgeConfig *constant.BadgeConfig
}

func (n *BadgeCommand) Exec() error {

	if command.CmdGlobalEntry().DryRun {
		if n.NoMarkdown {
			err := common_subcommand.PrintBadgeByConfig(
				n.BadgeConfig,
				n.LocalGitRemoteInfo.User,
				n.LocalGitRemoteInfo.Repo,
				n.LocalGitBranch,
			)
			if err != nil {
				return err
			}
			return nil
		}

		err := common_subcommand.PrintBadgeByConfigWithMarkdown(
			n.BadgeConfig,
			n.LocalGitRemoteInfo.User,
			n.LocalGitRemoteInfo.Repo,
			n.LocalGitBranch,
		)
		if err != nil {
			return err
		}
		return nil
	}
	targetAppendHeadBadge := ""
	if n.NoMarkdown {
		targetAppendHead, errBadge := common_subcommand.BadgeByConfig(
			n.BadgeConfig,
			n.LocalGitRemoteInfo.User,
			n.LocalGitRemoteInfo.Repo,
			n.LocalGitBranch,
		)
		if errBadge != nil {
			return errBadge
		}
		targetAppendHeadBadge = targetAppendHead
	} else {
		targetAppendHead, errBadge := common_subcommand.BadgeByConfigWithMarkdown(
			n.BadgeConfig,
			n.LocalGitRemoteInfo.User,
			n.LocalGitRemoteInfo.Repo,
			n.LocalGitBranch,
		)
		if errBadge != nil {
			return errBadge
		}
		targetAppendHeadBadge = targetAppendHead
	}

	var sb strings.Builder
	sb.WriteString(targetAppendHeadBadge)
	sb.WriteString("\n")

	sb.WriteString("\n")

	err := filepath_plus.AppendFileHead(n.TargetFile, []byte(sb.String()))
	if err != nil {
		return err
	}
	slog.Infof("-> finish append file head at: %s", n.TargetFile)

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "gitRootFolder",
			Usage: "set add target git root folder, defaults is cli run path",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "remote",
			Usage: "set git remote name, defaults is origin",
			Value: "origin",
		},
		&cli.BoolFlag{
			Name:  "no-markdown",
			Usage: "no add markdown badge",
		},
		&cli.StringFlag{
			Name:  "targetFile",
			Usage: "badge append file head defaults is README.md",
			Value: "README.md",
		},
	}
}

func withEntry(c *cli.Context) (*BadgeCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	remote := c.String("remote")
	gitRootFolder := c.String("gitRootFolder")
	if gitRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}
		gitRootFolder = dir
	}
	_, err := git_info.IsPathGitManagementRoot(gitRootFolder)
	if err != nil {
		return nil, err
	}
	fistRemoteInfo, err := git_info.RepositoryFistRemoteInfo(gitRootFolder, remote)
	if err != nil {
		return nil, err
	}
	branchByPath, err := git_info.RepositoryNowBranchByPath(gitRootFolder)
	if err != nil {
		return nil, err
	}

	targetFile := c.String("targetFile")
	targetFile = filepath.Join(gitRootFolder, targetFile)

	return &BadgeCommand{
		isDebug:            globalEntry.Verbose,
		GitRootPath:        gitRootFolder,
		Remote:             remote,
		LocalGitRemoteInfo: fistRemoteInfo,
		LocalGitBranch:     branchByPath,

		TargetFile:  targetFile,
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
