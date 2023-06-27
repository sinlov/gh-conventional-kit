package subcommand_template

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
)

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   "template",
			Usage:  "add conventional template at .github",
			Action: action,
			Flags:  flag(),
		},
	}
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "rootFolder",
			Usage: "set add badge target root folder, defaults is git_tools root path, value ''",
			Value: "",
		},
	}
}

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ badge ] start")
	entry, err := withEntry(c)
	newCommandEntry = entry
	if err != nil {
		return err
	}
	return newCommandEntry.Exec()
}

func withEntry(c *cli.Context) (*TemplateCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	targetRootFolder := c.String("rootFolder")
	if targetRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}
		targetRootFolder = dir
	}
	return &TemplateCommand{
		isDebug:          globalEntry.Verbose,
		TargetFolderPath: targetRootFolder,
	}, nil
}

var newCommandEntry *TemplateCommand

type TemplateCommand struct {
	isDebug          bool
	TargetFolderPath string
}

func (n *TemplateCommand) Exec() error {

	return nil
}
