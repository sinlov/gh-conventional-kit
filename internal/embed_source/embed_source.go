package embed_source

import (
	"embed"
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"sync"
)

var (
	ErrPathNotFoundFile = errors.New("path not found file")
)

type EmbedResource interface {
	IsDir() bool
	RelativePath() string
	Lang() string

	ChangeRootPath(root string) error

	RootPath() string

	FullPath() string

	FileName() string

	Raw() ([]byte, error)

	Data() string

	Render(payload interface{}) (result string, err error)
}

type embedContentResource struct {
	isDir        bool
	relativePath string
	lang         string

	rootPath string
	fullPath string

	mutex    *sync.RWMutex
	efs      embed.FS
	fileName string

	hasLoad bool
	loadErr error
	raw     []byte
}

// IsDir
// this path is Dir
func (e *embedContentResource) IsDir() bool {
	return e.isDir
}

func (e *embedContentResource) Lang() string {
	return e.lang
}

// RelativePath
// this is in embed.FS path
func (e *embedContentResource) RelativePath() string {
	return e.relativePath
}

// FileName
// if path is not Dir will return file name
func (e *embedContentResource) FileName() string {
	return e.fileName
}

// RootPath
// visual path for biz, this not embed.FS path
func (e *embedContentResource) RootPath() string {
	return e.rootPath
}

// ChangeRootPath
// change visual path for biz, this not embed.FS path
func (e *embedContentResource) ChangeRootPath(root string) error {
	if len(root) == 0 {
		return fmt.Errorf("change root path empty")
	}
	e.rootPath = root
	e.fullPath = path.Join(e.rootPath, e.relativePath)
	return nil
}

// FullPath
// visual path will append RootPath, this not embed.FS path
func (e *embedContentResource) FullPath() string {
	return e.fullPath
}

// Raw
// only file can use,
// this method will try load once
func (e *embedContentResource) Raw() ([]byte, error) {
	if !e.hasLoad && !e.isDir {
		e.mutex.Lock()

		bytes, errRead := e.efs.ReadFile(e.relativePath)
		if errRead != nil {
			loadError := fmt.Errorf("read file %s: %w", e.relativePath, errRead)
			e.loadErr = loadError
			return nil, loadError
		}
		e.raw = bytes
		e.hasLoad = true
		e.mutex.Unlock()
	}

	if e.loadErr != nil {
		return nil, e.loadErr
	}

	return e.raw, nil
}

// Data
// when Raw() success will return string
// when Raw() fail will return ""
func (e *embedContentResource) Data() string {
	if !e.hasLoad {
		raw, errLoad := e.Raw()
		if errLoad != nil {
			return ""
		}
		return string(raw)
	}
	if e.raw == nil {
		return ""
	}
	return string(e.raw)
}

// Render
// when Raw() success will return Render by https://handlebarsjs.com/ 3.0 as golang
func (e *embedContentResource) Render(payload interface{}) (result string, err error) {
	if e.isDir {
		return "", fmt.Errorf("Render not support dir %s", e.fullPath)
	}
	return Render(e.Data(), payload)
}

// NewEmbedResourceAsFile
// warning: will load file content to memory
//
//	loadNow: true will load to memory struct
func NewEmbedResourceAsFile(root string, efs embed.FS, pathFile string, lang string, loadNow bool) (EmbedResource, error) {
	_, err := efs.Open(pathFile)
	if err != nil {
		return nil, err
	}
	if loadNow {
		bytes, err := efs.ReadFile(pathFile)
		if err != nil {
			return nil, fmt.Errorf("read file %s: %w", pathFile, err)
		}
		return newResourceAsFile(root, pathFile, bytes, efs, lang), nil
	}
	return newResourceAsFile(root, pathFile, nil, efs, lang), nil
}

// NewEmbedResourceAsDir
// warning: will load file content to memory
//
//	loadNow: true will load to memory struct
func NewEmbedResourceAsDir(root string, efs embed.FS, pathDir string, lang string, loadNow bool) ([]EmbedResource, error) {

	dirEntries, err := efs.ReadDir(pathDir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", pathDir, err)
	}
	if len(dirEntries) == 0 {
		return nil, ErrPathNotFoundFile
	}

	var res []EmbedResource
	for _, entry := range dirEntries {
		innerPath := path.Join(pathDir, entry.Name())
		var item EmbedResource
		if entry.IsDir() {
			item = newResourceAsDir(root, innerPath, efs)
		} else {
			if loadNow {
				bytes, errRead := efs.ReadFile(innerPath)
				if errRead != nil {
					return nil, errRead
				}
				item = newResourceAsFile(root, innerPath, bytes, efs, lang)
			} else {
				item = newResourceAsFile(root, innerPath, nil, efs, lang)
			}

		}

		res = append(res, item)
	}
	return res, err
}

func newResourceAsDir(root string, innerPath string, efs embed.FS) EmbedResource {
	fullPath := path.Join(root, innerPath)
	return &embedContentResource{
		isDir:        true,
		relativePath: innerPath,
		rootPath:     root,
		fullPath:     fullPath,

		mutex: &sync.RWMutex{},
		efs:   efs,
	}
}

func newResourceAsFile(root string, innerPath string, bytes []byte, efs embed.FS, lang string) EmbedResource {
	fullPath := path.Join(root, innerPath)
	baseName := filepath.Base(fullPath)

	if len(bytes) == 0 {
		return &embedContentResource{
			isDir:        false,
			relativePath: innerPath,
			lang:         lang,

			rootPath: root,
			fullPath: fullPath,

			mutex: &sync.RWMutex{},
			efs:   efs,

			fileName: baseName,
			raw:      nil,
			hasLoad:  false,
		}
	}

	return &embedContentResource{
		isDir:        false,
		relativePath: innerPath,
		lang:         lang,

		rootPath: root,
		fullPath: fullPath,

		mutex: &sync.RWMutex{},
		efs:   efs,

		fileName: baseName,
		hasLoad:  true,
		raw:      bytes,
	}
}
