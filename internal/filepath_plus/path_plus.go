package filepath_plus

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// PathExists
//
//	path exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathExistsFast
//
//	path exists fast
func PathExistsFast(path string) bool {
	exists, _ := PathExists(path)
	return exists
}

// PathIsDir
//
//	path is dir
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// RmDir
//
//	remove dir by path
//
//nolint:golint,unused
func RmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// Mkdir
//
//	will use FileMode 0766
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// ReadFileAsByte
//
//	read file as byte array
func ReadFileAsByte(path string) ([]byte, error) {
	exists, err := PathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	if PathIsDir(path) {
		return nil, fmt.Errorf("path is dir: %s", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// WriteFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0766) os.FileMode(0666) os.FileMode(0644)
//	coverage true will coverage old
func WriteFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !PathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}
