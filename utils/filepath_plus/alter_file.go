package filepath_plus

import (
	"fmt"
	"os"
)

func AppendFileHead(path string, content []byte) error {
	if content == nil || len(content) == 0 {
		return fmt.Errorf("AppendFileHead content is empty")
	}
	fileAsByte, err := ReadFileAsByte(path)
	if err != nil {
		return err
	}
	fileAsByte = append([]byte(content), fileAsByte...)
	return AlterFile(path, fileAsByte)
}

func AppendFileTail(path string, content []byte) error {
	if content == nil || len(content) == 0 {
		return fmt.Errorf("AppendFileTail content is empty")
	}
	fileAsByte, err := ReadFileAsByte(path)
	if err != nil {
		return err
	}
	fileAsByte = append(fileAsByte, content...)
	return AlterFile(path, fileAsByte)
}

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
