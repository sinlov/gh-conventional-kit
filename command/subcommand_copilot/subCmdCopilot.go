package subcommand_copilot

import (
	"fmt"
	"os"
	"path"

	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli/cli_exit_urfave"
	"github.com/sinlov/gh-conventional-kit/internal/embed_operator"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"github.com/sinlov/gh-conventional-kit/resource"
	"github.com/urfave/cli/v2"
)

const commandName = "copilot"

var copilotCommandEntry *CopilotCommand

type CopilotCommand struct {
	isDebug bool

	GitRootPath              string
	CoverageTargetFolderFile bool
	LanguageSet              []string

	IsInitCommit bool
}

var maintainResourceCopilotFileList = []string{
	resource.KeyCopilotInstructionMd,
	resource.KeyGitCommitInstruction,
}

func (n *CopilotCommand) Exec() error {
	if n.IsInitCommit {
		errInitCommitConfig := n.initCommitConfig()
		if errInitCommitConfig != nil {
			slog.Errorf(errInitCommitConfig, "initCommitConfig")

			return cli_exit_urfave.Err(errInitCommitConfig)
		}
	}

	slog.Info("finish do github copilot config for git project")

	return nil
}

func (n *CopilotCommand) initCommitConfig() error {
	langSet := constant.LangEnUS
	if len(n.LanguageSet) > 0 {
		langSet = n.LanguageSet[0]
	}

	for _, maintainResourceFileItem := range maintainResourceCopilotFileList {
		maintainResEmbedItem, errMaintainResEmbed := embed_source.GetResourceByLanguage(
			resource.GroupCopilot,
			maintainResourceFileItem,
			langSet,
		)
		if errMaintainResEmbed != nil {
			slog.Errorf(
				errMaintainResEmbed,
				"GetResourceByLanguage for maintainResourceCopilotFileList %s",
				maintainResourceFileItem,
			)

			return errMaintainResEmbed
		}

		// to remove language set from path
		replacePath := path.Join(resource.GroupResource, resource.DirCopilot, langSet)

		errMaintainResourceWrite := embed_operator.WriteFileByEmbedResource(
			maintainResEmbedItem,
			n.GitRootPath, n.CoverageTargetFolderFile,
			replacePath,
			constant.PathGithubAction,
			maintainResourceFileItem)
		if errMaintainResourceWrite != nil {
			slog.Errorf(
				errMaintainResourceWrite,
				"WriteFileByEmbedResource for maintainResourceCopilotFileList %s",
				maintainResourceFileItem,
			)

			return cli_exit_urfave.Err(errMaintainResourceWrite)
		}
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  constant.CliNameGitRootFolder,
			Usage: "set add target git root folder, defaults is cli run path",
			Value: "",
		},
		&cli.BoolFlag{
			Name:  "coverage-folder-file",
			Usage: "coverage folder file under targetFolder, does not affect files that are not in the template",
			Value: false,
		},
		&cli.StringSliceFlag{
			Name:  "language",
			Usage: fmt.Sprintf("set language, support : %v", constant.SupportLanguage()),
			Value: cli.NewStringSlice(constant.LangEnUS),
		},

		&cli.BoolFlag{
			Name:  "init-commit",
			Usage: "Initialize github copilot commit message configuration commit, if set language, will use the first language as the commit message",
			Value: false,
		},
	}
}

func withEntry(c *cli.Context) (*CopilotCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	gitRootFolder := c.String(constant.CliNameGitRootFolder)
	if gitRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}

		gitRootFolder = dir
	}

	languageSet := c.StringSlice("language")
	if len(languageSet) > 0 {
		for _, lang := range languageSet {
			if !constant.StrInArr(lang, constant.SupportLanguage()) {
				return nil, fmt.Errorf(
					"not support language: %s, please check flag --language",
					lang,
				)
			}
		}
	}

	return &CopilotCommand{
		isDebug: globalEntry.Verbose,

		GitRootPath: gitRootFolder,

		CoverageTargetFolderFile: c.Bool("coverage-folder-file"),
		LanguageSet:              languageSet,

		IsInitCommit: c.Bool("init-commit"),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)

	entry, err := withEntry(c)
	if err != nil {
		return err
	}

	copilotCommandEntry = entry

	return copilotCommandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "add github copilot config for git project",
			Action: action,
			Flags:  flag(),
		},
	}
}
