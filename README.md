[![ci](https://github.com/sinlov/gh-conventional-kit/workflows/ci/badge.svg?branch=main)](https://github.com/sinlov/gh-conventional-kit/actions/workflows/ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov/gh-conventional-kit?label=go.mod)](https://github.com/sinlov/gh-conventional-kit)
[![GoDoc](https://godoc.org/github.com/sinlov/gh-conventional-kit?status.png)](https://godoc.org/github.com/sinlov/gh-conventional-kit/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/gh-conventional-kit)](https://goreportcard.com/report/github.com/sinlov/gh-conventional-kit)
[![codecov](https://codecov.io/gh/sinlov/gh-conventional-kit/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/gh-conventional-kit)
[![docker version semver](https://img.shields.io/docker/v/sinlov/gh-conventional-kit?sort=semver)](https://hub.docker.com/r/sinlov/gh-conventional-kit/tags?page=1&ordering=last_updated)
[![docker image size](https://img.shields.io/docker/image-size/sinlov/gh-conventional-kit)](https://hub.docker.com/r/sinlov/gh-conventional-kit)
[![docker pulls](https://img.shields.io/docker/pulls/sinlov/gh-conventional-kit)](https://hub.docker.com/r/sinlov/gh-conventional-kit/tags?page=1&ordering=last_updated)
[![github release](https://img.shields.io/github/v/release/sinlov/gh-conventional-kit?style=social)](https://github.com/sinlov/gh-conventional-kit/releases)

## for what

- this cli generate conventional markdown and add `.github` template to your project

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/gh-conventional-kit)](https://github.com/sinlov/gh-conventional-kit/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [x] `badge` add badge at github project
    - `--common-badges` add common badges for github project
    - `--golang` add golang badges for github project
    - `--rust` add rust badges for github project
    - `--node` add node badges for github project
    - `-npm` add npm badges for github project
    - `--docker-user` `--docker-repo` add docker badges for github project
- [X] `markdown` subcommand generate markdown badge by program language or framework
- [X] `template` add conventional template at .github and try add badge at README.md
    - conventional contributing support `--language`
        - `en-US`
        - `zh-CN`
- [ ] more perfect test case coverage

## usage

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/sinlov/gh-conventional-kit/cmd/gh-conventional-kit@latest
# install version v1.5.0
$ go install -v github.com/sinlov/gh-conventional-kit/cmd/gh-conventional-kit@v1.5.0

# usa as docker cli tools
$ docker run --rm sinlov/gh-conventional-kit:latest -h
# use as docker cli tools with version
$ docker run --rm sinlov/gh-conventional-kit:1.5.0 -h
```

- please install [git](https://git-scm.com/) before use this cli

```bash
## show commands and global options
$ gh-conventional-kit --help

## markdown tools help
$ gh-conventional-kit markdown -h
# show common badges
$ gh-conventional-kit markdown -u [user] -r [repo]
# this project is golang
$ gh-conventional-kit markdown -u [user] -r [repo] --golang
# this project is rust
$ gh-conventional-kit markdown -u [user] -r [repo] --rust
#  crates name not same as repo name
$ gh-conventional-kit markdown -u [user] -r [repo] --rust --rust-crates some-rs
# this project is node
$ gh-conventional-kit markdown -u [user] -r [repo] --node
# multiple programming languages
$ gh-conventional-kit markdown -u [user] -r [repo] --golang --node

## badge tools help
$ gh-conventional-kit badge --help

# try at your git project root path golang
$ gh-conventional-kit --dry-run badge --golang
# append at README.md head
$ gh-conventional-kit badge --golang

# if use other language like rust
$ gh-conventional-kit --dry-run badge --rust
# multiple programming languages
$ gh-conventional-kit --dry-run badge --rust --npm
# docker badges
$ gh-conventional-kit --dry-run badge --rust --docker-user [user] --docker-repo [repo]

## template tools help
$ gh-conventional-kit template --help
# try add template at .github and try add badge at README.md
$ gh-conventional-kit --dry-run template --language en-US,zh-CN
# this project is golang
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --golang
# this project is rust
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --rust
#  crates name not same as repo name
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --rust --rust-crates some-rs
# this project is node
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --node
# multiple programming languages
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --golang --node
# do template and add badge at README.md
$ gh-conventional-kit template --language en-US,zh-CN --golang
# coverage template .github folder files
$ gh-conventional-kit template --language en-US,zh-CN --golang --coverage-folder-file
```

## evn

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                       | version             |
|:------------------------------------------|:--------------------|
| https://github.com/stretchr/testify       | v1.8.4              |
| https://github.com/sebdah/goldie          | v2.5.3              |
| https://github.com/urfave/cli/            | v2.23.7             |
| https://github.com/gookit/color           | v1.5.3              |
| https://github.com/bar-counter/slog       | v1.4.0              |
| https://github.com/sinlov-go/go-git-tools | v1.0.0              |
| https://github.com/sinlov-go/badges       | v1.3.0              |
| https://github.com/aymerick/raymond       | v2.0.2+incompatible |

# dev

## depends

in go mod project

```bash
# warning use private git_tools host must set
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
├── build                          # build output
├── cmd                            # command line main package
│         ├── main.go                 # command line entry
│         └── main_test.go            # integrated test entry
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
├── package.json                   # command line profile information
├── resource.go                    # embed resource 
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
