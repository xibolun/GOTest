package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 获取文件列表
// 给定一个目录，递归创建文件
// 给定一个目录，递归写入文件
// 获取一个文件path的父目录

// GetParentPath 获取上级目录
func GetParentPath(path string) string {
	return filepath.Dir(path)
}

// CreateFileRecursive 递归创建文件
func CreateFileRecursive(path string) error {
	if IsExist(path) {
		return os.ErrExist
	}
	pPath := filepath.Dir(path)
	if !IsExist(pPath) {
		if err := os.MkdirAll(pPath, 0755); err != nil {
			return fmt.Errorf("mkdir %s fail,%s", pPath, err.Error())
		}
	}

	if _, err := os.Create(path); err != nil {
		return err
	}

	return nil
}

// GetFileList 获取目录下的文件列表
func GetFileList(dir, excludes string) []*os.File {
	var files []*os.File

	root, err := os.Open(dir)
	if err != nil {
		return files
	}
	fi, err := root.Stat()
	if err != nil {
		return files
	}
	if !fi.IsDir() {
		files = append(files, root)
		return files
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return files
	}

	for _, item := range fis {
		if strings.Index(excludes, item.Name()) != -1 {
			continue
		}
		fileFullPath := filepath.Join(dir, item.Name())

		if !item.IsDir() {
			f, _ := os.Open(fileFullPath)
			files = append(files, f)
			continue
		}
		files = append(files, GetFileList(fileFullPath, excludes)...)
	}
	return files
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	return os.IsExist(err) && err != nil
}
