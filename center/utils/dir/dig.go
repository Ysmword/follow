/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/2 18:20
// @Desc 定义跟文件夹相关操作函数

package dir

import (
	"fmt"
	"os"
)

const (
	defaultCreateDirMode os.FileMode = 0777
)

// CreateDir 创建文件夹
func CreateDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return os.MkdirAll(dirPath, defaultCreateDirMode)
	}
	if !info.IsDir() {
		return fmt.Errorf("the [%s] already exist, but not a directory", dirPath)
	}
	return nil
}

// RemoveDir 删除文件夹
func RemoveDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil
	}

	if !info.IsDir() {
		return fmt.Errorf("[%s] is not a directory", dirPath)
	}

	return os.RemoveAll(dirPath)
}

// IsDir 返回是否是一个文件夹
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
