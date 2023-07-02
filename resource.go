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
		slog.Debugf("TemplateGitRootWalk check or create file: %s", tPath)
	}
	return nil
}

var languageConventional map[string]string

func LanguageConventional() map[string]string {
	if languageConventional == nil {
		languageConventional = map[string]string{
			"en-US": enUsConventionalReadme,
			"zh-CN": zhCNConventionalReadme,
		}
	}
	return languageConventional
}

type TemplateFunc func(config template_file.ConventionalConfig, dir string) error

var languageConventionalTemplate map[string]TemplateFunc

func LanguageConventionalTemplate() map[string]TemplateFunc {
	if languageConventionalTemplate == nil {
		languageConventionalTemplate = map[string]TemplateFunc{
			"en-US": templateGithubDotWalkEnUs,
			"zh-CN": templateGithubDotWalkZhCn,
		}
	}
	return languageConventionalTemplate
}

var (
	//go:embed resource/template_file/github_template/en-US/conventional_readme.md
	enUsConventionalReadme string

	//go:embed resource/template_file/github_template/en-US/pull_request_template.md
	enUSPullRequestTemplate string

	//go:embed resource/template_file/github_template/en-US/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md
	enUSCodeOfConduct string

	//go:embed resource/template_file/github_template/en-US/CONTRIBUTING_DOC/CONTRIBUTING.md
	enUSContributing string

	//go:embed resource/template_file/github_template/en-US/ISSUE_TEMPLATE/bug_report.md
	enUSBugReport string

	//go:embed resource/template_file/github_template/en-US/ISSUE_TEMPLATE/feature_request.md
	enUSFeatureRequest string

	//go:embed resource/template_file/github_template/en-US/ISSUE_TEMPLATE/help_wanted.md
	enUSHelpWanted string

	//go:embed resource/template_file/github_template/en-US/ISSUE_TEMPLATE/question.md
	enUSQuestion string
)

func templateGithubDotWalkEnUs(config template_file.ConventionalConfig, targetRootDir string) error {
	walkList := []template_file.ConventionalFile{
		{
			Name:         "pull_request_template.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "pull_request_template.md"},
			Content:      enUSPullRequestTemplate,
		},
		{
			Name:         "CODE_OF_CONDUCT.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "CODE_OF_CONDUCT.md"},
			Content:      enUSCodeOfConduct,
		},
		{
			Name:         "CONTRIBUTING.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "CONTRIBUTING.md"},
			Content:      enUSContributing,
		},
		{
			Name:         "bug_report.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "bug_report.md"},
			Content:      enUSBugReport,
		},
		{
			Name:         "feature_request.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "feature_request.md"},
			Content:      enUSFeatureRequest,
		},
		{
			Name:         "help_wanted.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "help_wanted.md"},
			Content:      enUSHelpWanted,
		},
		{
			Name:         "question.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "ISSUE_TEMPLATE", "question.md"},
			Content:      enUSQuestion,
		},
	}
	for _, cFile := range walkList {
		tPath := filepath.Join(cFile.FullPaths...)
		err := filepath_plus.CheckOrCreateFileWithStringFast(tPath, cFile.Content)
		if err != nil {
			return err
		}
		slog.Debugf("TemplateGithubDotWalkEnUs check or create file: %s", tPath)
	}
	return nil
}

var (
	//go:embed resource/template_file/github_template/zh-CN/conventional_readme.md
	zhCNConventionalReadme string

	//go:embed resource/template_file/github_template/zh-CN/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md
	zhCNCodeOfConduct string

	//go:embed resource/template_file/github_template/zh-CN/CONTRIBUTING_DOC/CONTRIBUTING.md
	zhCNContributing string
)

func templateGithubDotWalkZhCn(config template_file.ConventionalConfig, targetRootDir string) error {
	walkList := []template_file.ConventionalFile{
		{
			Name:         "CODE_OF_CONDUCT.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "zh-CN", "CODE_OF_CONDUCT.md"},
			Content:      zhCNCodeOfConduct,
		},
		{
			Name:         "CONTRIBUTING.md",
			GitOwnerName: config.GitOwnerName,
			GitRepoName:  config.GitOwnerName,
			FullPaths:    []string{targetRootDir, "CONTRIBUTING_DOC", "zh-CN", "CONTRIBUTING.md"},
			Content:      zhCNContributing,
		},
	}
	for _, cFile := range walkList {
		tPath := filepath.Join(cFile.FullPaths...)
		err := filepath_plus.CheckOrCreateFileWithStringFast(tPath, cFile.Content)
		if err != nil {
			return err
		}
		slog.Debugf("TemplateGithubDotWalkEnUs check or create file: %s", tPath)
	}
	return nil
}
