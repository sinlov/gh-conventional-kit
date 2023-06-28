package subcommand_template

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
)

const commandName = "template"

var commandEntry *TemplateCommand

type TemplateCommand struct {
	isDebug     bool
	GitRootPath string
}

func (n *TemplateCommand) Exec() error {

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "gitRootFolder",
			Usage: "set add badge target git root folder, defaults is git_tools root path, value ''",
			Value: "",
		},
	}
}

func withEntry(c *cli.Context) (*TemplateCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	gitRootFolder := c.String("gitRootFolder")
	if gitRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}
		gitRootFolder = dir
	}
	return &TemplateCommand{
		isDebug:     globalEntry.Verbose,
		GitRootPath: gitRootFolder,
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)
	entry, err := withEntry(c)
	commandEntry = entry
	if err != nil {
		return err
	}
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "add conventional template at .github",
			Action: action,
			Flags:  flag(),
		},
	}
}
