package filepath_plus

import (
	"fmt"
	"io/fs"
	"os"
)

// AppendFileHead
//
//	append file head
//	@param path string
//	@param content []byte
//
// this method will keep file mode
// if file not exist, will not create file and return error
func AppendFileHead(path string, content []byte) error {
	if len(content) == 0 {
		return fmt.Errorf("AppendFileHead content is empty")
	}
	fileAsByte, err := ReadFileAsByte(path)
	if err != nil {
		return err
	}
	fileAsByte = append([]byte(content), fileAsByte...)
	return AlterFile(path, fileAsByte)
}

// AppendFileTail
//
//	append file tail
//	@param path string
//	@param content []byte
//
// this method will keep file mode
// if file not exist, will not create file and return error
func AppendFileTail(path string, content []byte) error {
	if len(content) == 0 {
		return fmt.Errorf("AppendFileTail content is empty")
	}
	fileAsByte, err := ReadFileAsByte(path)
	if err != nil {
		return err
	}
	fileAsByte = append(fileAsByte, content...)
	return AlterFile(path, fileAsByte)
}

// AlterFile
//
//	alter file
//	@param path string
//	@param data []byte
//
// this method will keep file mode
// if file not exist, will not create file and return error
func AlterFile(path string, data []byte) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("AlterFile get file info at path: %v, err: %v", path, err)
	}
	err = os.WriteFile(path, data, fileInfo.Mode())
	if err != nil {
		return fmt.Errorf("cAlterFile write data at path: %v, err: %v", path, err)
	}
	return nil
}

// CheckOrCreateFileWithStringFast
//
//	check or create file with content, if file exist, will not coverage
//	crate folder of file with os.FileMode(0766)
//	@param path string
//	@param content string
func CheckOrCreateFileWithStringFast(path string, context string) error {
	return CheckOrCreateFileWithString(path, context, os.FileMode(0766))
}

// CheckOrCreateFileWithString
//
//	check or create file with content, if file exist, will not coverage
//	@param path string
//	@param content string
//	@param fileMod os.FileMode(0766) os.FileMode(0666) os.FileMode(0644)
func CheckOrCreateFileWithString(path string, content string, fileMod fs.FileMode) error {
	if content == "" {
		return fmt.Errorf("CheckOrCreateFileWithString content is empty")
	}
	if PathExistsFast(path) {
		return nil
	}
	return WriteFileByByte(path, []byte(content), fileMod, true)
}

// AlterFileWithStringFast
//
//	check or file with content, if file exist, will coverage
//	crate folder of file with os.FileMode(0766)
//	@param path string
//	@param content string
func AlterFileWithStringFast(path string, content string) error {
	return AlterFileWithString(path, content, os.FileMode(0766))
}

// AlterFileWithString
//
//	check or file with content, if file exist, will coverage
//	@param path string
//	@param content string
//	@param fileMod os.FileMode(0766) os.FileMode(0666) os.FileMode(0644)
func AlterFileWithString(path string, content string, fileMod fs.FileMode) error {
	if content == "" {
		return fmt.Errorf("CheckOrCreateFileWithString content is empty")
	}
	return WriteFileByByte(path, []byte(content), fileMod, true)
}
