package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	mysqlInit = 0
	db        *gorm.DB
)
var (
	ErrInitialized    = errors.New("initialized")
	ErrNotInitialized = errors.New("not initialized")
)

// InitMysql 初始化mysql连接
func InitMysql(c MysqlConfig) error {
	if err := c.Check(); err != nil {
		return fmt.Errorf("config invalid: %w", err)

	}
	if mysqlInit != 0 {
		return ErrInitialized
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.DBName)

	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	db.SingularTable(true) // 默认不添加s

	mysqlInit = 1
	return nil
}

func getDB() (*gorm.DB, error) {
	if mysqlInit == 0 {
		return nil, ErrNotInitialized
	}
	return db, nil
}
