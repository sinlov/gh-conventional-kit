package subcommand_action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bar-counter/slog"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli/cli_exit_urfave"
	"github.com/sinlov/gh-conventional-kit/internal/embed_operator"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"github.com/sinlov/gh-conventional-kit/resource"
	"github.com/urfave/cli/v2"
)

const commandName = "action"

var commandEntry *ActionCommand

const (
	resourceLanguage = ""
)

var resourceActionDeployTagsFileList = []string{
	resource.KeyGithubActionVersionYml,
	resource.KeyGithubActionDeployTagsYml,
	resource.KeyGithubActionCIYml,
}

type ActionCommand struct {
	isDebug bool

	GitRootPath              string
	TargetActionFolder       string
	CoverageTargetFolderFile bool

	IsCiDeployTag bool
}

func (n *ActionCommand) Exec() error {
	if n.IsCiDeployTag {
		for _, maintainResourceFileItem := range resourceActionDeployTagsFileList {
			maintainResEmbedItem, errMaintainResEmbed := embed_source.GetResourceByLanguage(
				resource.GroupResourceActionWorkflowsDeployTags,
				maintainResourceFileItem,
				resourceLanguage,
			)
			if errMaintainResEmbed != nil {
				slog.Errorf(
					errMaintainResEmbed,
					"GetResource for resourceActionDeployTagsFileList %s",
					maintainResourceFileItem,
				)

				return errMaintainResEmbed
			}

			errMaintainResourceWrite := embed_operator.WriteFileByEmbedResource(
				maintainResEmbedItem,
				n.GitRootPath,
				n.CoverageTargetFolderFile,
				resource.GroupResourceActionWorkflowsDeployTags,
				filepath.Join(constant.PathGithubAction, constant.PathWorkflows),
				maintainResourceFileItem,
			)
			if errMaintainResourceWrite != nil {
				slog.Errorf(
					errMaintainResourceWrite,
					"WriteFileByEmbedResource for maintainResourceActionFileList %s",
					maintainResourceFileItem,
				)

				return cli_exit_urfave.Err(errMaintainResourceWrite)
			}
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
		&cli.StringFlag{
			Name:  "targetFolder",
			Usage: "change target folder",
			Value: filepath.Join(constant.PathGithubAction, constant.PathWorkflows),
		},
		&cli.BoolFlag{
			Name:  "coverage-folder-file",
			Usage: "coverage folder or file under targetFolder, does not affect files that are not in the template",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "ci-deploy-tag",
			Usage: "add sample deploy by tag",
			Value: false,
		},
	}
}

func withEntry(c *cli.Context) (*ActionCommand, error) {
	globalEntry := command.CmdGlobalEntry()

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

	return &ActionCommand{
		isDebug:     globalEntry.Verbose,
		GitRootPath: gitRootFolder,

		CoverageTargetFolderFile: c.Bool("coverage-folder-file"),
		TargetActionFolder:       c.String("targetFolder"),

		IsCiDeployTag: c.Bool("ci-deploy-tag"),
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
			Usage:  "fast add github action workflow",
			Action: action,
			Flags:  flag(),
		},
	}
}
