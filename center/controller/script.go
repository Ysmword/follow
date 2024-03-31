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

	"github.com/follow/config"
	"github.com/follow/cron"
	"github.com/follow/model"
	"github.com/follow/utils/file"
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

// AuScript 添加或更新脚本
func AuScript(c *gin.Context) {
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

	if !s.Status {
		sfa, err := config.GetScriptFileAddr()
		if err != nil {
			response.FailWithReason(c, fmt.Sprintf("获取脚本地址失败：%v", err))
			return
		}
		if err := s.Build(sfa); err != nil {
			response.FailWithReason(c, fmt.Sprintf("编译失败%v", err))
			return
		}
		spec := fmt.Sprintf("*/%d * * * *", s.Cycle)
		entryID, err := cron.RegistrateTask(spec, s.RunShell, cron.SetTaskName(s.Username, s.Name))
		if err != nil {
			slog.Error(fmt.Sprintf("registrate task failed: %v", err))
			return
		}
		s.CronID = entryID
	} else {
		// 删除任务
		cron.RemoveTask(cron.SetTaskName(s.Username, s.Name))
	}

	if err := s.CreateOrUpdate(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("创建或更新脚本失败：%v", err))
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
	if err := file.WriteFile(tempFile, []byte(s.Code)); err != nil {
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

func DeleteScript(c *gin.Context) {
	s := model.Script{}
	if err := c.BindJSON(&s); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}

	if s.ID <= 0 {
		response.FailWithReason(c, "请输入脚本ID")
		return
	}

	if err := s.Delete(); err != nil {
		slog.Error(fmt.Sprintf("delete [%+v] failed:", s))
		response.FailWithReason(c, "删除失败")
		return
	}
	response.Success(c)
}

func GetScriptByID(c *gin.Context) {
	s := model.Script{}
	if err := c.BindJSON(&s); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}
	if s.ID <= 0 {
		response.FailWithReason(c, "请输入脚本ID")
		return
	}

	if err := s.GetByID(); err != nil {
		slog.Error(fmt.Sprintf("根据ID获取脚本信息失败：%v", err))
		response.FailWithReason(c, bindJsonFail(err))
		return
	}
	response.SuccessWithData(c, s)
}
