package subcommand_badge

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/sinlov/gh-conventional-kit/utils/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/utils/urfave_cli"
	"github.com/urfave/cli/v2"
)

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   "badge",
			Usage:  "add badge at github project",
			Action: action,
			Flags:  flag(),
		},
	}
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "target",
			Usage: "set add badge target, defaults is: README.md",
			Value: "README.md",
		},
		&cli.BoolFlag{
			Name:  "golang",
			Usage: "set add badge target is golang project, defaults is: false",
		},
	}
}

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ badge ] start")
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	newCommandEntry = entry
	return newCommandEntry.Exec()
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	globalEntry := command.CmdGlobalEntry()
	targetPath := c.String("target")
	if !filepath_plus.PathExistsFast(targetPath) {
		return nil, fmt.Errorf("target file not exists: %s", targetPath)
	}
	if filepath_plus.PathIsDir(targetPath) {
		return nil, fmt.Errorf("target file is dir: %s", targetPath)
	}
	return &NewCommand{
		isDebug:    globalEntry.Verbose,
		TargetFile: targetPath,
	}, nil
}

var newCommandEntry *NewCommand

type NewCommand struct {
	isDebug    bool
	TargetFile string
}

func (n *NewCommand) Exec() error {

	if n.TargetFile == "" {
		return fmt.Errorf("target file is empty")
	}
	return nil
}
