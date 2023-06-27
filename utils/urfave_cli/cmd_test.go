package urfave_cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestUrfaveCliAppendCliCommand(t *testing.T) {
	t.Logf("~> mock UrfaveCliAppendCliCommand")
	// mock UrfaveCliAppendCliCommand

	t.Logf("~> do UrfaveCliAppendCliCommand")
	// do UrfaveCliAppendCliCommand
	var appCommands []*cli.Command
	appCommands = UrfaveCliAppendCliCommand(appCommands, []*cli.Command{})
	assert.Equal(t, 0, len(appCommands))

	appCommands = UrfaveCliAppendCliCommand(appCommands, []*cli.Command{
		{
			Name:  "new",
			Usage: "new command",
		},
	})

	// verify UrfaveCliAppendCliCommand
	assert.Equal(t, 1, len(appCommands))
}
