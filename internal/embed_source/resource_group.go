package embed_source

import (
	"embed"
	"fmt"
	"os"
	"path"
	"sort"
)

var (
	ErrorNotSetResourceRootPath = fmt.Errorf("not set resource root path")

	markKey          = string(os.PathListSeparator)
	resourceDirKey   = fmt.Sprintf("%s%s%s", markKey, markKey, markKey)
	resourceGroup    = map[string][]EmbedResource{}
	resourceRootPath = ""
)

func SettingResourceRootPath(root string) {
	if len(resourceRootPath) == 0 {
		if !pathExistsFast(root) {
			panic(fmt.Sprintf("SettingResourceRootPath not exists root %s", root))
		}
		resourceRootPath = root
		return
	}
	panic(fmt.Sprintf("SettingResourceRootPath already set root %s", resourceRootPath))
}

func checkResourceRootPathNotSet() bool {
	return len(resourceRootPath) == 0
}

// GetResourceByLanguage
// this method will check lang and parent dir
// if not found will use default lang with GetResourceByLanguageDefault
//
//	group: resource group, most of the use path of resource folder
//	innerPath: inner path of resource folder
//	lang: language as i18n
//
// return EmbedResource
func GetResourceByLanguage(group string, innerPath string, lang string) (EmbedResource, error) {
	checkPath := path.Join(group, innerPath)
	parentDirPath := path.Dir(checkPath)
	fileName := path.Base(checkPath)
	wantKey := fmt.Sprintf("%s%s%s%s%s", group, markKey, lang, markKey, parentDirPath)
	resources, ok := resourceGroup[wantKey]
	if !ok {
		return GetResourceByLanguageDefault(group, innerPath)
	}
	if len(resources) == 0 {
		return nil, fmt.Errorf("not found GetResourceListByLanguage group %s lang %s item %s", group, lang, innerPath)
	}
	var target EmbedResource
	for _, resource := range resources {
		if resource.Lang() != "" && resource.Lang() != lang {
			continue
		}
		if resource.FileName() == fileName {
			target = resource
		}
	}

	if target != nil {
		return target, nil
	}

	return GetResourceByLanguageDefault(group, innerPath)
}

func GetResourceByLanguageDefault(group string, innerPath string) (EmbedResource, error) {
	checkPath := path.Join(group, innerPath)
	parentDirPath := path.Dir(checkPath)
	fileName := path.Base(checkPath)
	defaultKey := fmt.Sprintf("%s%s%s%s", group, markKey, markKey, parentDirPath)
	resourcesDefault, okDefault := resourceGroup[defaultKey]
	if !okDefault {
		return nil, fmt.Errorf("not found GetResourceByLanguageDefault group %s item %s, as item %s", group, innerPath, defaultKey)
	}
	if len(resourcesDefault) == 0 {
		return nil, fmt.Errorf("not found GetResourceByLanguageDefault group %s item %s", group, innerPath)
	}
	for _, resource := range resourcesDefault {
		if resource.FileName() == fileName {
			return resource, nil
		}
	}
	return nil, fmt.Errorf("not found GetResourceByLanguageDefault group %s item %s", group, innerPath)
}

func GetResourceListByLanguage(group string, innerScanPath string, lang string) ([]EmbedResource, error) {
	relativePath := path.Join(group, innerScanPath)
	wantKey := fmt.Sprintf("%s%s%s%s%s", group, markKey, lang, markKey, relativePath)
	resources, ok := resourceGroup[wantKey]
	if !ok {
		defaultKey := fmt.Sprintf("%s%s%s%s", group, markKey, markKey, relativePath)
		resourcesDefault, okDefault := resourceGroup[defaultKey]
		if !okDefault {
			return nil, fmt.Errorf("not found GetResourceListByLanguage group %s lang %s item %s, as item %s", group, lang, innerScanPath, defaultKey)
		}
		return resourcesDefault, nil
	}
	return resources, nil
}

// InitResourceGroupByLanguage
//
//	group: resource group, most of the use path of resource folder
//	embFs: embed.FS some as group package by go:embed
//	innerScanPath: inner scan path, will use "" will scan group current path
//	langList: lang load list
func InitResourceGroupByLanguage(group string, embFs embed.FS, innerScanPath string, langList []string) error {
	if checkResourceRootPathNotSet() {
		return ErrorNotSetResourceRootPath
	}
	// check inner language resource
	for _, lang := range langList {
		if len(lang) == 0 {
			return fmt.Errorf("check InitResourceGroupByLanguage langList item is empty")
		}
	}

	checkPath := path.Join(group, innerScanPath)

	embOpen, errEmbOpen := embFs.Open(checkPath)
	if errEmbOpen != nil {
		return fmt.Errorf("check InitResourceGroupByLanguage %s err: %s", innerScanPath, errEmbOpen)
	}
	embOpenInfo, errStat := embOpen.Stat()
	if errStat != nil {
		return fmt.Errorf("check InitResourceGroupByLanguage %s err: %s", innerScanPath, errStat)
	}

	if embOpenInfo.IsDir() {
		targetDirs, err := embFs.ReadDir(checkPath)
		if err != nil {
			return fmt.Errorf("check InitResourceGroupByLanguage read dir %s err: %s", innerScanPath, err)
		}
		var defaultItems []string
		for _, dirItem := range targetDirs {
			if dirItem.IsDir() {
				continue
			}
			itemKeyName := dirItem.Name()
			relativePath := path.Join(group, innerScanPath, itemKeyName)
			resourceFile, errNewResAsFile := NewEmbedResourceAsFile(resourceRootPath, embFs, relativePath, "", false)
			if errNewResAsFile != nil {
				return fmt.Errorf("default InitResourceGroupByLanguage %s err %s", relativePath, errNewResAsFile)
			}
			newKey := fmt.Sprintf("%s%s%s%s", group, markKey, markKey, checkPath)
			resourceGroup[newKey] = append(resourceGroup[newKey], resourceFile)
			defaultItems = append(defaultItems, itemKeyName)
		}
		if len(defaultItems) > 0 {
			for _, dirItem := range targetDirs {
				if !dirItem.IsDir() {
					continue
				}
				langDirName := dirItem.Name()
				i18nDir := path.Join(checkPath, langDirName)
				if !strInArr(langDirName, langList) {
					//return fmt.Errorf("check InitResourceGroupByLanguage langDirName %s not setting in path %s", langDirName, i18nDir)
					continue
				}
				langTargetDirs, errLangRead := embFs.ReadDir(i18nDir)
				if errLangRead != nil {
					//return fmt.Errorf("check InitResourceGroupByLanguage %s err: %s", i18nDir, errLangRead)
					continue
				}
				for _, i18nInnerDir := range langTargetDirs {
					if i18nInnerDir.IsDir() {
						continue
					}
					itemKeyName := i18nInnerDir.Name()
					if !strInArr(itemKeyName, defaultItems) {
						//return fmt.Errorf("check InitResourceGroupByLanguage langDirName %s not has itemKeyName %s in path %s", langDirName, itemKeyName, i18nDir)
						continue
					}
					relativePath := path.Join(group, innerScanPath, langDirName, itemKeyName)
					resourceFile, errNewResAsFile := NewEmbedResourceAsFile(resourceRootPath, embFs, relativePath, langDirName, false)
					if errNewResAsFile != nil {
						return fmt.Errorf("load InitResourceGroupByLanguage langDirName %s %s err %s", langDirName, relativePath, errNewResAsFile)
					}

					newKey := fmt.Sprintf("%s%s%s%s%s", group, markKey, langDirName, markKey, checkPath)
					resourceGroup[newKey] = append(resourceGroup[newKey], resourceFile)
				}
			}
		}

	} else {
		parentDirPath := path.Dir(checkPath)
		defaultItem := path.Base(checkPath)
		relativePath := path.Join(group, innerScanPath)
		resourceFile, errNewResAsFile := NewEmbedResourceAsFile(resourceRootPath, embFs, relativePath, "", false)
		if errNewResAsFile != nil {
			return fmt.Errorf("default InitResourceGroupByLanguage %s err %s", relativePath, errNewResAsFile)
		}
		newKey := fmt.Sprintf("%s%s%s%s", group, markKey, markKey, parentDirPath)
		resourceGroup[newKey] = append(resourceGroup[newKey], resourceFile)
		parentDirs, errReadParentDir := embFs.ReadDir(parentDirPath)
		if errReadParentDir == nil {
			for _, parentDir := range parentDirs {
				if !parentDir.IsDir() {
					continue
				}
				langDirName := parentDir.Name()
				if !strInArr(langDirName, langList) {
					continue
				}
				i18nDir := path.Join(parentDirPath, langDirName)
				langTargetDirs, errLangRead := embFs.ReadDir(i18nDir)
				if errLangRead != nil {
					//return fmt.Errorf("check InitResourceGroupByLanguage %s err: %s", i18nDir, errLangRead)
					continue
				}
				for _, i18nInnerDir := range langTargetDirs {
					if i18nInnerDir.IsDir() {
						continue
					}
					itemKeyName := i18nInnerDir.Name()
					if itemKeyName != defaultItem {
						continue
					}
					relativePathLang := path.Join(i18nDir, itemKeyName)
					resourceFileLang, errNewResAsFileLang := NewEmbedResourceAsFile(resourceRootPath, embFs, relativePathLang, langDirName, false)
					if errNewResAsFileLang != nil {
						return fmt.Errorf("load InitResourceGroupByLanguage langDirName %s %s err %s", langDirName, relativePathLang, errNewResAsFileLang)
					}

					newKeyLang := fmt.Sprintf("%s%s%s%s%s", group, markKey, langDirName, markKey, parentDirPath)
					resourceGroup[newKeyLang] = append(resourceGroup[newKeyLang], resourceFileLang)
				}
			}
		}

	}

	return nil
}

// InitResourceByDir
// this function only load first folder then bind second file as EmbedResource list
//
//	group: resource group, most of the use path of resource folder
//	embFs: embed.FS some as group package by go:embed
//	innerScanPath: inner scan path, will use "" will scan group current path
func InitResourceByDir(group string, embFs embed.FS, innerScanPathList []string) error {
	if checkResourceRootPathNotSet() {
		return ErrorNotSetResourceRootPath
	}

	if len(innerScanPathList) == 0 {
		return fmt.Errorf("check InitResourceByDir %s innerScanPathList is empty", group)
	}

	for _, itemPath := range innerScanPathList {
		targetDirs, errReadDir := embFs.ReadDir(itemPath)
		if errReadDir != nil {
			return fmt.Errorf("check InitResourceByDir %s err %s", itemPath, errReadDir)
		}
		if len(targetDirs) == 0 {
			return fmt.Errorf("check InitResourceByDir %s is empty", itemPath)
		}
		innerPath := itemPath
		resource, errNewRes := NewEmbedResourceAsDir(resourceRootPath, embFs, innerPath, "", false)
		if errNewRes != nil {
			return fmt.Errorf("InitResourceByDir %s err %s", innerPath, errNewRes)
		}
		newKey := fmt.Sprintf("%s%s%s", group, resourceDirKey, innerPath)
		resourceGroup[newKey] = append(resourceGroup[newKey], resource...)
	}

	return nil
}

func GetResourceGroupByDir(group string, innerScanPath string) ([]EmbedResource, error) {
	absKey := path.Join(group, innerScanPath)
	wantKey := fmt.Sprintf("%s%s%s", group, resourceDirKey, absKey)
	resources, ok := resourceGroup[wantKey]
	if !ok {
		return nil, fmt.Errorf("not found GetResourceGroupByDir group %s, by want key %s", group, absKey)
	}
	return resources, nil
}

func strInArr(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// PathExistsFast
//
//	path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// pathExists
//
//	path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
