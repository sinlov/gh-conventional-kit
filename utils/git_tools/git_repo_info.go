package git_tools

import (
	"github.com/go-git/go-git/v5"
	_ "github.com/go-git/go-git/v5"
)

func IsPathUnderGitManagement(path string) bool {
	_, err := git.PlainOpen(path)
	return err == nil
}

func RepositoryWithPath(path string) (*git.Repository, error) {
	return git.PlainOpen(path)
}
