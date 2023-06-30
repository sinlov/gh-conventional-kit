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
	_, err = git_tools.RepositoryFistRemoteInfo(projectRootPath, "")
	if err == nil {
		t.Fatal("RepositoryFistRemoteInfo with empty remote should not be nil")
	}
	assert.Equal(t, "RepositoryFistRemoteInfo remote is empty", err.Error())
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

func TestIsPathGitManagementRoot(t *testing.T) {
	t.Logf("~> mock IsPathGitManagementRoot")
	// mock IsPathGitManagementRoot

	t.Logf("~> do IsPathGitManagementRoot")
	// do IsPathGitManagementRoot
	isRoot, err := git_tools.IsPathGitManagementRoot(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isRoot)

	dirRoot, err := git_tools.IsPathGitManagementRoot(filepath.Dir(projectRootPath))
	if err == nil {
		t.Fatal("should not be nil")
	}
	t.Logf("dirRoot IsPathGitManagementRoot err: %v", err)
	assert.False(t, dirRoot)
	docRoot, err := git_tools.IsPathGitManagementRoot(filepath.Join(projectRootPath, "doc"))
	if err == nil {
		t.Fatal("should not be nil")
	}
	t.Logf("docRoot IsPathGitManagementRoot err: %v", err)
	assert.False(t, docRoot)

	// verify IsPathGitManagementRoot
}

func TestRepositoryHeadByPath(t *testing.T) {
	t.Logf("~> mock RepositoryHeadByPath")
	// mock RepositoryHeadByPath
	headByPath, err := git_tools.RepositoryHeadByPath(projectRootPath)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("headByPath: %v", headByPath)

	_, err = git_tools.RepositoryHeadByPath(filepath.Dir(projectRootPath))
	if err == nil {
		t.Fatal("should not be nil")
	}
}

func TestRepositoryNowBranchByPath(t *testing.T) {
	t.Logf("~> mock RepositoryNowBranchByPath")
	// mock RepositoryNowBranchByPath
	branchByPath, err := git_tools.RepositoryNowBranchByPath(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("branchByPath: %s", branchByPath)
	_, err = git_tools.RepositoryNowBranchByPath(filepath.Dir(projectRootPath))
	if err == nil {
		t.Fatal("should not be nil")
	}
}
