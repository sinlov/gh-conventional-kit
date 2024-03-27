package subcommand_template

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/command/common_subcommand"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/embed_operator"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"github.com/sinlov/gh-conventional-kit/internal/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli/cli_exit_urfave"
	"github.com/sinlov/gh-conventional-kit/resource"
	"github.com/sinlov/gh-conventional-kit/resource/contributing_doc"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const commandName = "template"

var commandEntry *TemplateCommand

type TemplateCommand struct {
	isDebug            bool
	GitRootPath        string
	Remote             string
	LocalGitRemoteInfo *git_info.GitRemoteInfo
	LocalGitBranch     string

	CoverageTargetFolderFile bool
	TargetFile               string
	ConventionalTargetFolder string
	LanguageSet              []string
	BadgeConfig              *constant.BadgeConfig
}

var (
	maintainResourceActionFileList = []string{
		resource.KeyPullRequestTemplate,
		resource.KeyDependabotConfig,
	}
)

func (n *TemplateCommand) Exec() error {

	readmeAppendHead, err := common_subcommand.BadgeByConfigWithMarkdown(
		n.BadgeConfig,
		n.LocalGitRemoteInfo.User,
		n.LocalGitRemoteInfo.Repo,
		n.LocalGitBranch,
	)
	if err != nil {
		slog.Errorf(err, "BadgeByConfigWithMarkdown")
		return cli_exit_urfave.ErrMsg(err, "BadgeByConfigWithMarkdown")
	}

	var sb strings.Builder
	sb.WriteString(readmeAppendHead)
	sb.WriteString("\n")

	conventionalConfig := contributing_doc.ConventionalTitleConfig{
		GitOwnerName: n.LocalGitRemoteInfo.User,
		GitRepoName:  n.LocalGitRemoteInfo.Repo,
	}

	conventionalTitleEmbed, err := embed_source.GetResourceByLanguage(resource.GroupResource,
		path.Join(resource.DirNameContributingDoc, resource.KeyConventionalReadmeTitle), constant.LangEnUS)
	if err != nil {
		slog.Errorf(err, "GetResourceByLanguage")
		return cli_exit_urfave.Err(err)
	}

	ccTitleRenderRes, err := conventionalTitleEmbed.Render(conventionalConfig)
	if err != nil {
		slog.Errorf(err, "Render")
		return cli_exit_urfave.Err(err)
	}
	slog.Debugf("-> conventionalTitleEmbed render: %s", ccTitleRenderRes)
	sb.WriteString(ccTitleRenderRes)

	if len(n.LanguageSet) > 0 {

		for _, l := range n.LanguageSet {
			ccReadmeEmbed, errCCReadMe := embed_source.GetResourceByLanguage(resource.GroupResource,
				path.Join(resource.DirNameContributingDoc, resource.KeyConventionalReadmeI18n), l)
			if errCCReadMe != nil {
				slog.Warnf("get conventional readme i18n err: %v", errCCReadMe)
				continue
			}
			ccReadmeContent, errCCReadMe := ccReadmeEmbed.Raw()
			if errCCReadMe != nil {
				slog.Warnf("get conventional readme i18n Raw err: %v", errCCReadMe)
				continue
			}
			sb.WriteString("\n")
			sb.WriteString(string(ccReadmeContent))
			sb.WriteString("\n")
		}
	}
	sb.WriteString("\n\n")

	if command.CmdGlobalEntry().DryRun {
		color.Bluef("-> dry run, not add to file: %s\n\n", n.TargetFile)

		color.Greenf("template append head at path: %s \n", n.TargetFile)
		color.Grayf("%s\n", sb.String())
		return nil
	}

	resIssues, err := embed_source.GetResourceGroupByDir(resource.GroupResource, path.Join(resource.DirGithubAction, resource.KeyIssueTemplate))
	if err != nil {
		slog.Errorf(err, "GetResourceGroupByDir")
		return cli_exit_urfave.ErrMsg(err, "GetResourceGroupByDir")
	}
	err = embed_operator.WriteFileByEmbedResources(resIssues,
		n.GitRootPath, n.CoverageTargetFolderFile,
		path.Join(resource.GroupResource, resource.DirGithubAction),
		constant.PathGithubAction,
		"issue template")
	if err != nil {
		slog.Errorf(err, "WriteFileByEmbedResources")
		return cli_exit_urfave.ErrMsg(err, "WriteFileByEmbedResources")
	}
	slog.Infof("-> finish check issues template at: %s", constant.PathGithubAction)

	for _, maintainResourceFileItem := range maintainResourceActionFileList {
		maintainResEmbedItem, errMaintainResEmbed := embed_source.GetResourceByLanguage(resource.GroupResourceAction, maintainResourceFileItem, constant.LangEnUS)
		if errMaintainResEmbed != nil {
			slog.Errorf(errMaintainResEmbed, "GetResourceByLanguage for maintainResourceActionFileList %s", maintainResourceFileItem)
			return cli_exit_urfave.ErrMsg(errMaintainResEmbed, "GetResourceByLanguage for maintainResourceActionFileList")
		}
		errMaintainResourceWrite := embed_operator.WriteFileByEmbedResource(maintainResEmbedItem,
			n.GitRootPath, n.CoverageTargetFolderFile,
			path.Join(resource.GroupResource, resource.DirGithubAction),
			constant.PathGithubAction,
			maintainResourceFileItem)
		if errMaintainResourceWrite != nil {
			slog.Errorf(errMaintainResourceWrite, "WriteFileByEmbedResource for maintainResourceActionFileList %s", maintainResourceFileItem)
			return cli_exit_urfave.Err(errMaintainResourceWrite)
		}
	}

	for _, l := range n.LanguageSet {
		ccEmbedAsLang, errCCEmbedAsLang := embed_source.GetResourceListByLanguage(resource.GroupResourceAction, resource.KeyContributingDoc, l)
		if errCCEmbedAsLang != nil {
			slog.Errorf(errCCEmbedAsLang, "GetResourceByLanguage")
			return cli_exit_urfave.ErrMsg(errCCEmbedAsLang, "GetResourceByLanguage")
		}
		errWriteCCAsLang := embed_operator.WriteFileByEmbedResources(ccEmbedAsLang,
			n.GitRootPath, n.CoverageTargetFolderFile,
			path.Join(resource.GroupResource, resource.DirGithubAction),
			n.ConventionalTargetFolder,
			"contributing template",
		)
		if errWriteCCAsLang != nil {
			slog.Errorf(errWriteCCAsLang, "WriteFileByEmbedResources")
			return cli_exit_urfave.ErrMsg(errWriteCCAsLang, "WriteFileByEmbedResources")
		}
	}

	slog.Infof("-> finish add contributing template at: %s", n.ConventionalTargetFolder)

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
			Name:  constant.CliNameGitRootFolder,
			Usage: "set add target git root folder, defaults is cli run path",
			Value: "",
		},
		&cli.StringFlag{
			Name:  constant.CliNameGitRemote,
			Usage: "set git remote name, defaults is: origin",
			Value: "origin",
		},
		&cli.StringFlag{
			Name:  "targetFile",
			Usage: "set conventional entrance file defaults is README.md",
			Value: "README.md",
		},
		&cli.StringFlag{
			Name:  "conventional-targetFolder",
			Usage: "set conventional folder",
			Value: constant.PathGithubAction,
		},
		&cli.BoolFlag{
			Name:  "coverage-folder-file",
			Usage: "coverage folder file under targetFolder, does not affect files that are not in the template",
			Value: false,
		},
		&cli.StringSliceFlag{
			Name:  "language",
			Usage: fmt.Sprintf("set language, support : %v", constant.SupportLanguage()),
			Value: cli.NewStringSlice(constant.LangEnUS, constant.LangZhCN),
		},
	}
}

func withEntry(c *cli.Context) (*TemplateCommand, error) {
	globalEntry := command.CmdGlobalEntry()

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
	branchByPath, err := git_info.RepositoryNowBranchByPath(gitRootFolder)
	if err != nil {
		return nil, err
	}

	targetFile := c.String("targetFile")
	targetFile = filepath.Join(gitRootFolder, targetFile)
	conventionalTargetFolder := c.String("conventional-targetFolder")

	languageSet := c.StringSlice("language")
	if len(languageSet) > 0 {
		for _, lang := range languageSet {
			if !constant.StrInArr(lang, constant.SupportLanguage()) {
				return nil, fmt.Errorf("not support language: %s, please check flag --language", lang)
			}
		}
	}

	return &TemplateCommand{
		isDebug:            globalEntry.Verbose,
		GitRootPath:        gitRootFolder,
		Remote:             remote,
		LocalGitRemoteInfo: fistRemoteInfo,
		LocalGitBranch:     branchByPath,

		CoverageTargetFolderFile: c.Bool("coverage-folder-file"),
		TargetFile:               targetFile,
		ConventionalTargetFolder: conventionalTargetFolder,
		LanguageSet:              languageSet,
		BadgeConfig:              constant.BindBadgeConfig(c),
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
