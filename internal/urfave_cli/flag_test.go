package urfave_cli

import (
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestGlobalFlag(t *testing.T) {
	t.Logf("~> mock GlobalFlag")
	// mock GlobalFlag

	t.Logf("~> do GlobalFlag")
	// do GlobalFlag
	flags := UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	// verify GlobalFlag
	assert.NotEqual(t, 0, len(flags))
}

func TestUrfaveCliAppendCliFlag(t *testing.T) {
	t.Logf("~> mock UrfaveCliAppendCliFlag")
	// mock UrfaveCliAppendCliFlag

	t.Logf("~> do UrfaveCliAppendCliFlag")
	// do UrfaveCliAppendCliFlag
	flags := UrfaveCliAppendCliFlag([]cli.Flag{}, []cli.Flag{})
	assert.Equal(t, 0, len(flags))

	// verify UrfaveCliAppendCliFlag
	flags = UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	assert.NotEqual(t, 0, len(flags))
}
