package shields_helper

import "fmt"

// GithubTagLatestMarkdown
//
//	return markdown for github tag latest
//
// See https://shields.io/badges/git-hub-tag-latest-sem-ver-pre-release
func GithubTagLatestMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub latest SemVer tag)](%s/github/v/tag/%s/%s)](https://github.com/%s/%s/tags)",
		ShieldsUrl, user, repo, user, repo)
}

func GitHubReleaseMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub release)](%s/github/v/release/%s/%s)](https://github.com/%s/%s/releases)",
		ShieldsUrl, user, repo, user, repo)
}
