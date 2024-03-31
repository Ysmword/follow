/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/2 18:57
// @Desc 定义配置文件操作函数

package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/follow/model"
)

var (
	ErrNotInitialized = errors.New("not initialized")
)

var (
	appConf *Config
)

type Config struct {
	Log        LogConfig
	Server     ServerConfig
	Mysql      model.MysqlConfig
	ScriptFile model.ScriptFileAddr `toml:"script_file"`
}

type LogConfig struct {
	FileName   string `toml:"fileName"`
	MaxSize    int    `toml:"maxSize"`
	MaxAge     int    `toml:"maxAge"`
	MaxBackups int    `toml:"maxBackups"`
	Level      int    `toml:"level"`
}

type ServerConfig struct {
	Addr string `toml:"addr"`
}

func LoadConfig(filePath string) error {
	var c Config
	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		return fmt.Errorf("decode file [%s] failed: %w", filePath, err)
	}

	if err := c.Log.check(); err != nil {
		return fmt.Errorf("check log config failed: %w", err)
	}

	if err := c.Server.check(); err != nil {
		return fmt.Errorf("check server config failed: %w", err)
	}

	if err := c.Mysql.Check(); err != nil {
		return fmt.Errorf("check mysql config failed: %w", err)
	}

	if err := c.ScriptFile.Check(); err != nil {
		return fmt.Errorf("check script file addr failed: %w", err)
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

func (s *ServerConfig) check() error {
	if s.Addr == "" {
		return fmt.Errorf("addr [%s] expect", s.Addr)
	}
	return nil
}

func GetLogConfig() (LogConfig, error) {
	if appConf == nil {
		return LogConfig{}, ErrNotInitialized
	}
	return appConf.Log, nil
}

func GetServerConfig() (ServerConfig, error) {
	if appConf == nil {
		return ServerConfig{}, ErrNotInitialized
	}
	return appConf.Server, nil
}

func GetMysqlConfig() (model.MysqlConfig, error) {
	if appConf == nil {
		return model.MysqlConfig{}, ErrNotInitialized
	}
	return appConf.Mysql, nil
}

func GetScriptFileAddr() (model.ScriptFileAddr, error) {
	if appConf == nil {
		return model.ScriptFileAddr{}, ErrNotInitialized
	}

	return appConf.ScriptFile, nil
}
