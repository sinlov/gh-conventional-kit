package subcommand_new

import (
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli"
	"github.com/urfave/cli/v2"
)

const commandName = "new"

var commandEntry *NewCommand

type NewCommand struct {
	isDebug bool
}

func (n *NewCommand) Exec() error {

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template_file",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Set the resulting package name, defaults to the directory name",
		},
	}
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &NewCommand{
		isDebug: globalEntry.Verbose,
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
			Usage:  "",
			Action: action,
			Flags:  flag(),
		},
	}
}
