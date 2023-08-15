package command

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/log"
	"github.com/sinlov/gh-conventional-kit/internal/pkgJson"
	"github.com/sinlov/gh-conventional-kit/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
	"os"
)

type GlobalConfig struct {
	LogLevel      string
	TimeoutSecond uint
	RunRootPath   string
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name    string
		Version string
		Verbose bool
		DryRun  bool
		RootCfg GlobalConfig
	}
)

var (
	cmdGlobalEntry *GlobalCommand
)

// CmdGlobalEntry
//
//	return global command entry
func CmdGlobalEntry() *GlobalCommand {
	return cmdGlobalEntry
}

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}
	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return err
	}
	return nil
}

func (c *GlobalCommand) globalExec() error {

	return nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	err := log.InitLog(isVerbose, !isVerbose)
	if err != nil {
		panic(err)
	}
	cliVersion := pkgJson.GetPackageJsonVersionGoStyle(false)
	if isVerbose {
		slog.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	appName := pkgJson.GetPackageJsonName()
	cmdGlobal, errFlag := withGlobalFlag(c, cliVersion, appName)
	if errFlag != nil {
		return cli_exit_urfave.ErrMsg(errFlag, "init global flag")
	}

	cmdGlobalEntry = cmdGlobal
	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	if isVerbose {
		slog.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
	}
	return nil
}

func withGlobalFlag(c *cli.Context, cliVersion, cliName string) (*GlobalCommand, error) {
	isVerbose := c.Bool(constant.NameCliVerbose)
	cliRunRootPath := c.String(constant.NameCliRunPath)
	if len(cliRunRootPath) == 0 {
		rootDir, err := os.Getwd()
		if err != nil {
			slog.Errorf(err, "get rooted path name corresponding to the current directory path err")
			return nil, cli_exit_urfave.Err(err)
		}
		cliRunRootPath = rootDir
	}
	err := gh_conventional_kit.CheckAllResource(cliRunRootPath)
	if err != nil {
		slog.Errorf(err, "check all resource err")
		return nil, cli_exit_urfave.Err(err)
	}

	config := GlobalConfig{
		LogLevel:      c.String(constant.NameLogLevel),
		TimeoutSecond: c.Uint(constant.NameCliTimeoutSecond),
		RunRootPath:   cliRunRootPath,
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		DryRun:  c.Bool(constant.NameCliDryRun),
		RootCfg: config,
	}
	return &p, nil
}
