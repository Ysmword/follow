/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/17 10:43
// @Desc

package controller

import (
	"fmt"

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

	if err := s.Create(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("创建脚本失败：%v", err))
		return
	}

	response.SuccessWithData(c, "创建脚本成功")
}
