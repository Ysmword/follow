/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/16 17:31
// @Desc

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusCodeSuccess = iota
	StatusCodeFail
)

const (
	DefSuccessFlag = "success"
)

func commonResp(status int, msg string, data interface{}) gin.H {
	return gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	}
}

// Success 成功响应，不会携带数据
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, commonResp(StatusCodeSuccess, DefSuccessFlag, nil))
}

// SuccessWithData 成功响应，会携带数据
func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, commonResp(StatusCodeSuccess, DefSuccessFlag, data))
}

// FailWithReason 失败响应
func FailWithReason(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, commonResp(StatusCodeFail, msg, nil))
}
