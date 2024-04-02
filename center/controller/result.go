package controller

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/follow/model"
	"github.com/follow/utils/response"
)

func GetAllResult(c *gin.Context) {
	r := model.Result{}
	results, err := r.GetAll()
	if err != nil {
		slog.Error(fmt.Sprintf("get all result failed: %v", err))
		response.FailWithReason(c, "获取所有的结果失败")
		return
	}
	response.SuccessWithData(c, results)
}

func GetResultByU(c *gin.Context) {
	r := model.Result{}
	if err := c.BindJSON(&r); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}

	if r.Username == "" {
		response.FailWithReason(c, "用户名必填")
		return
	}
	results, err := r.GetByUsername()
	if err != nil {
		slog.Error(fmt.Sprintf("get result by username [%s] failed: %v", r.Username, err))
		response.FailWithReason(c, "根据用户名获取失败")
		return
	}
	response.SuccessWithData(c, results)
}

func GetResultByUT(c *gin.Context) {
	r := model.Result{}
	if err := c.BindJSON(&r); err != nil {
		response.FailWithReason(c, bindJsonFail(err))
		return
	}
	if r.Username == "" {
		response.FailWithReason(c, "用户名必填")
		return
	}
	if r.Type == "" {
		response.FailWithReason(c, "类型必填")
		return
	}

	results, err := r.GetByUT()
	if err != nil {
		slog.Error(fmt.Sprintf("get result by username [%s] and type [%s] failed: %v", r.Username, r.Type, err))
		response.FailWithReason(c, "根据用户名和类型获取失败")
		return
	}
	response.SuccessWithData(c, results)
}
