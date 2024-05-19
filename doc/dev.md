# dev

## evn

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version
- change `golangci-lint@v1.53.3` from [golangci-lint version release](https://github.com/golangci/golangci-lint/releases) to new version
    - more info see [golangci-lint local-installation](https://golangci-lint.run/usage/install/#local-installation)

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

## dev tasks

```bash
# It needs to be executed after the first use or update of dependencies.
make init dep
```

- test code

```bash
make test
# benchmark and coverage show
make ci.test.benchmark ci.coverage.show
```

- ci to fast check as CI pipeline

```bash
# check style at local
make style

# run ci at local
make ci
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
