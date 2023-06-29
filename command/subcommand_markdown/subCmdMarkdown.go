package subcommand_markdown

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"

	"github.com/sinlov-go/badges/codecov_badges"
	"github.com/sinlov-go/badges/golang_badges"
	"github.com/sinlov-go/badges/npm_badges"
	"github.com/sinlov-go/badges/rust_badges"
	"github.com/sinlov-go/badges/shields_badges"
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

	if n.BadgeConfig.GolangBadges {
		if n.UserName == "" || n.RepoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		fmt.Println("\nGolang badges:")
		fmt.Println(golang_badges.GithubGoModVersionMarkdown(n.UserName, n.RepoName))
		fmt.Println(golang_badges.GithubGoDocMarkdown(n.UserName, n.RepoName))
		fmt.Println(golang_badges.GithubGoReportCardMarkdown(n.UserName, n.RepoName))
	}

	if n.BadgeConfig.RustBadges {
		if n.UserName == "" || n.RepoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		fmt.Println("\nRust badges:")
		if n.BadgeConfig.RustVersion != "" {
			fmt.Println(fmt.Sprintf("[![rust version](%s)](https://github.com/%s/%s)",
				shields_badges.StaticBadgeOrange("rust", n.BadgeConfig.RustVersion),
				n.UserName, n.RepoName,
			))
		}
		fmt.Println(rust_badges.DocsRsMarkdown(n.RepoName))
		fmt.Println(rust_badges.CratesVersionMarkdown(n.RepoName))
		fmt.Println(rust_badges.CratesDownloadLatestMarkdown(n.RepoName))
		fmt.Println(rust_badges.CratesLicenseMarkdown(n.RepoName))
		fmt.Println(rust_badges.DepsRsCrateLatestMarkdown(n.RepoName))
	}

	if n.BadgeConfig.NpmPackage != "" {
		fmt.Println("\nnpm badges:")
		fmt.Println(npm_badges.VersionLatestMarkdown(n.BadgeConfig.NpmPackage))
		fmt.Println(npm_badges.NodeLtsVersionMarkdown(n.BadgeConfig.NpmPackage))
		fmt.Println(npm_badges.LicenseMarkdown(n.BadgeConfig.NpmPackage))
		fmt.Println(npm_badges.DownloadLatestMonthMarkdown(n.BadgeConfig.NpmPackage))
		fmt.Println(npm_badges.CollaboratorsMarkdown(n.BadgeConfig.NpmPackage))
	}

	if n.BadgeConfig.DockerUser != "" {
		if n.BadgeConfig.DockerRepo == "" {
			return fmt.Errorf("need set --docker-repo")
		}
		fmt.Println("\nDocker badges:")
		fmt.Println(shields_badges.DockerHubImageVersionSemverMarkdown(n.BadgeConfig.DockerUser, n.BadgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImageSizeMarkdown(n.BadgeConfig.DockerUser, n.BadgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImagePullMarkdown(n.BadgeConfig.DockerUser, n.BadgeConfig.DockerRepo))
	}

	if !n.BadgeConfig.NoCommonBadges {
		if n.UserName == "" || n.RepoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}

		fmt.Println("\nCommon badges:")
		fmt.Println(shields_badges.GithubLicenseMarkdown(n.UserName, n.RepoName))
		fmt.Println(codecov_badges.GithubMarkdown(n.UserName, n.RepoName, n.Branch))
		fmt.Println(shields_badges.GithubLatestSemVerTagMarkdown(n.UserName, n.RepoName))
		fmt.Println(shields_badges.GithubReleaseMarkdown(n.UserName, n.RepoName))
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
