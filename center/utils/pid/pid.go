/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/2 18:14
// @Desc 定义pid文件相关操作

package pid

import (
	"errors"
	"fmt"
	"github.com/follow/utils/dir"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	NotExistCode int = -1
)

var (
	ErrPidAlreadyExist = errors.New("pid exist already")
)

// CheckPidExist 检查pid文件是否存在
//
// PARAMS:
// - dirName: 文件夹名字
// - fileName: 文件名字
//
// RETURNS:
// - int:：pid，如果有错误，返回-1
// - error：错误信息
func CheckPidExist(dirName, fileName string) (int, error) {
	// 读取文件
	pidFileName := path.Join(dirName, fileName)
	idByte, err := os.ReadFile(pidFileName)
	if err != nil {
		return NotExistCode, fmt.Errorf("read pid from file [%s] failed: %w", pidFileName, err)
	}

	// 文件内容转化
	idStr := strings.TrimSpace(string(idByte))
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		return NotExistCode, fmt.Errorf("expect number format pid: %s", idStr)
	}
	return idNum, nil
}

// CreateNewPid 创建新的pid文件
//
// PARAMS:
// - dirName: 文件夹名字
// - fileName: 文件名字
//
// RETURNS:
// - int:：pid，如果有错误，返回-1
// - error：错误信息
func CreateNewPid(dirName, fileName string) (int, error) {
	// 创建文件夹
	if err := dir.CreateDir(dirName); err != nil {
		return NotExistCode, fmt.Errorf("create directory [%s] failed: %w", dirName, err)
	}

	// 将文件写入到文件
	idNum := os.Getpid()
	pidFileName := path.Join(dirName, fileName)
	if err := os.WriteFile(pidFileName, []byte(fmt.Sprintf("%d", idNum)), 0644); err != nil {
		return idNum, fmt.Errorf("write pid [%d] into file [%s] failed: %w", idNum, pidFileName, err)
	}
	return idNum, nil
}

// RemovePid 删除pid文件
//
// PARAMS:
// - dirName: 文件夹名字
// - fileName: 文件名字
//
// RETURNS:
// - error：错误信息
func RemovePid(dirName, fileName string) error {
	if _, err := CheckPidExist(dirName, fileName); err != nil {
		return err
	}

	pidFileName := path.Join(dirName, fileName)
	if err := os.Remove(pidFileName); err != nil {
		return fmt.Errorf("remove pid file [%s] failed: %w", pidFileName, err)
	}
	return nil
}
