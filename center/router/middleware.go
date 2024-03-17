/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/3 10:16
// @Desc 定义中间件

package router

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/follow/utils/response"
	"github.com/gin-gonic/gin"
)

// 硬编码权限验证
const (
	username = "follow"
	password = "follow@123456"
)

func registerMiddleware(r *gin.Engine) {
	slog.Info("start register middleware")
	defer slog.Info("register middleware end")

	r.Use(setLogFormat())
	r.Use(auth)
	r.Use(gin.Recovery())
}

func setLogFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.DateTime),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	})
}

func auth(c *gin.Context) {
	u := c.Query("username")
	p := c.Query("pwd")
	if u != username || p != password {
		c.Abort()
		response.FailWithReason(c, "权限验证失败")
	}
	c.Next()
}
