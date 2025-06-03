//go:build !test

package main

import (
	"fmt"
	os "os"

	"github.com/gookit/color"
	ghconventionalkit "github.com/sinlov/gh-conventional-kit"
	"github.com/sinlov/gh-conventional-kit/cmd/cli"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/pkg_kit"
	"github.com/sinlov/gh-conventional-kit/internal/cli_kit/urfave_cli/cli_exit_urfave"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
)

const (
	// exitCodeCmdArgs SIGINT as 1.
	exitCodeCmdArgs = 1
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build.
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	pkg_kit.InitPkgJsonContent(ghconventionalkit.PackageJson)
	embed_source.RegisterSettings(embed_source.DefaultFunctions)
	cli_exit_urfave.ChangeDefaultExitCode(exitCodeCmdArgs)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		pkg_kit.GetPackageJsonDescription(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)
	pkg_kit.SaveBuildInfo(&bdInfo)

	app := cli.NewCliApp(bdInfo)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(2)
	}

	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
