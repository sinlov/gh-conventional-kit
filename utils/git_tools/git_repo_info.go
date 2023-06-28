package git_tools

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	gitConfig "github.com/go-git/go-git/v5/config"
	gitUrls "github.com/whilp/git-urls"
	"net/url"
	"strings"
)

func IsPathUnderGitManagement(path string) bool {
	_, err := git.PlainOpen(path)
	return err == nil
}

// RepositoryConfigPath
//
//	cfg, err := git_tools.RepositoryConfigPath(projectRootPath)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Logf("config: %v", cfg.Remotes["origin"].URLs[0])
//
// more use See https://github.com/go-git/go-git/blob/v5.7.0/config/config_test.go#L18
func RepositoryConfigPath(path string) (*gitConfig.Config, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("RepositoryConfigPath can not open repository: %s", err)
	}

	return repository.Config()

}

type GitRemoteInfo struct {
	Scheme   string
	Host     string
	Hostname string
	UserInfo *url.Userinfo
	User     string
	Repo     string
}

// RepositoryFistRemoteInfo
//
//	url, err := git_tools.RepositoryFistRemoteInfo(projectRootPath, "origin")
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Logf("url.Host: %s", url.Host)
//	t.Logf("url.User: %s", url.User)
//	t.Logf("url.Repo: %s", url.Repo)
//
// See: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html
func RepositoryFistRemoteInfo(path string, remote string) (*GitRemoteInfo, error) {
	cfg, err := RepositoryConfigPath(path)
	if err != nil {
		return nil, err
	}
	remoteConfig := cfg.Remotes[remote]
	if remoteConfig == nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s not found", remote)
	}
	if len(remoteConfig.URLs) == 0 {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs is empty", remote)
	}
	urlStr := remoteConfig.URLs[0]
	parse, err := gitUrls.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse err: %s", remote, urlStr, err)
	}
	if parse.Path == "" {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse path is empty", remote, urlStr)
	}
	removeGitPath := strings.Replace(parse.Path, ".git", "", -1)
	pathSplit := strings.Split(removeGitPath, "/")
	if len(pathSplit) < 2 {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse path not support", remote, parse.Path)
	}
	return &GitRemoteInfo{
		Scheme:   parse.Scheme,
		Host:     parse.Host,
		Hostname: parse.Hostname(),
		UserInfo: parse.User,
		User:     pathSplit[0],
		Repo:     pathSplit[1],
	}, nil
}
