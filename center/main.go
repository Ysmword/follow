package main

import (
	"errors"
	"fmt"
	"github.com/follow/config"
	"github.com/follow/utils/log"
	"github.com/follow/utils/pid"
	"log/slog"
	"os"
	"path"
)

// pid
const (
	defaultPidDir  = "var"
	defaultPidFile = "follow.pid"
)

const (
	defaultConfFile = "conf/app.toml"
)

func main() {
	if err := createPid(); err != nil {
		fmt.Println("create new pid failed:", err)
		return
	}
	defer removePid()

	if err := config.LoadConfig(defaultConfFile); err != nil {
		fmt.Println("load app config failed:", err)
		return
	}

	lc, err := config.GetLogConfig()
	if err != nil {
		fmt.Println("get log config failed:", err)
	}
	log.InitLog(lc.FileName, lc.Level, lc.MaxSize, lc.MaxBackups, lc.MaxAge)

	slog.Info("follow启动")
}

func createPid() error {
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get work directory failed: %w", err)
	}
	id, err := pid.CreateNewPid(path.Join(workDir, defaultPidDir), defaultPidFile)
	if err == nil && id != pid.NotExistCode {
		fmt.Println("create new pid success")
		return nil
	}

	if errors.Is(err, pid.ErrPidAlreadyExist) {
		return fmt.Errorf("pid [%s] already exist", id)
	}
	return err
}

func removePid() {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println("get work directory failed:", err)
		return
	}
	if err = pid.RemovePid(path.Join(workDir, defaultPidDir), defaultPidFile); err != nil {
		fmt.Println("remove pid failed:", err)
		return
	}
	fmt.Println("remove pid successfully")
}
