package utils

import (
	"os"
	"path/filepath"
)

func GetDirPath(dirname string) (string, error) {
	msgDir := os.Getenv("APP_MSG_DIR")
	if msgDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		msgDir = filepath.Join(wd, "message", dirname)
	} else {
		msgDir = filepath.Join(msgDir, dirname)
	}

	return msgDir, nil
}

func GetFilePath(dirname string, filename string) (string, error) {
	msgDir, err := GetDirPath(dirname)
	if err != nil {
		return "", err
	}

	return filepath.Join(msgDir, filename), nil 
}