package command

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/urfave/cli/v2"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    constant.NameCliVerbose,
			Usage:   "open cli verbose mode",
			Value:   false,
			EnvVars: []string{constant.EnvKeyCliVerbose},
		},
		&cli.BoolFlag{
			Name:  constant.NameCliDryRun,
			Usage: "See the commands that running standard-version would run",
			Value: false,
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    constant.NameCliTimeoutSecond,
			Usage:   "command timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{constant.EnvKeyCliTimeoutSecond},
		},
		&cli.StringFlag{
			Name:    constant.NameLogLevel,
			Usage:   fmt.Sprintf("command clog level. default %s", slog.INFO),
			Value:   slog.INFO,
			Hidden:  true,
			EnvVars: []string{constant.EnvLogLevel},
		},
		&cli.StringFlag{
			Name:   constant.NameCliRunPath,
			Usage:  "command run path. default current path",
			Hidden: true,
			Value:  "",
		},
	}
}
