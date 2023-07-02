package gh_conventional_kit

import (
	_ "embed"
	_ "github.com/aymerick/raymond"
	"github.com/sinlov/gh-conventional-kit/resource/template_file"
)

//go:embed package.json
var PackageJson string

//go:embed resource/template_file/conventional_readme.md
var conventionalReadme string

func TemplateConventionalReadme(readme template_file.ConventionalReadme) (string, error) {
	return template_file.Render(conventionalReadme, readme)
}
