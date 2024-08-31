package embed_operator

import (
	"github.com/bar-counter/slog"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov/gh-conventional-kit/internal/embed_source"
	"os"
	"path/filepath"
	"strings"
)

func WriteFileByEmbedResources(resourceList []embed_source.EmbedResource, root string, isCoverage bool, innerPath string, replacePath string, addFileTag string) error {
	for _, embedRes := range resourceList {
		if embedRes.IsDir() {
			continue
		}
		err := WriteFileByEmbedResource(embedRes, root, isCoverage, innerPath, replacePath, addFileTag)
		if err != nil {
			return err
		}

	}
	return nil
}

func WriteFileByEmbedResource(embedSource embed_source.EmbedResource, root string, isCoverage bool, innerPath string, replacePath string, addFileTag string) error {
	relativePath := embedSource.RelativePath()
	pathReplace := strings.Replace(relativePath, innerPath, replacePath, 1)
	targetPath := filepath.Join(root, pathReplace)
	if filepath_plus.PathExistsFast(targetPath) {
		if isCoverage {
			bytes, err := embedSource.Raw()
			if err != nil {
				return err
			}
			errWrite := filepath_plus.WriteFileByByte(targetPath, bytes, os.FileMode(0o666), true)
			if errWrite != nil {
				return errWrite
			}
			slog.Infof("-> rewrite by embed file %s file: %s", addFileTag, targetPath)
		} else {
			slog.Debugf("-> skip by embed file exists %s path: %s", addFileTag, targetPath)
		}
	} else {
		bytes, err := embedSource.Raw()
		if err != nil {
			return err
		}
		errWrite := filepath_plus.WriteFileByByte(targetPath, bytes, os.FileMode(0o666), false)
		if errWrite != nil {
			return errWrite
		}
		slog.Infof("-> add by embed file %s path: %s", addFileTag, targetPath)
	}
	return nil
}
