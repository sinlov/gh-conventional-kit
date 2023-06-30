package subcommand_badge

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"os"
)

const commandName = "badge"

var commandEntry *BadgeCommand

type BadgeCommand struct {
	isDebug     bool
	TargetFile  string
	GitRootPath string
}

func (n *BadgeCommand) Exec() error {

	if n.TargetFile == "" {
		return fmt.Errorf("target file is empty")
	}
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "target",
			Usage: "set add badge target, defaults is: README.md",
			Value: "README.md",
		},
		&cli.StringFlag{
			Name:  "gitRootFolder",
			Usage: "set add badge target git root folder, defaults is git_tools root path, value ''",
			Value: "",
		},
	}
}

func withEntry(c *cli.Context) (*BadgeCommand, error) {
	globalEntry := command.CmdGlobalEntry()

	gitRootFolder := c.String("gitRootFolder")
	if gitRootFolder == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("can not get target foler err: %v", err)
		}
		gitRootFolder = dir
	}

	targetPath := c.String("target")
	if !filepath_plus.PathExistsFast(targetPath) {
		return nil, fmt.Errorf("target file not exists: %s", targetPath)
	}
	if filepath_plus.PathIsDir(targetPath) {
		return nil, fmt.Errorf("target file is dir: %s", targetPath)
	}
	return &BadgeCommand{
		isDebug:     globalEntry.Verbose,
		TargetFile:  targetPath,
		GitRootPath: gitRootFolder,
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "add badge at github project",
			Action: action,
			Flags:  flag(),
		},
	}
}
