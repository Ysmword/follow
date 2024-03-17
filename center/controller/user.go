/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/16 17:59
// @Desc

package controller

import (
	"fmt"
	"log/slog"

	"github.com/follow/model"
	"github.com/follow/utils/response"
	"github.com/follow/utils/times"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func bindJsonFail(err error) string {
	return fmt.Sprintf("数据解析失败：%v", err)
}

// UserExist 判断用户是否存在
func UserExist(c *gin.Context) {
	u := model.User{}
	if err := c.BindJSON(&u); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}

	err := u.Get()
	if gorm.IsRecordNotFoundError(err) {
		response.FailWithReason(c, "用户名或密码不正确")
		return
	}

	if err != nil {
		slog.Error(fmt.Sprintf("查询用户失败: %v", err))
		response.FailWithReason(c, "系统异常")
		return
	}

	response.Success(c)
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	u := model.User{}
	if err := c.BindJSON(&u); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}
	u.CreateTime = times.GetCurTimeInt()
	if err := u.Check(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("用户信息检查失败：%v", err))
		return
	}

	if err := u.Create(); err != nil {
		response.FailWithReason(c, fmt.Sprintf("创建用户失败：%v", err))
		return
	}

	response.SuccessWithData(c, "创建用户成功")
}
