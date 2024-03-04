/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/2 18:10
// @Desc 提供初始化全局slog函数

package log

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLog 初始化日志全局变量，默认不压缩文件和使用本地时间创建时间戳，日志支持切割
//
// PARAMS:
// - filename：日志文件名字
// - level：显示日志等级
// - maxSize：日志文件占用最大存储空间，单位MB
// - maxBackups：日志文件备份数量
// - maxAge：保留旧日志文件的最大天数
//
// RETURNS:
// -
func InitLog(filename string, level int, maxSize int, maxBackups int, maxAge int) {
	log := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  true,
		Compress:   false,
	}
	logger := slog.New(slog.NewJSONHandler(log, nil))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(logLevel(level))
	slog.Info("初始化日志成功")
}

const (
	LevelDebug int = -4
	LevelInfo  int = 0
	LevelWarn  int = 4
	LevelError int = 8
)

func logLevel(level int) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
