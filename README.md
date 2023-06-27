[![golang-ci](https://github.com/sinlov/gh-conventional-kit/workflows/golang-ci/badge.svg?branch=main)](https://github.com/sinlov/gh-conventional-kit/actions/workflows/golang-ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov/gh-conventional-kit?label=go.mod)](https://github.com/sinlov/gh-conventional-kit)
[![GoDoc](https://godoc.org/github.com/sinlov/gh-conventional-kit?status.png)](https://godoc.org/github.com/sinlov/gh-conventional-kit/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/gh-conventional-kit)](https://goreportcard.com/report/github.com/sinlov/gh-conventional-kit)
[![codecov](https://codecov.io/gh/sinlov/gh-conventional-kit/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/gh-conventional-kit)
[![docker version semver](https://img.shields.io/docker/v/sinlov/gh-conventional-kit?sort=semver)](https://hub.docker.com/r/sinlov/gh-conventional-kit/tags?page=1&ordering=last_updated)
[![docker image size](https://img.shields.io/docker/image-size/sinlov/gh-conventional-kit)](https://hub.docker.com/r/sinlov/gh-conventional-kit)
[![docker pulls](https://img.shields.io/docker/pulls/sinlov/gh-conventional-kit)](https://hub.docker.com/r/sinlov/gh-conventional-kit/tags?page=1&ordering=last_updated)
[![github release](https://img.shields.io/github/v/release/sinlov/gh-conventional-kit?style=social)](https://github.com/sinlov/gh-conventional-kit/releases)

## for what

- this project used to cli with golang

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/gh-conventional-kit)](https://github.com/sinlov/gh-conventional-kit/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below
    - `github.com/sinlov/gh-conventional-kit` to your package name
    - `sinlov` to your owner name
    - `gh-conventional-kit` to your project name

## evn

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/urfave/cli/      | v2.23.7 |
| https://github.com/bar-counter/slog | v1.4.0  |
| https://github.com/go-git/go-git    | v5.7.0  |

# dev

## depends

in go mod project

```bash
# warning use privte git_tools host must set
# global set for once
# add private git_tools host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git_tools config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git_tools config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git_tools config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git_tools ls-remote -q http://github.com/sinlov/gh-conventional-kit.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov/gh-conventional-kit
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov/gh-conventional-kit | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## local dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run and shell help
$ make devHelp

# run at CLI_VERBOSE=true
$ make dev

# run at ordinary mode
$ make run
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

### docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

### EngineeringStructure

```
.
├── Dockerfile                     # ci docker build
├── Dockerfile.s6                  # local docker build
├── Makefile                       # make entry
├── README.md
├── command                        # command line package
│         ├── TestMain.go             # common entry in unit test package
│         ├── flag.go                 # global flag
│         ├── global.go               # global command
│         ├── global_test.go          # global command unit test
│         ├── golder_data_test.go     # unit test test data case
│         ├── init_test.go            # unit test initialization tool
│         └── subcommand_new          # subcommandPackage new
├── constant                       # constant package 
│         └── env.go                  # constant environment variable
├── doc                            # command line tools documentation
│         └── cmd.md
├── go.mod
├── go.sum
├── main.go                        # command line entry
├── main_test.go                   # integrated test entry
├── package.json                   # command line profile information
├── utils                          # toolkit package
│         ├── env_kit                 # environment variables toolkit
│         ├── log                     # log toolkit
│         ├── pkgJson                 # package.json toolkit
│         └── urfave_cli              # urfave/cli toolkit
├── vendor
└── z-MakefileUtils                # make toolkit

```

### log

- cli log use [github.com/sinlov/go-logger](https://github.com/bar-counter/slog)
    - open debug log by env `CLI_VERBOSE=true` or global flag `--verbose`

```go
package foo

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ new ] start") // this not show at CLI_VERBOSE=false

	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	return nil
}
```