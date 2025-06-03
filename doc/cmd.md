## gh-conventional-kit

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