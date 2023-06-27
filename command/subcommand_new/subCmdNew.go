package subcommand_new

import (
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
)

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   "new",
			Usage:  "new command",
			Action: action,
			Flags:  flag(),
		},
	}
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template",
			Value: false,
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Set the resulting package name, defaults to the directory name",
		},
	}
}

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ new ] start")
	newCommandEntry = withGlobalFlag(c)
	return newCommandEntry.Exec()
}

func withGlobalFlag(c *cli.Context) *NewCommand {
	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &NewCommand{
		isDebug: globalEntry.Verbose,
	}
}

var newCommandEntry *NewCommand

type NewCommand struct {
	isDebug bool
}

func (n *NewCommand) Exec() error {
	return nil
}
