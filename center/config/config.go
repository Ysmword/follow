/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/2 18:57
// @Desc 定义配置文件操作函数

package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	ErrNotInitialized = errors.New("not initialized")
)

var (
	appConf *Config
)

type Config struct {
	Log LogConfig
}

type LogConfig struct {
	FileName   string `toml:"fileName"`
	MaxSize    int    `toml:"maxSize"`
	MaxAge     int    `toml:"maxAge"`
	MaxBackups int    `toml:"maxBackups"`
	Level      int    `toml:"level"`
}

func LoadConfig(filePath string) error {
	var c Config
	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		return fmt.Errorf("decode file [%s] failed: %w", filePath, err)
	}

	if err := c.Log.check(); err != nil {
		return fmt.Errorf("check log config failed: %w", err)
	}

	appConf = &c
	return nil
}

func (l *LogConfig) check() error {
	if l.FileName == "" {
		return fmt.Errorf("fileName [%s] expect", l.FileName)
	}
	if l.MaxSize <= 0 {
		return fmt.Errorf("maxSize [%d] expect > 0", l.MaxSize)
	}
	if l.MaxAge <= 0 {
		return fmt.Errorf("maxAge [%d] expect > 0", l.MaxAge)
	}
	if l.MaxBackups <= 0 {
		return fmt.Errorf("maxBackups [%d] expect > 0", l.MaxBackups)
	}
	return nil
}

func GetLogConfig() (LogConfig, error) {
	if appConf == nil {
		return LogConfig{}, ErrNotInitialized
	}
	return appConf.Log, nil
}
