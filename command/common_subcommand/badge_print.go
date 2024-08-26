package common_subcommand

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov-go/badges"
	"github.com/sinlov-go/badges/codecov_badges"
	"github.com/sinlov-go/badges/golang_badges"
	"github.com/sinlov-go/badges/node_badges"
	"github.com/sinlov-go/badges/npm_badges"
	"github.com/sinlov-go/badges/rust_badges"
	"github.com/sinlov-go/badges/shields_badges"
	"github.com/sinlov/gh-conventional-kit/constant"
)

func PrintBadgeByConfigWithMarkdown(
	badgeConfig *constant.BadgeConfig,
	userName, repoName string,
	branch string,
) error {

	htmlMarkdownSize := badges.MarkdownImgSizes

	if badgeConfig.GolangBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nGolang badges:\n")

		if badgeConfig.EnableHtmlFlag {
			fmt.Println(golang_badges.GithubGoModVersionHtmlMarkdown(userName, repoName, htmlMarkdownSize))
			fmt.Println(golang_badges.GithubGoDocHtmlMarkdown(userName, repoName, htmlMarkdownSize))
			fmt.Println(golang_badges.GithubGoReportCardHtmlMarkdown(userName, repoName, htmlMarkdownSize))
		} else {
			fmt.Println(golang_badges.GithubGoModVersionMarkdown(userName, repoName))
			fmt.Println(golang_badges.GithubGoDocMarkdown(userName, repoName))
			fmt.Println(golang_badges.GithubGoReportCardMarkdown(userName, repoName))
		}
	}

	if badgeConfig.RustBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nRust badges:\n")
		if badgeConfig.RustVersion != "" {
			staticBadgeOrange := shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion)
			fmt.Printf("[![rust version](%s)](https://github.com/%s/%s)\n", staticBadgeOrange, userName, repoName)
		}
		cratesName := badgeConfig.RustCratesName
		if cratesName == "" {
			cratesName = repoName
		}

		if badgeConfig.EnableHtmlFlag {
			fmt.Println(rust_badges.DocsRsHtmlMarkdown(cratesName, htmlMarkdownSize))
			fmt.Println(rust_badges.CratesVersionHtmlMarkdown(cratesName, htmlMarkdownSize))
			fmt.Println(rust_badges.CratesDownloadLatestHtmlMarkdown(cratesName, htmlMarkdownSize))
			fmt.Println(rust_badges.CratesLicenseHtmlMarkdown(cratesName, htmlMarkdownSize))
			fmt.Println(rust_badges.DepsRsCrateLatestHtmlMarkdown(cratesName, htmlMarkdownSize))
		} else {

			fmt.Println(rust_badges.DocsRsMarkdown(cratesName))
			fmt.Println(rust_badges.CratesVersionMarkdown(cratesName))
			fmt.Println(rust_badges.CratesDownloadLatestMarkdown(cratesName))
			fmt.Println(rust_badges.CratesLicenseMarkdown(cratesName))
			fmt.Println(rust_badges.DepsRsCrateLatestMarkdown(cratesName))
		}
	}

	if badgeConfig.NodeBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nnode badges:\n")
		if badgeConfig.EnableHtmlFlag {
			fmt.Println(node_badges.GitHubPackageJsonVersionHtmlMarkdown(userName, repoName, htmlMarkdownSize))
		} else {
			fmt.Println(node_badges.GitHubPackageJsonVersionMarkdown(userName, repoName))
		}
	}

	if badgeConfig.NpmPackage != "" {
		color.Greenf("\nnpm badges:\n")
		if badgeConfig.EnableHtmlFlag {
			fmt.Println(npm_badges.VersionLatestHtmlMarkdown(badgeConfig.NpmPackage, htmlMarkdownSize))
			fmt.Println(npm_badges.NodeLtsVersionHtmlMarkdown(badgeConfig.NpmPackage, htmlMarkdownSize))
			fmt.Println(npm_badges.LicenseMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.DownloadLatestTimesMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.CollaboratorsMarkdown(badgeConfig.NpmPackage))
		} else {
			fmt.Println(npm_badges.VersionLatestMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.NodeLtsVersionMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.LicenseMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.DownloadLatestMonthMarkdown(badgeConfig.NpmPackage))
			fmt.Println(npm_badges.CollaboratorsMarkdown(badgeConfig.NpmPackage))
		}
	}

	if badgeConfig.DockerUser != "" {
		if badgeConfig.DockerRepo == "" {
			return fmt.Errorf("need set --docker-repo")
		}
		color.Greenf("\nDocker badges:\n")
		if badgeConfig.EnableHtmlFlag {
			fmt.Println(shields_badges.DockerHubImageVersionSemverHtmlMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo, htmlMarkdownSize))
			fmt.Println(shields_badges.DockerHubImageSizeHtmlMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo, htmlMarkdownSize))
			fmt.Println(shields_badges.DockerHubImagePullHtmlMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo, htmlMarkdownSize))
		} else {

			fmt.Println(shields_badges.DockerHubImageVersionSemverMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
			fmt.Println(shields_badges.DockerHubImageSizeMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
			fmt.Println(shields_badges.DockerHubImagePullMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		}
	}

	if !badgeConfig.NoCommonBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}

		color.Greenf("\nCommon badges:\n")
		if badgeConfig.EnableHtmlFlag {
			fmt.Println(shields_badges.GithubLicenseHtmlMarkdown(userName, repoName, htmlMarkdownSize))
		} else {
			fmt.Println(shields_badges.GithubLicenseMarkdown(userName, repoName))
		}
		if branch != "" {

			if badgeConfig.CodeCovBadges {
				if badgeConfig.EnableHtmlFlag {
					fmt.Println(codecov_badges.GithubHtmlMarkdown(userName, repoName, branch, htmlMarkdownSize))
				} else {
					fmt.Println(codecov_badges.GithubMarkdown(userName, repoName, branch))
				}
			}

		}
		if badgeConfig.EnableHtmlFlag {
			fmt.Println(shields_badges.GithubLatestSemVerTagHtmlMarkdown(userName, repoName, htmlMarkdownSize))
			fmt.Println(shields_badges.GithubReleaseHtmlMarkdown(userName, repoName, htmlMarkdownSize))
		} else {
			fmt.Println(shields_badges.GithubLatestSemVerTagMarkdown(userName, repoName))
			fmt.Println(shields_badges.GithubReleaseMarkdown(userName, repoName))
		}
	}
	return nil
}

func PrintBadgeByConfig(
	badgeConfig *constant.BadgeConfig,
	userName, repoName string,
	branch string,
) error {
	if badgeConfig.GolangBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nGolang badges:\n")
		fmt.Println(golang_badges.GithubGoModVersion(userName, repoName))
		fmt.Println(golang_badges.GithubGoDoc(userName, repoName))
		fmt.Println(golang_badges.GithubGoReportCard(userName, repoName))
	}

	if badgeConfig.RustBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nRust badges:\n")
		if badgeConfig.RustVersion != "" {
			fmt.Println(shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion))
		}
		cratesName := badgeConfig.RustCratesName
		if cratesName == "" {
			cratesName = repoName
		}
		fmt.Println(rust_badges.DocsRs(cratesName))
		fmt.Println(rust_badges.CratesVersion(cratesName))
		fmt.Println(rust_badges.CratesDownloadLatest(cratesName))
		fmt.Println(rust_badges.CratesLicense(cratesName))
		fmt.Println(rust_badges.DepsRsCrateLatest(cratesName))
	}

	if badgeConfig.NodeBadges {
		color.Greenf("\nnode badges:\n")
		fmt.Println(node_badges.GitHubPackageJsonVersion(userName, repoName))
	}

	if badgeConfig.NpmPackage != "" {
		color.Greenf("\nnpm badges:\n")
		fmt.Println(npm_badges.VersionLatest(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.NodeLtsVersion(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.License(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.DownloadLatestMonth(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.Collaborators(badgeConfig.NpmPackage))
	}

	if badgeConfig.DockerUser != "" {
		if badgeConfig.DockerRepo == "" {
			return fmt.Errorf("need set --docker-repo")
		}
		color.Greenf("\nDocker badges:\n")
		fmt.Println(shields_badges.DockerHubImageVersionSemver(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImageSize(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImagePull(badgeConfig.DockerUser, badgeConfig.DockerRepo))
	}

	if !badgeConfig.NoCommonBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}

		color.Greenf("\nCommon badges:\n")
		fmt.Println(shields_badges.GithubLicense(userName, repoName))
		if branch != "" {
			if badgeConfig.CodeCovBadges {
				fmt.Println(codecov_badges.Github(userName, repoName, branch))
			}
		}
		fmt.Println(shields_badges.GithubLatestSemVerTag(userName, repoName))
		fmt.Println(shields_badges.GithubRelease(userName, repoName))
	}
	return nil
}
