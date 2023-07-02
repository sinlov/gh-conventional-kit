package gh_conventional_kit

import (
	_ "embed"
	_ "github.com/aymerick/raymond"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/resource/template_file"
	"github.com/sinlov/gh-conventional-kit/utils/filepath_plus"
	"path/filepath"
)

//go:embed package.json
var PackageJson string

//go:embed resource/template_file/conventional_readme.md
var conventionalReadme string

func TemplateConventionalReadme(readme template_file.ConventionalConfig) (string, error) {
	return template_file.Render(conventionalReadme, readme)
}

var (
	//go:embed resource/template_file/versionrc.json
	versionrc string
)

func TemplateGitRootWalk(config template_file.ConventionalConfig, gitRootDir string) error {
	walkList := []template_file.ConventionalFile{
		{
			Name:         ".versionrc",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{gitRootDir, ".versionrc"},
			Content:      versionrc,
		},
	}
	for _, cFile := range walkList {
		tPath := filepath.Join(cFile.FullPaths...)
		err := filepath_plus.CheckOrCreateFileWithStringFast(tPath, cFile.Content)
		if err != nil {
			return err
		}
		slog.Debugf("TemplateGitRootWalk create file: %s", tPath)
	}
	return nil
}

var (
	//go:embed resource/template_file/github_template/pull_request_template.md
	pullRequestTemplate string

	//go:embed resource/template_file/github_template/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md
	codeOfConduct string

	//go:embed resource/template_file/github_template/CONTRIBUTING_DOC/CONTRIBUTING.md
	contributing string

	//go:embed resource/template_file/github_template/ISSUE_TEMPLATE/bug_report.md
	bugReport string

	//go:embed resource/template_file/github_template/ISSUE_TEMPLATE/feature_request.md
	featureRequest string

	//go:embed resource/template_file/github_template/ISSUE_TEMPLATE/help_wanted.md
	helpWanted string

	//go:embed resource/template_file/github_template/ISSUE_TEMPLATE/question.md
	question string
)

func TemplateGithubDotWalk(config template_file.ConventionalConfig, targetRootDir string) error {
	walkList := []template_file.ConventionalFile{
		{
			Name:         "pull_request_template.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "pull_request_template.md"},
			Content:      pullRequestTemplate,
		},
		{
			Name:         "CODE_OF_CONDUCT.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "CODE_OF_CONDUCT.md"},
			Content:      codeOfConduct,
		},
		{
			Name:         "CONTRIBUTING.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "CONTRIBUTING.md"},
			Content:      contributing,
		},
		{
			Name:         "bug_report.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "bug_report.md"},
			Content:      bugReport,
		},
		{
			Name:         "feature_request.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "feature_request.md"},
			Content:      featureRequest,
		},
		{
			Name:         "help_wanted.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "help_wanted.md"},
			Content:      helpWanted,
		},
		{
			Name:         "question.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "question.md"},
			Content:      question,
		},
	}
	for _, cFile := range walkList {
		tPath := filepath.Join(cFile.FullPaths...)
		err := filepath_plus.CheckOrCreateFileWithStringFast(tPath, cFile.Content)
		if err != nil {
			return err
		}
		slog.Debugf("TemplateGithubDotWalk create file: %s", tPath)
	}
	return nil
}
