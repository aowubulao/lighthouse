package utils

import (
	"os"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetDirFiles(path string) (string, error) {
	rd, err := os.ReadDir(path)
	resStr := ""
	if err != nil {
		return resStr, err
	}
	for _, file := range rd {
		if !file.IsDir() {
			resStr += file.Name()
		}
	}
	return resStr, nil
}
