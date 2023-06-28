package git_tools_test

import (
	"github.com/sinlov/gh-conventional-kit/utils/git_tools"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestRepositoryFistRemote(t *testing.T) {
	t.Logf("~> mock RepositoryFistRemoteInfo")
	// mock RepositoryFistRemoteInfo
	url, err := git_tools.RepositoryFistRemoteInfo(projectRootPath, "origin")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("url.Host: %s", url.Host)
	t.Logf("url.User: %s", url.User)
	t.Logf("url.Repo: %s", url.Repo)

	t.Logf("~> do RepositoryFistRemoteInfo")
	// do RepositoryFistRemoteInfo

	// verify RepositoryFistRemoteInfo
}

func TestRepositoryConfigPath(t *testing.T) {
	t.Logf("~> mock RepositoryConfigPath")
	// mock RepositoryConfigPath
	cfg, err := git_tools.RepositoryConfigPath(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config: %v", cfg.Remotes["origin"].URLs[0])

	t.Logf("~> do RepositoryConfigPath")
	// do RepositoryConfigPath

	// verify RepositoryConfigPath
}

func TestIsPathUnderGitManagement(t *testing.T) {
	t.Logf("~> mock IsPathUnderGitManagement")
	// mock IsPathUnderGitManagement
	t.Logf("projectRoot: %s", projectRootPath)

	t.Logf("~> do IsPathUnderGitManagement")
	// do IsPathUnderGitManagement
	assert.True(t, git_tools.IsPathUnderGitManagement(projectRootPath))

	// verify IsPathUnderGitManagement

	topDir := filepath.Dir(projectRootPath)
	assert.False(t, git_tools.IsPathUnderGitManagement(topDir))
}
