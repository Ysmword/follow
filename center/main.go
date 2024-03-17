package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/follow/config"
	"github.com/follow/model"
	"github.com/follow/router"
	"github.com/follow/utils/log"
	"github.com/follow/utils/pid"
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

	// 	初始化log
	lc, err := config.GetLogConfig()
	if err != nil {
		fmt.Println("get log config failed:", err)
		return
	}
	log.InitLog(lc.FileName, lc.Level, lc.MaxSize, lc.MaxBackups, lc.MaxAge)

	// 初始化mysql数据库
	mc, err := config.GetMysqlConfig()
	if err != nil {
		fmt.Println("get mysql config failed:", err)
		return
	}
	if err := model.InitMysql(mc); err != nil {
		fmt.Println("init mysql failed:", err)
		return
	}

	// 初始化server
	sc, err := config.GetServerConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("get server failed: %v", err))
		return
	}
	if err := router.GraceFulStartServer(sc.Addr); err != nil {
		slog.Error(fmt.Sprintf("graceful start server failed: %v", err))
		return
	}
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
		return fmt.Errorf("pid [%d] already exist", id)
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
