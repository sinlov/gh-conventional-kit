package subcommand_template

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/gookit/color"
	"github.com/sinlov/gh-conventional-kit"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/resource/template_file"
	"github.com/sinlov/gh-conventional-kit/utils/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/utils/git_tools"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

const commandName = "template"

var commandEntry *TemplateCommand

type TemplateCommand struct {
	isDebug            bool
	GitRootPath        string
	Remote             string
	LocalGitRemoteInfo *git_tools.GitRemoteInfo
	LocalGitBranch     string

	TargetFile   string
	TargetFolder string
	LanguageSet  []string
	NoVersionRc  bool
	BadgeConfig  *constant.BadgeConfig
}

func (n *TemplateCommand) Exec() error {

	readmeAppendHead, err := common_subcommand.BadgeByConfigWithMarkdown(
		n.BadgeConfig,
		n.LocalGitRemoteInfo.User,
		n.LocalGitRemoteInfo.Repo,
		n.LocalGitBranch,
	)
	if err != nil {
		return err
	}

	conventionalConfig := template_file.ConventionalConfig{
		GitOwnerName: n.LocalGitRemoteInfo.User,
		GitRepoName:  n.LocalGitRemoteInfo.Repo,
	}
	readmeAppendConventional, err := gh_conventional_kit.TemplateConventionalReadme(conventionalConfig)
	if err != nil {
		return err
	}
	var sb strings.Builder
	sb.WriteString(readmeAppendHead)
	sb.WriteString("\n")
	sb.WriteString(readmeAppendConventional)
	if len(n.LanguageSet) > 0 {
		languageConventional := gh_conventional_kit.LanguageConventional()
		for _, l := range n.LanguageSet {
			_, ok := languageConventional[l]
			if ok {
				sb.WriteString("\n")
				sb.WriteString(languageConventional[l])
				sb.WriteString("\n")
			} else {
				return fmt.Errorf("tempalte not support language: %s", l)
			}
		}
	}
	sb.WriteString("\n\n")

	if command.CmdGlobalEntry().DryRun {
		color.Greenf("template append head at path: %s \n", n.TargetFile)
		color.Grayf("%s\n", sb.String())
		return nil
	}

	err = gh_conventional_kit.TemplateGitRootWalk(conventionalConfig, n.GitRootPath)
	if err != nil {
		return err
	}

	if len(n.LanguageSet) > 0 {
		conventionalTemplate := gh_conventional_kit.LanguageConventionalTemplate()
		for _, l := range n.LanguageSet {
			_, ok := conventionalTemplate[l]
			if ok {
				fun := conventionalTemplate[l]
				err = fun(conventionalConfig, n.TargetFolder)
				if err != nil {
					return err
				}
				slog.Debugf("-> finish add template at language: %s", l)
			}
		}
	}
	slog.Infof("-> finish at template at: %s", n.TargetFolder)

	err = filepath_plus.AppendFileHead(n.TargetFile, []byte(sb.String()))
	if err != nil {
		return err
	}
	slog.Infof("-> finish add template at: %s", n.TargetFile)
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "gitRootFolder",
			Usage: "set git root folder, defaults is git_tools root path, value ''",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "remote",
			Usage: "set git remote name, defaults is: origin",
			Value: "origin",
		},
		&cli.StringFlag{
			Name:  "targetFile",
			Usage: "set conventional entrance file defaults is README.md",
			Value: "README.md",
		},
		&cli.StringFlag{
			Name:  "targetFolder",
			Usage: "set conventional folder, defaults is: .github",
			Value: ".github",
		},
		&cli.BoolFlag{
			Name:  "noVersionRc",
			Usage: "set no .versionrc, defaults is: false",
		},
		&cli.StringSliceFlag{
			Name:  "language",
			Usage: "set language, support : en-US,zh-CN",
			Value: cli.NewStringSlice("en-US"),
		},
	}
}

func withEntry(c *cli.Context) (*TemplateCommand, error) {
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
	_, err := git_tools.IsPathGitManagementRoot(gitRootFolder)
	if err != nil {
		return nil, err
	}
	fistRemoteInfo, err := git_tools.RepositoryFistRemoteInfo(gitRootFolder, remote)
	if err != nil {
		return nil, err
	}
	branchByPath, err := git_tools.RepositoryNowBranchByPath(gitRootFolder)
	if err != nil {
		return nil, err
	}

	targetFile := c.String("targetFile")
	targetFile = filepath.Join(gitRootFolder, targetFile)
	targetFolder := c.String("targetFolder")
	targetFolder = filepath.Join(gitRootFolder, targetFolder)

	return &TemplateCommand{
		isDebug:            globalEntry.Verbose,
		GitRootPath:        gitRootFolder,
		Remote:             remote,
		LocalGitRemoteInfo: fistRemoteInfo,
		LocalGitBranch:     branchByPath,

		TargetFile:   targetFile,
		TargetFolder: targetFolder,
		NoVersionRc:  c.Bool("noVersionRc"),
		LanguageSet:  c.StringSlice("language"),
		BadgeConfig:  constant.BindBadgeConfig(c),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)
	entry, err := withEntry(c)
	commandEntry = entry
	if err != nil {
		return err
	}
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "add conventional template at .github and try add badge at README.md (can change)",
			Action: action,
			Flags:  urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.BadgeFlags()),
		},
	}
}
