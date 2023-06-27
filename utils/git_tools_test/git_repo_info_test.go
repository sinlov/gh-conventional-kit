package git_tools_test

import (
	"github.com/sinlov/gh-conventional-kit/utils/git_tools"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

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
