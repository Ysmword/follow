/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/17 10:07
// @Desc

package times

import "time"

// GetCurTimeInt 获取当前时间
func GetCurTimeInt() int64 {
	return time.Now().Unix()
}
