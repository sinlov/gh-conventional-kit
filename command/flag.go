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
			Name:    "verbose",
			Usage:   "open cli verbose mode",
			Value:   false,
			EnvVars: []string{constant.EnvKeyCliVerbose},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second",
			Usage:   "command timeout setting second. default 10",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{constant.EnvKeyCliTimeoutSecond},
		},
		&cli.StringFlag{
			Name:    "config.log_level",
			Usage:   fmt.Sprintf("command clog level. default %s", slog.INFO),
			Value:   slog.INFO,
			Hidden:  true,
			EnvVars: []string{constant.EnvLogLevel},
		},
	}
}
