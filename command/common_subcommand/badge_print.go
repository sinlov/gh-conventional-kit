package common_subcommand

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov-go/badges/codecov_badges"
	"github.com/sinlov-go/badges/golang_badges"
	"github.com/sinlov-go/badges/node_badges"
	"github.com/sinlov-go/badges/npm_badges"
	"github.com/sinlov-go/badges/rust_badges"
	"github.com/sinlov-go/badges/shields_badges"
	"github.com/sinlov/gh-conventional-kit/constant"
	"strings"
)

func BadgeByConfigWithMarkdown(
	badgeConfig *constant.BadgeConfig,
	userName, repoName string,
	branch string,
) (string, error) {
	var sb strings.Builder
	sb.Grow(0)

	if badgeConfig.GolangBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(golang_badges.GithubGoModVersionMarkdown(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(golang_badges.GithubGoDocMarkdown(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(golang_badges.GithubGoReportCardMarkdown(userName, repoName))
		sb.WriteString("\n")
	}

	if badgeConfig.RustBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		if badgeConfig.RustVersion != "" {
			staticBadgeOrange := shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion)
			sb.WriteString(fmt.Sprintf("[![rust version](%s)](https://github.com/%s/%s)\n", staticBadgeOrange, userName, repoName))
			sb.WriteString("\n")
		}
		sb.WriteString(rust_badges.DocsRsMarkdown(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesVersionMarkdown(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesDownloadLatestMarkdown(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesLicenseMarkdown(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.DepsRsCrateLatestMarkdown(repoName))
		sb.WriteString("\n")
	}

	if badgeConfig.NodeBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(node_badges.GitHubPackageJsonVersionMarkdown(userName, repoName))
		sb.WriteString("\n")
	}
	if badgeConfig.NpmPackage != "" {
		sb.WriteString(npm_badges.VersionLatestMarkdown(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.NodeLtsVersionMarkdown(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.LicenseMarkdown(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.DownloadLatestMonthMarkdown(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.CollaboratorsMarkdown(badgeConfig.NpmPackage))
		sb.WriteString("\n")
	}
	if badgeConfig.DockerUser != "" {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(shields_badges.DockerHubImageVersionSemverMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.DockerHubImageSizeMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.DockerHubImagePullMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
	}

	if !badgeConfig.NoCommonBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(shields_badges.GithubLicenseMarkdown(userName, repoName))
		sb.WriteString("\n")

		if branch != "" {
			sb.WriteString(codecov_badges.GithubMarkdown(userName, repoName, branch))
			sb.WriteString("\n")
		}
		sb.WriteString(shields_badges.GithubLatestSemVerTagMarkdown(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.GithubReleaseMarkdown(userName, repoName))
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func BadgeByConfig(
	badgeConfig *constant.BadgeConfig,
	userName, repoName string,
	branch string,
) (string, error) {
	var sb strings.Builder
	sb.Grow(0)

	if badgeConfig.GolangBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(golang_badges.GithubGoModVersion(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(golang_badges.GithubGoDoc(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(golang_badges.GithubGoReportCard(userName, repoName))
		sb.WriteString("\n")
	}

	if badgeConfig.RustBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		if badgeConfig.RustVersion != "" {
			staticBadgeOrange := shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion)
			sb.WriteString(staticBadgeOrange)
			sb.WriteString("\n")
		}
		sb.WriteString(rust_badges.DocsRs(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesVersion(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesDownloadLatest(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesLicense(repoName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.DepsRsCrateLatest(repoName))
		sb.WriteString("\n")
	}

	if badgeConfig.NodeBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(node_badges.GitHubPackageJsonVersion(userName, repoName))
		sb.WriteString("\n")
	}
	if badgeConfig.NpmPackage != "" {
		sb.WriteString(npm_badges.VersionLatest(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.NodeLtsVersion(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.License(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.DownloadLatestMonth(badgeConfig.NpmPackage))
		sb.WriteString("\n")
		sb.WriteString(npm_badges.Collaborators(badgeConfig.NpmPackage))
		sb.WriteString("\n")
	}
	if badgeConfig.DockerUser != "" {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(shields_badges.DockerHubImageVersionSemver(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.DockerHubImageSize(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.DockerHubImagePull(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		sb.WriteString("\n")
	}

	if !badgeConfig.NoCommonBadges {
		if userName == "" || repoName == "" {
			return sb.String(), fmt.Errorf("need set --user and --repo")
		}
		sb.WriteString(shields_badges.GithubLicense(userName, repoName))
		sb.WriteString("\n")

		if branch != "" {
			sb.WriteString(codecov_badges.Github(userName, repoName, branch))
			sb.WriteString("\n")
		}
		sb.WriteString(shields_badges.GithubLatestSemVerTag(userName, repoName))
		sb.WriteString("\n")
		sb.WriteString(shields_badges.GithubRelease(userName, repoName))
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func PrintBadgeByConfigWithMarkdown(
	badgeConfig *constant.BadgeConfig,
	userName, repoName string,
	branch string,
) error {
	if badgeConfig.GolangBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nGolang badges:\n")
		fmt.Println(golang_badges.GithubGoModVersionMarkdown(userName, repoName))
		fmt.Println(golang_badges.GithubGoDocMarkdown(userName, repoName))
		fmt.Println(golang_badges.GithubGoReportCardMarkdown(userName, repoName))
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
		fmt.Println(rust_badges.DocsRsMarkdown(repoName))
		fmt.Println(rust_badges.CratesVersionMarkdown(repoName))
		fmt.Println(rust_badges.CratesDownloadLatestMarkdown(repoName))
		fmt.Println(rust_badges.CratesLicenseMarkdown(repoName))
		fmt.Println(rust_badges.DepsRsCrateLatestMarkdown(repoName))
	}

	if badgeConfig.NodeBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}
		color.Greenf("\nnode badges:\n")
		fmt.Println(node_badges.GitHubPackageJsonVersionMarkdown(userName, repoName))
	}

	if badgeConfig.NpmPackage != "" {
		color.Greenf("\nnpm badges:\n")
		fmt.Println(npm_badges.VersionLatestMarkdown(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.NodeLtsVersionMarkdown(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.LicenseMarkdown(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.DownloadLatestMonthMarkdown(badgeConfig.NpmPackage))
		fmt.Println(npm_badges.CollaboratorsMarkdown(badgeConfig.NpmPackage))
	}

	if badgeConfig.DockerUser != "" {
		if badgeConfig.DockerRepo == "" {
			return fmt.Errorf("need set --docker-repo")
		}
		color.Greenf("\nDocker badges:\n")
		fmt.Println(shields_badges.DockerHubImageVersionSemverMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImageSizeMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
		fmt.Println(shields_badges.DockerHubImagePullMarkdown(badgeConfig.DockerUser, badgeConfig.DockerRepo))
	}

	if !badgeConfig.NoCommonBadges {
		if userName == "" || repoName == "" {
			return fmt.Errorf("need set --user and --repo")
		}

		color.Greenf("\nCommon badges:\n")
		fmt.Println(shields_badges.GithubLicenseMarkdown(userName, repoName))
		if branch != "" {
			fmt.Println(codecov_badges.GithubMarkdown(userName, repoName, branch))
		}
		fmt.Println(shields_badges.GithubLatestSemVerTagMarkdown(userName, repoName))
		fmt.Println(shields_badges.GithubReleaseMarkdown(userName, repoName))
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
		fmt.Println(rust_badges.DocsRs(repoName))
		fmt.Println(rust_badges.CratesVersion(repoName))
		fmt.Println(rust_badges.CratesDownloadLatest(repoName))
		fmt.Println(rust_badges.CratesLicense(repoName))
		fmt.Println(rust_badges.DepsRsCrateLatest(repoName))
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
			fmt.Println(codecov_badges.Github(userName, repoName, branch))
		}
		fmt.Println(shields_badges.GithubLatestSemVerTag(userName, repoName))
		fmt.Println(shields_badges.GithubRelease(userName, repoName))
	}
	return nil
}
