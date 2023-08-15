package gh_conventional_kit

import (
	"embed"
	_ "embed"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"github.com/sinlov/gh-conventional-kit/resource"
	"path"
)

//go:embed package.json
var PackageJson string

func CheckAllResource(root string) error {

	embed_source.SettingResourceRootPath(root)

	err := embed_source.InitResourceByDir(resource.GroupResource, embedResourceFiles, embedDotGithubList)
	if err != nil {
		return err
	}

	for _, resActionItem := range embedResourceActionList {
		err = embed_source.InitResourceGroupByLanguage(resource.GroupResourceAction, embedResourceActionFiles, resActionItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	for _, resPathItem := range embedResourceCCDocPathList {
		err = embed_source.InitResourceGroupByLanguage(resource.GroupResource, embedResourceContributingDocFiles, resPathItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	return nil
}

var (
	//go:embed resource
	embedResourceFiles embed.FS

	embedDotGithubList = []string{
		path.Join(resource.GroupResource, resource.DirGithubAction, resource.KeyIssueTemplate),
	}

	//go:embed resource/action
	embedResourceActionFiles embed.FS

	embedResourceActionList = []string{
		resource.KeyContributingDoc,
		resource.KeyPullRequestTemplate,
		resource.KeyDependabotConfig,
	}

	//go:embed resource/contributing_doc
	embedResourceContributingDocFiles embed.FS

	embedResourceCCDocPathList = []string{
		path.Join(resource.DirNameContributingDoc, resource.KeyConventionalReadmeTitle),
		path.Join(resource.DirNameContributingDoc, resource.KeyConventionalReadmeI18n),
	}
)
