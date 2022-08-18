package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
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

func WriteFileByPath(path string, str string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}
	return WriteFileByObj(file, str)
}

func WriteFileByObj(file *os.File, str string) error {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	md5Byte := md5.Sum([]byte(str))
	_, err := file.WriteString(str)
	if err != nil {
		return err
	}
	split := strings.Split(file.Name(), "/")
	versionFileName := "version/" + split[len(split)-1]
	versionExist := FileIsExist(versionFileName)
	if !versionExist {
		versionFile, err := os.Create(versionFileName)
		defer func(versionFile *os.File) {
			err := versionFile.Close()
			if err != nil {
				return
			}
		}(versionFile)
		if err != nil {
			return err
		}
		_, err = versionFile.WriteString(fmt.Sprintf("%x", md5Byte))
		if err != nil {
			return err
		}
	} else {
		versionFile, err := os.OpenFile(versionFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
		defer func(versionFile *os.File) {
			err := versionFile.Close()
			if err != nil {
			}
		}(versionFile)
		if err != nil {
			return err
		}
		_, err = io.WriteString(versionFile, fmt.Sprintf("%x", md5Byte))
		if err != nil {
			return err
		}
	}
	return nil
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
