/* Copyright 2024 follow. All Rights Reserved */
// @Author miaomiao
// @Date 2024/3/3 09:44
// @Desc 定义路由以及启动server的方法

package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/follow/controller"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	defaultLogFileNamePattern string = "log/service/service.log.%Y%m%d%H%M"
	defaultLogFile                   = "log/service/service.log"
	defaultLogMaxAge                 = 3 * 24 * time.Hour
	defaultRotationTime              = 1 * time.Hour
)

func initRouter() (*gin.Engine, error) {
	router := gin.Default()

	if err := setLogWriter(); err != nil {
		return nil, fmt.Errorf("set log writer failed: %v", err)
	}

	registerMiddleware(router)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "start")
	})

	// user
	router.POST("/userExist", controller.UserExist)
	router.POST("/addUser", controller.AddUser)

	// script
	router.POST("/getUserAllScript", controller.GetUserAllScript)
	router.POST("/addScript", controller.AddScript)

	return router, nil
}

// GraceFulStartServer 优雅的启动或重启server，需要注意：如果在重启的过程失败，会Panic
func GraceFulStartServer(addr string) error {
	router, err := initRouter()
	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen failed: %v", err)
		}
	}()
	slog.Info("server start")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shut down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shut down failed: %v", err)
	}
	return nil
}

func setLogWriter() error {
	writer, err := rotatelogs.New(
		defaultLogFileNamePattern,
		rotatelogs.WithLinkName(defaultLogFile),
		rotatelogs.WithMaxAge(defaultLogMaxAge),
		rotatelogs.WithRotationTime(defaultRotationTime),
	)
	if err != nil {
		return fmt.Errorf("new rotatelogs failed: %w", err)
	}
	gin.DefaultWriter = writer
	gin.DefaultErrorWriter = writer
	gin.DisableConsoleColor()
	return nil
}
