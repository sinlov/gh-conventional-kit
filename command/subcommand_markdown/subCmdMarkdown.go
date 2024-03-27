package subcommand_markdown

import (
	"fmt"
	"github.com/bar-counter/slog"
	giturls "github.com/chainguard-dev/git-urls"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

const commandName = "markdown"

var commandEntry *MarkdownCommand

type MarkdownCommand struct {
	isDebug bool

	UserName string
	RepoName string
	Branch   string

	BadgeConfig *constant.BadgeConfig
}

func (n *MarkdownCommand) Exec() error {

	err := common_subcommand.PrintBadgeByConfigWithMarkdown(n.BadgeConfig, n.UserName, n.RepoName, n.Branch)
	if err != nil {
		return err
	}
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  constant.CliNameGitRootFolder,
			Usage: "set add target git root folder, defaults is cli run path, this will ignore --user and --repo or use arg0 to input git url",
			Value: "",
		},
		&cli.StringFlag{
			Name:  constant.CliNameGitRemote,
			Usage: "set git remote name, defaults is origin, this will ignore --user and --repo or use arg0 to input git url",
			Value: "origin",
		},

		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Repository at git user name, can cover by arg0",
		},
		&cli.StringFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Repository at git repo, can cover by arg0",
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

	user := c.String("user")
	repo := c.String("repo")

	if user == "" || repo == "" {
		if c.Args().Len() > 0 {
			gitCmdUrl := c.Args().First()
			slog.Debugf("input gitCmdUrl: %s", gitCmdUrl)
			gitUrl, err := giturls.Parse(gitCmdUrl)
			if err != nil {
				return nil, fmt.Errorf("parse gitCmdUrl: %s error: %s", gitCmdUrl, err)
			}
			urlPath := gitUrl.Path
			if gitUrl.Scheme == "ssh" {
				splitPath := strings.Split(urlPath, "/")
				if len(splitPath) < 2 {
					return nil, fmt.Errorf("gitCmdUrl: %s not find user or repo", gitCmdUrl)
				}

				user = splitPath[0]
				repo = splitPath[1]
			} else {
				splitPath := strings.Split(urlPath, "/")
				if len(splitPath) < 3 {
					return nil, fmt.Errorf("gitCmdUrl: %s not find user or repo", gitCmdUrl)
				}

				user = splitPath[1]
				repo = splitPath[2]
			}
		} else {
			slog.Debug("try find out at git repo")
			remote := c.String(constant.CliNameGitRemote)
			gitRootFolder := c.String(constant.CliNameGitRootFolder)
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
			user = fistRemoteInfo.User
			slog.Debugf("find git repo user: %s", user)
			repo = fistRemoteInfo.Repo
			slog.Debugf("find git repo: %s", repo)
		}
	}

	if user == "" {
		return nil, fmt.Errorf("need set --user, or can not find git repo user")
	}
	if repo == "" {
		return nil, fmt.Errorf("need set --repo, or can not find git repo")
	}

	return &MarkdownCommand{
		isDebug:     globalEntry.Verbose,
		UserName:    user,
		RepoName:    repo,
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
			Name:      commandName,
			Usage:     "generate markdown badge by program language or framework for this repo",
			UsageText: "markdown [command options] <gitUrl>\nmarkdown --golang https://github.com/sinlov/gh-conventional-kit",
			Action:    action,
			Flags:     urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.BadgeFlags()),
		},
	}
}
