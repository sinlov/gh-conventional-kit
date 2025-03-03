# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.17.1](https://github.com/sinlov/gh-conventional-kit/compare/1.17.0...v1.17.1) (2024-11-25)

### 🐛 Bug Fixes

* update `badge-markdown` subcommand for Rust crates ([e6a69318](https://github.com/sinlov/gh-conventional-kit/commit/e6a69318c118183151bde89186be919e146cded7))

## [1.17.0](https://github.com/sinlov/gh-conventional-kit/compare/1.16.2...v1.17.0) (2024-11-24)

### ✨ Features

* update dependency badge to use GitHub reference ([507bc493](https://github.com/sinlov/gh-conventional-kit/commit/507bc493661e2923cee60e0990b090bf029e94db)), feat [#63](https://github.com/sinlov/gh-conventional-kit/issues/63)

### 👷‍ Build System

* bump codecov/codecov-action from 4.5.0 to 4.6.0 ([f1c2bd94](https://github.com/sinlov/gh-conventional-kit/commit/f1c2bd9468239b796f3baff5880de10bc37d41c7))

## [1.16.2](https://github.com/sinlov/gh-conventional-kit/compare/1.16.1...v1.16.2) (2024-08-31)

### 📝 Documentation

* update resouce of doc for `CONTRIBUTING_DOC` ([60cef99a](https://github.com/sinlov/gh-conventional-kit/commit/60cef99aa4110a75daa8bfed003d5d62d70d0a77))

* update `CONTRIBUTING_DOC` link ([cb2c0acc](https://github.com/sinlov/gh-conventional-kit/commit/cb2c0acced67b89499587df72ac72612de053bef)), doc [#60](https://github.com/sinlov/gh-conventional-kit/issues/60)

## [1.16.1](https://github.com/sinlov/gh-conventional-kit/compare/1.16.0...v1.16.1) (2024-08-31)

### 🐛 Bug Fixes

* replace to 0o666 and use github.com/sinlov-go/go-common-lib ([ea92514d](https://github.com/sinlov/gh-conventional-kit/commit/ea92514dde721b382f8bb62d1ce643b432355b0c)), fix [#58](https://github.com/sinlov/gh-conventional-kit/issues/58)

### 📝 Documentation

* add usage of flag `--html-img` ([8cbccfca](https://github.com/sinlov/gh-conventional-kit/commit/8cbccfcacf00ab4043811a587ef2ceb3de9274dd))

## [1.16.0](https://github.com/sinlov/gh-conventional-kit/compare/1.15.0...v1.16.0) (2024-08-26)

### ✨ Features

* add HTML markdown support for badges by flag `--html-img` ([f373d06c](https://github.com/sinlov/gh-conventional-kit/commit/f373d06c3419e4cc0655d9e4438130c84b0217b3)), fe [#56](https://github.com/sinlov/gh-conventional-kit/issues/56)

### 👷‍ Build System

* bump github.com/sebdah/goldie/v2 from 2.5.3 to 2.5.5 ([4df1cd7f](https://github.com/sinlov/gh-conventional-kit/commit/4df1cd7f060cd74efda93b88eca80d83fc5f3ae8))

* upgrade github.com/sinlov-go/badges to v1.4.0Update the github.com/sinlov-go/badges dependency from v1.3.2 to v1.4.0 for improved ([fb28c237](https://github.com/sinlov/gh-conventional-kit/commit/fb28c23782dbc0d0341acf0641040cbeaf6083c1))

## [1.15.0](https://github.com/sinlov/gh-conventional-kit/compare/1.14.0...v1.15.0) (2024-08-01)

### ✨ Features

* change file template for change `dependabot.yml` for `commit-message` ([776d07e2](https://github.com/sinlov/gh-conventional-kit/commit/776d07e2b68281622c45d07f52954e6adc9e6f67)), feat [#53](https://github.com/sinlov/gh-conventional-kit/issues/53)

### 👷‍ Build System

* add `main.buildID` inject at build pipeline ([f33e7486](https://github.com/sinlov/gh-conventional-kit/commit/f33e7486df5bd64977bcb5870ba45ab3185688e0))

## [1.14.0](https://github.com/sinlov/gh-conventional-kit/compare/1.13.0...v1.14.0) (2024-05-20)

### ✨ Features

* update docker bake-hubdocker build pipline ([4b9cce79](https://github.com/sinlov/gh-conventional-kit/commit/4b9cce796f875781b73421df5883d342b1403442)), feat [#46](https://github.com/sinlov/gh-conventional-kit/issues/46)

* github action auto depends update, full make task update ([000455b9](https://github.com/sinlov/gh-conventional-kit/commit/000455b9613675a56e2be68462e9f081075d8db3)), feat [#46](https://github.com/sinlov/gh-conventional-kit/issues/46)

## [1.13.0](https://github.com/sinlov/gh-conventional-kit/compare/1.12.0...v1.13.0) (2024-04-28)

### ✨ Features

* pipeline change to use docker baka ([fc7b3168](https://github.com/sinlov/gh-conventional-kit/commit/fc7b3168746391514a6d47503bf63fffda0a1369))

### 🐛 Bug Fixes

* config file `docker-bake.hcl` add ([f77093fe](https://github.com/sinlov/gh-conventional-kit/commit/f77093fef834b9b69c4b1ae8d0705d527127a15c))

## [1.12.0](https://github.com/sinlov/gh-conventional-kit/compare/1.11.2...v1.12.0) (2024-03-27)

### ✨ Features

* markdown support find by current folder to get git info ([e0a1d1b8](https://github.com/sinlov/gh-conventional-kit/commit/e0a1d1b8065756cc7a5ef38ac2c63f7d4b3c9ec2))

* add flag `--codecov-badges` to open show codecov.io badge ([4f5fd059](https://github.com/sinlov/gh-conventional-kit/commit/4f5fd059239ef572665a25d9eef1d83cb533e4fc))

## [1.11.2](https://github.com/sinlov/gh-conventional-kit/compare/1.11.1...v1.11.2) (2024-01-26)

### 🐛 Bug Fixes

* rust badge docs.rs url error update github.com/sinlov-go/badges v1.3.2 ([34219ad3](https://github.com/sinlov/gh-conventional-kit/commit/34219ad3947d267bea608c42884d2fb44fbc8c65)), fix [#31](https://github.com/sinlov/gh-conventional-kit/issues/31)

### 📝 Documentation

* update usage doc to v1.11.2 ([7d583a9f](https://github.com/sinlov/gh-conventional-kit/commit/7d583a9ff89dc64f6d84051c615e2db597d6be86))

## [1.11.1](https://github.com/sinlov/gh-conventional-kit/compare/1.11.0...v1.11.1) (2024-01-26)

### 📝 Documentation

* update usage of Features, support actions/download-artifact@v4 ([7d6dd193](https://github.com/sinlov/gh-conventional-kit/commit/7d6dd193151c432247044be9042de4b82d19215c)), doc [#29](https://github.com/sinlov/gh-conventional-kit/issues/29)

## [1.11.0](https://github.com/sinlov/gh-conventional-kit/compare/1.10.0...v1.11.0) (2024-01-23)

### ✨ Features

* change to use github.com/chainguard-dev/git-urls v1.0.2 ([b62b1cc7](https://github.com/sinlov/gh-conventional-kit/commit/b62b1cc7160ec3cfb6679b082f047d7bbedcc884))

* add subcommand action ([545be1b0](https://github.com/sinlov/gh-conventional-kit/commit/545be1b0403a5e4c2a04446cdbe190af59393e7c)), feature [#23](https://github.com/sinlov/gh-conventional-kit/issues/23)

* change pull_request_template.md be sample ([2bca4de8](https://github.com/sinlov/gh-conventional-kit/commit/2bca4de8cece66df1eb9ab0c54766e2a8f621e27))

* update template of issue support form ([20f697ae](https://github.com/sinlov/gh-conventional-kit/commit/20f697ae7ba62f99e67ae1cc3f14b6ccf9a45e6f)), feat [#13](https://github.com/sinlov/gh-conventional-kit/issues/13)

* workflows-template can render by git info ([4f5714f9](https://github.com/sinlov/gh-conventional-kit/commit/4f5714f9fbb05398156cc25a8712cb80fdb3c198))

* embed_source package to management template resource ([13f4b6df](https://github.com/sinlov/gh-conventional-kit/commit/13f4b6df34a98b1ebfb0f23f799949c219b43908)), feat [#10](https://github.com/sinlov/gh-conventional-kit/issues/10)

* let markdown each group have return and update dependabot.yml resource` ([e84b6918](https://github.com/sinlov/gh-conventional-kit/commit/e84b69186e27a4c00aa5bf74b6d7b1c7e3528dea))

* let subcommond markdown support <gitUrl> for fast got git info ([88d8e254](https://github.com/sinlov/gh-conventional-kit/commit/88d8e254b41c352c7efe5cda6205ef80fcac7a42))

* add go verison show and change versionrc.json ([47bf4187](https://github.com/sinlov/gh-conventional-kit/commit/47bf4187ac1771c2e49f77e54fdecf64b39107b8))

* update template file and add notice at --dry-run ([38ae66fe](https://github.com/sinlov/gh-conventional-kit/commit/38ae66fe27a556e3b24405d91593a99de6762274))

* add dependabot.yml at github_template when use commands `template` ([47b75dfc](https://github.com/sinlov/gh-conventional-kit/commit/47b75dfc415fd3fd10fa1cdff95b0718f2110c47))

* add --rust-crates to support rust crates name different from repo name ([e0d0dba2](https://github.com/sinlov/gh-conventional-kit/commit/e0d0dba24f76344a5f427aa947005f978de39705))

* add docker image for not install go-sdk useage ([145e4d7f](https://github.com/sinlov/gh-conventional-kit/commit/145e4d7f44bf9cbaee5570e74d81c6798f8f39ee))

* change to github.com/sinlov-go/go-git-tools v1.0.0 ([6c53aeb4](https://github.com/sinlov/gh-conventional-kit/commit/6c53aeb4cfb42890768b28a499110a27932e02f4))

* --coverage-folder-file for template, change install folder for this cli ([70c67a47](https://github.com/sinlov/gh-conventional-kit/commit/70c67a47bdaa4888ef5cda23ac34bcb09010dd68))

* remove useless template noversion rc and let bages support --no-markdown at append README.md ([f18babbc](https://github.com/sinlov/gh-conventional-kit/commit/f18babbcc1a5b5fafd98a58bebfa262fc11af4f5))

* update help of cli and let `badge` can append README.md ([16fccbd3](https://github.com/sinlov/gh-conventional-kit/commit/16fccbd30b8127dba847847409721d78bfc85551))

* add template command support --language ([c2592554](https://github.com/sinlov/gh-conventional-kit/commit/c25925549b0d2de7ffbbaf94f72054774f116b21))

* add TemplateGitRootWalk and TemplateGithubDotWalk can try add template file by embed ([ed672486](https://github.com/sinlov/gh-conventional-kit/commit/ed67248655cd50cbb0fb783269ce1e3095b431d9))

* add template can add README.md info ([5353a626](https://github.com/sinlov/gh-conventional-kit/commit/5353a6263b7d099852d86e695c3cc69ebea61cec))

* badge commond can gen by git workspace root folder ([aa04173c](https://github.com/sinlov/gh-conventional-kit/commit/aa04173c18b34f8ba925b5720ac509c9dc453898))

* add badge node ([b37a97a2](https://github.com/sinlov/gh-conventional-kit/commit/b37a97a25819f0a600f02a547883cede7653d978))

* github.com/sinlov-go/badges@1.1.0 ([b2ea8941](https://github.com/sinlov/gh-conventional-kit/commit/b2ea8941246ee27d96b226dd8a2ae4a51dac175c))

* use lib github.com/sinlov-go/badges at markdown subcommand ([2c1af69b](https://github.com/sinlov/gh-conventional-kit/commit/2c1af69bc867fc2e5fc01893d702a1f9b888f580))

* add subCmdMarkdown common github info ([0f5c6baa](https://github.com/sinlov/gh-conventional-kit/commit/0f5c6baa51da5953ea3d430ce51273c21919bb68))

### 🐛 Bug Fixes

* fix windows inner path will be replace error ([66240198](https://github.com/sinlov/gh-conventional-kit/commit/66240198d54cdf546e159ea82ebc16919a96f831))

* change github.com/sinlov-go/go-git-tools v1.6.0 to fix local parse error ([e2e8905b](https://github.com/sinlov/gh-conventional-kit/commit/e2e8905b41f94d371574341d3de803a30b3479ed))

* common_subcommand.PrintBadgeByConfigWithMarkdown not append return ([54b5f8e8](https://github.com/sinlov/gh-conventional-kit/commit/54b5f8e83a44f7335c67c9e135d9a971fd7c5778))

### ♻ Refactor

* rename util package to internal ([dd782237](https://github.com/sinlov/gh-conventional-kit/commit/dd78223754857f3a967b0a70cc3cc9af59459da6))

* show more run info and move to github.com/sinlov/gh-conventional-kit/cmd/cli ([8375fa30](https://github.com/sinlov/gh-conventional-kit/commit/8375fa304b92e0da4221cc26429b6de57072c4aa))

### 👷‍ Build System

* support actions/upload-artifact/tree/v4 ([ea2be1fc](https://github.com/sinlov/gh-conventional-kit/commit/ea2be1fc2572bd227bea6ff79a10bbb47eff0b8c))

* change golangci/golangci-lint-action use version latest ([1c71e7af](https://github.com/sinlov/gh-conventional-kit/commit/1c71e7af7864c4eebce14e38f75f3287e156254f))

* change build goversion1.19 for base of build ([b0147e03](https://github.com/sinlov/gh-conventional-kit/commit/b0147e034e0966a82533a80903dd3759ba76ecf4))

* try fix build of verison ([e4bd0bda](https://github.com/sinlov/gh-conventional-kit/commit/e4bd0bda3ebd96ff80dec7034152e8bd82cf63ba))

* fix github action echo words error ([635b78a3](https://github.com/sinlov/gh-conventional-kit/commit/635b78a3278adebb12824ac5c0fcf7a44874a967))

* change new version of CI pipline ([f595eac2](https://github.com/sinlov/gh-conventional-kit/commit/f595eac23f2647cdb107f1da9a68f5a637713de9))

* github.com/sinlov-go/badges v1.3.1 ([82445027](https://github.com/sinlov/gh-conventional-kit/commit/824450274064d5f3deeea4aa4b4dfc9321a02fd7))

* change to new build workflows by ci.yml ([d8fc77b4](https://github.com/sinlov/gh-conventional-kit/commit/d8fc77b4a71c575015c1c253cea4b3a38cb24c55))

## [1.10.0](https://github.com/sinlov/gh-conventional-kit/compare/1.9.0...v1.10.0) (2023-12-29)

### ✨ Features

* add subcommand action ([75eea7e8](https://github.com/sinlov/gh-conventional-kit/commit/75eea7e8e7ed3ee9a216a918734c5a1de9057f2f)), feature [#23](https://github.com/sinlov/gh-conventional-kit/issues/23)

### 👷‍ Build System

* change build goversion1.19 for base of build ([b0147e03](https://github.com/sinlov/gh-conventional-kit/commit/b0147e034e0966a82533a80903dd3759ba76ecf4))

## [1.9.0](https://github.com/sinlov/gh-conventional-kit/compare/1.8.0...v1.9.0) (2023-09-24)

### ✨ Features

* change pull_request_template.md be sample ([2bca4de8](https://github.com/sinlov/gh-conventional-kit/commit/2bca4de8cece66df1eb9ab0c54766e2a8f621e27))

* update template of issue support form ([20f697ae](https://github.com/sinlov/gh-conventional-kit/commit/20f697ae7ba62f99e67ae1cc3f14b6ccf9a45e6f)), feat [#13](https://github.com/sinlov/gh-conventional-kit/issues/13)

## [1.8.0](https://github.com/sinlov/gh-conventional-kit/compare/1.7.1...v1.8.0) (2023-08-15)

### ✨ Features

* workflows-template can render by git info ([4f5714f9](https://github.com/sinlov/gh-conventional-kit/commit/4f5714f9fbb05398156cc25a8712cb80fdb3c198))

* embed_source package to management template resource ([13f4b6df](https://github.com/sinlov/gh-conventional-kit/commit/13f4b6df34a98b1ebfb0f23f799949c219b43908)), feat [#10](https://github.com/sinlov/gh-conventional-kit/issues/10)

### 🐛 Bug Fixes

* fix windows inner path will be replace error ([66240198](https://github.com/sinlov/gh-conventional-kit/commit/66240198d54cdf546e159ea82ebc16919a96f831))

### ♻ Refactor

* rename util package to internal ([dd782237](https://github.com/sinlov/gh-conventional-kit/commit/dd78223754857f3a967b0a70cc3cc9af59459da6))

## [1.7.1](https://github.com/sinlov/gh-conventional-kit/compare/1.7.0...v1.7.1) (2023-08-05)

### 👷‍ Build System

* try fix build of verison ([e4bd0bda](https://github.com/sinlov/gh-conventional-kit/commit/e4bd0bda3ebd96ff80dec7034152e8bd82cf63ba))

* fix github action echo words error ([635b78a3](https://github.com/sinlov/gh-conventional-kit/commit/635b78a3278adebb12824ac5c0fcf7a44874a967))

## [1.7.0](https://github.com/sinlov/gh-conventional-kit/compare/1.6.0...v1.7.0) (2023-08-05)

### ✨ Features

* let markdown each group have return and update dependabot.yml resource` ([e84b6918](https://github.com/sinlov/gh-conventional-kit/commit/e84b69186e27a4c00aa5bf74b6d7b1c7e3528dea))

## [1.6.0](https://github.com/sinlov/gh-conventional-kit/compare/1.5.0...v1.6.0) (2023-08-01)

### ✨ Features

* let subcommond markdown support <gitUrl> for fast got git info ([88d8e254](https://github.com/sinlov/gh-conventional-kit/commit/88d8e254b41c352c7efe5cda6205ef80fcac7a42))

* add go verison show and change versionrc.json ([47bf4187](https://github.com/sinlov/gh-conventional-kit/commit/47bf4187ac1771c2e49f77e54fdecf64b39107b8))

### ♻ Refactor

* show more run info and move to github.com/sinlov/gh-conventional-kit/cmd/cli ([8375fa30](https://github.com/sinlov/gh-conventional-kit/commit/8375fa304b92e0da4221cc26429b6de57072c4aa))

### 👷‍ Build System

* change new version of CI pipline ([f595eac2](https://github.com/sinlov/gh-conventional-kit/commit/f595eac23f2647cdb107f1da9a68f5a637713de9))

## [1.5.0](https://github.com/sinlov/gh-conventional-kit/compare/v1.4.0...v1.5.0) (2023-07-20)

### 🐛 Bug Fixes

* change github.com/sinlov-go/go-git-tools v1.6.0 to fix local parse error ([e2e8905](https://github.com/sinlov/gh-conventional-kit/commit/e2e8905b41f94d371574341d3de803a30b3479ed))

## [1.4.0](https://github.com/sinlov/gh-conventional-kit/compare/v1.3.0...v1.4.0) (2023-07-20)

### ✨ Features

* update template file and add notice at --dry-run ([38ae66f](https://github.com/sinlov/gh-conventional-kit/commit/38ae66fe27a556e3b24405d91593a99de6762274))

## [1.3.0](https://github.com/sinlov/gh-conventional-kit/compare/v1.2.0...v1.3.0) (2023-07-18)

### ✨ Features

* add dependabot.yml at github_template when use commands `template` ([47b75df](https://github.com/sinlov/gh-conventional-kit/commit/47b75dfc415fd3fd10fa1cdff95b0718f2110c47))

### 👷‍ Build System

* change to new build workflows by ci.yml ([d8fc77b](https://github.com/sinlov/gh-conventional-kit/commit/d8fc77b4a71c575015c1c253cea4b3a38cb24c55))

* **gomod:** github.com/sinlov-go/badges v1.3.1 ([8244502](https://github.com/sinlov/gh-conventional-kit/commit/824450274064d5f3deeea4aa4b4dfc9321a02fd7))

## [1.2.0](https://github.com/sinlov/gh-conventional-kit/compare/v1.1.1...v1.2.0) (2023-07-11)

### ✨ Features

* add --rust-crates to support rust crates name different from repo name ([e0d0dba](https://github.com/sinlov/gh-conventional-kit/commit/e0d0dba24f76344a5f427aa947005f978de39705))

### [1.1.1](https://github.com/sinlov/gh-conventional-kit/compare/v1.1.0...v1.1.1) (2023-07-11)

### 🐛 Bug Fixes

* common_subcommand.PrintBadgeByConfigWithMarkdown not append return ([54b5f8e](https://github.com/sinlov/gh-conventional-kit/commit/54b5f8e83a44f7335c67c9e135d9a971fd7c5778))

## [1.1.0](https://github.com/sinlov/gh-conventional-kit/compare/v1.0.0...v1.1.0) (2023-07-05)

### ✨ Features

* add docker image for not install go-sdk useage ([145e4d7](https://github.com/sinlov/gh-conventional-kit/commit/145e4d7f44bf9cbaee5570e74d81c6798f8f39ee))

## 1.0.0 (2023-07-04)

### ✨ Features

* --coverage-folder-file for template, change install folder for this cli ([70c67a4](https://github.com/sinlov/gh-conventional-kit/commit/70c67a47bdaa4888ef5cda23ac34bcb09010dd68))

* add badge node ([b37a97a](https://github.com/sinlov/gh-conventional-kit/commit/b37a97a25819f0a600f02a547883cede7653d978))

* add subCmdMarkdown common github info ([0f5c6ba](https://github.com/sinlov/gh-conventional-kit/commit/0f5c6baa51da5953ea3d430ce51273c21919bb68))

* add template can add README.md info ([5353a62](https://github.com/sinlov/gh-conventional-kit/commit/5353a6263b7d099852d86e695c3cc69ebea61cec))

* add template command support --language ([c259255](https://github.com/sinlov/gh-conventional-kit/commit/c25925549b0d2de7ffbbaf94f72054774f116b21))

* add TemplateGitRootWalk and TemplateGithubDotWalk can try add template file by embed ([ed67248](https://github.com/sinlov/gh-conventional-kit/commit/ed67248655cd50cbb0fb783269ce1e3095b431d9))

* badge commond can gen by git workspace root folder ([aa04173](https://github.com/sinlov/gh-conventional-kit/commit/aa04173c18b34f8ba925b5720ac509c9dc453898))

* change to github.com/sinlov-go/go-git-tools v1.0.0 ([6c53aeb](https://github.com/sinlov/gh-conventional-kit/commit/6c53aeb4cfb42890768b28a499110a27932e02f4))

* github.com/sinlov-go/badges@1.1.0 ([b2ea894](https://github.com/sinlov/gh-conventional-kit/commit/b2ea8941246ee27d96b226dd8a2ae4a51dac175c))

* remove useless template noversion rc and let bages support --no-markdown at append README.md ([f18babb](https://github.com/sinlov/gh-conventional-kit/commit/f18babbcc1a5b5fafd98a58bebfa262fc11af4f5))

* update help of cli and let `badge` can append README.md ([16fccbd](https://github.com/sinlov/gh-conventional-kit/commit/16fccbd30b8127dba847847409721d78bfc85551))

* use lib github.com/sinlov-go/badges at markdown subcommand ([2c1af69](https://github.com/sinlov/gh-conventional-kit/commit/2c1af69bc867fc2e5fc01893d702a1f9b888f580))
