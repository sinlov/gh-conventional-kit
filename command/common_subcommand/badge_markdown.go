package common_subcommand

import (
	"fmt"
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
		sb.WriteString("\n")
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
		sb.WriteString("\n")
		if badgeConfig.RustVersion != "" {
			staticBadgeOrange := shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion)
			sb.WriteString(fmt.Sprintf("[![rust version](%s)](https://github.com/%s/%s)\n", staticBadgeOrange, userName, repoName))
			sb.WriteString("\n")
		}

		cratesName := badgeConfig.RustCratesName
		if cratesName == "" {
			cratesName = repoName
		}

		sb.WriteString(rust_badges.DocsRsMarkdown(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesVersionMarkdown(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesDownloadLatestMarkdown(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesLicenseMarkdown(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.DepsRsCrateLatestMarkdown(cratesName))
		sb.WriteString("\n")
	}

	if badgeConfig.NodeBadges || badgeConfig.NpmPackage != "" {
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
		sb.WriteString("\n")
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
		sb.WriteString("\n")
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
		sb.WriteString("\n")

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

		sb.WriteString("\n")

		if badgeConfig.RustVersion != "" {
			staticBadgeOrange := shields_badges.StaticBadgeOrange("rust", badgeConfig.RustVersion)
			sb.WriteString(staticBadgeOrange)
			sb.WriteString("\n")
		}

		cratesName := badgeConfig.RustCratesName
		if cratesName == "" {
			cratesName = repoName
		}

		sb.WriteString(rust_badges.DocsRs(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesVersion(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesDownloadLatest(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.CratesLicense(cratesName))
		sb.WriteString("\n")
		sb.WriteString(rust_badges.DepsRsCrateLatest(cratesName))
		sb.WriteString("\n")
	}

	if badgeConfig.NodeBadges || badgeConfig.NpmPackage != "" {
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

		sb.WriteString("\n")

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

		sb.WriteString("\n")

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
