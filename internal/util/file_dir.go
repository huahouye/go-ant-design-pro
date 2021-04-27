package util

import (
	"h-trader/internal/logger"
	"os"
	"path/filepath"
)

func EnsureDir(dirName string) bool {

	err := os.Mkdir(dirName, os.ModePerm)
	if err == nil || os.IsExist(err) {
		return true
	} else {
		return false
	}
}

func CreateFileWithDirs(filename string) (*os.File, error) {

	absFilename, err := filepath.Abs(filename)
	logger.Infof("absFilename: %s", absFilename)
	if err == nil {
		if err := os.MkdirAll(filepath.Dir(absFilename), os.ModePerm); err != nil {
			return nil, err
		}
		return os.Create(absFilename)
	}

	return nil, err
}
