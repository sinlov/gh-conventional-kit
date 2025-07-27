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

## Features

- [X] `markdown` subcommand generate markdown badge by program language or framework, this
  subcommand `only output into stdout`
    - markdown can use git url or http url to get git info (only support github now)
        - if not set `--user` or `--repo` can input git url like `https//...` or `git@...` (v1.11.+)
        - with empty `--user` or `--repo` and empty `arg0` will try to find out git info from current git project (
          v1.12.+)
    - can cover by `arg0` to fast get badge output, event not in git managed project
    - `--no-common-badges` no badges common subcommand for this repo (default: false)
    - `--golang` add golang badges for github project
    - `--rust` add rust badges for github project
        - `--rust-crates` value crates.io name badges for this repo, if not set, use repo name
        - use github reps badge for [deps.rs](https://deps.rs/) (1.17+)
    - `--node` add node badges for github project
    - `--npm` add npm badges for github project
    - `--docker-user` `--docker-repo` add docker badges for github project
- [x] `badge` add badge at github project
    - support badge same as `markdown` subcommand, different must use in `git` managed project
- [X] `template` add conventional template at .github and try add badge at README.md
    - support badge same as `markdown` subcommand, different must use in `git` managed project
    - `--coverage-folder-file` coverage folder file under targetFolder, does not affect files that are not in the
      template (default: false)
    - conventional contributing support `--language`
        - `en-US`
        - `zh-CN`
- [x] `copilot` add github copilot config (1.18.+)
    - `--init-commit` init commit message config, default: false
- [x] add `--html-img` add html img badge (1.16.+)
- [X] `action` fast add github action workflow (1.10.+), must set `--ci-*` to effective
    - `--coverage-folder-file` coverage folder or file under targetFolder, does not affect files that are not in the
      template (default: false)
    - `--ci-deploy-tag` add sample deploy by tag
- [ ] more perfect test case coverage

## usage

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/sinlov/gh-conventional-kit/cmd/gh-conventional-kit@latest
# install version v1.11.2
$ go install -v github.com/sinlov/gh-conventional-kit/cmd/gh-conventional-kit@v1.12.0

# usa as docker cli tools
$ docker run --rm sinlov/gh-conventional-kit:latest -h
# use as docker cli tools with version
$ docker run --rm sinlov/gh-conventional-kit:1.12.0 -h
```

- please install [git](https://git-scm.com/) before use this cli

```bash
## show commands and global options
$ gh-conventional-kit --help

## markdown tools help
$ gh-conventional-kit markdown -h
# show common badges by gitUrl 1.6.+ support
$ gh-conventional-kit markdown --golang <gitUrl>
# example
$ gh-conventional-kit markdown --golang git@github.com:sinlov/gh-conventional-kit.git
$ gh-conventional-kit markdown --golang https://github.com/sinlov/gh-conventional-kit.git

# when project is golang
$ gh-conventional-kit markdown --golang <gitUrl>
# when project is rust
$ gh-conventional-kit markdown --rust <gitUrl>
# crates name not same as repo name
$ gh-conventional-kit markdown --rust --rust-crates some-rs <gitUrl>
# when project is node
$ gh-conventional-kit markdown --node <gitUrl>
# multiple programming languages
$ gh-conventional-kit markdown --golang --node <gitUrl>
# show common badges by -u and r
$ gh-conventional-kit markdown -u [user] -r [repo]

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

## copilot tools help (1.18.+)
$ gh-conventional-kit copilot --help
# try add copilot git commit message config
$ gh-conventional-kit copilot --init-commit
# coverage folder file is exist
$ gh-conventional-kit copilot --coverage-folder-file --init-commit
# init copilot git commit message config with language default en-US
$ gh-conventional-kit copilot --init-commit --language zh-CN

## template tools help
$ gh-conventional-kit template --help
# try add template at .github and try add badge at README.md
$ gh-conventional-kit --dry-run template --language en-US,zh-CN
# this project is golang
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --golang
# this project is rust
$ gh-conventional-kit --dry-run template --language en-US,zh-CN --rust
# crates name not same as repo name
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

## dev

- see [doc/dev.md](doc/dev.md)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/gh-conventional-kit)](https://github.com/sinlov/gh-conventional-kit/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息