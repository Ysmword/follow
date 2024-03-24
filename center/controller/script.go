/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/17 10:43
// @Desc

package controller

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/follow/model"
	"github.com/follow/utils/response"
	"github.com/follow/utils/times"
)

// GetUserAllScript 获取用户的所有脚本
func GetUserAllScript(c *gin.Context) {
	s := model.Script{}
	if err := c.BindJSON(&s); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}

	if s.Username == "" {
		response.FailWithReason(c, "用户名必填")
		return
	}

	scripts, err := s.GetByUsername()
	if gorm.IsRecordNotFoundError(err) || err == nil {
		response.SuccessWithData(c, scripts)
		return
	}

	response.FailWithReason(c, fmt.Sprintf("系统bug: %v", err))
}

// AddScript 添加脚本
func AddScript(c *gin.Context) {
	s := model.Script{}
	if err := c.BindJSON(&s); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}
	s.CreateTime = times.GetCurTimeInt()
	s.UpdateTime = times.GetCurTimeInt()
	if err := s.Check(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("脚本检查失败：%v", err))
		return
	}

	exist, err := s.Exist()
	if err != nil || exist {
		response.FailWithReason(c, fmt.Sprintf("查询脚本失败：%v 或者 脚本已经存在：%v", err, exist))
		return
	}

	if err := s.Create(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("创建脚本失败：%v", err))
		return
	}

	response.SuccessWithData(c, "创建脚本成功")
}

// RunDebug 脚本试运行,这里带优化
func RunDebug(c *gin.Context) {
	s := model.Script{}
	if err := c.BindJSON(&s); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}

	tempFile := fmt.Sprintf("%s_temp_%d.go", s.Username, times.GetCurTimeInt())
	if err := writeFile(tempFile, []byte(s.Code)); err != nil {
		slog.Error(fmt.Sprintf("创建文件失败：%v", err))
		response.FailWithReason(c, "系统出现bug，请联系RD进行处理")
		return
	}

	cmd := exec.Command("go", "run", tempFile)
	ouput, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(fmt.Sprintf("运行脚本失败：%v", err))
		response.FailWithReason(c, fmt.Sprintf("运行脚本失败：%v", err))
		return
	}
	response.SuccessWithData(c, string(ouput))

	if err := os.Remove(tempFile); err != nil {
		slog.Error(fmt.Sprintf("删除脚本失败：%v", err))
	}
}

func writeFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
