package utils

import (
	"os"
	"path/filepath"
)

// GetDirPath는 환경변수의 경로 or Getwd()/dirname 디렉토리 경로를 반환한다.
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

// GetFilePath는 GetDirPath()/filename의 파일 경로를 반환한다.
func GetFilePath(dirname string, filename string) (string, error) {
	msgDir, err := GetDirPath(dirname)
	if err != nil {
		return "", err
	}

	return filepath.Join(msgDir, filename), nil 
}

// IsDirExisting는 디렉토리가 존재하는지에 대한 여부를 bool로 반환한다.
// 존재하면 true, 아니면 false
func IsDirExisting(dirname string) (bool, error) {
	dirPath, err := GetDirPath(dirname)
	if err != nil {
		return false, err
	}

	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	
	return info.IsDir(), nil
}