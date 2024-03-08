package model

import (
	"fmt"
	"reflect"
)

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DBName   string `toml:"dbName"`
	RunMode  string `toml:"runMode"`
}

// Check 检查配置
func (m *MysqlConfig) Check() error {
	elem := reflect.ValueOf(m).Elem()
	elemType := elem.Type()
	for index := 0; index < elem.NumField(); index++ {
		field, fieldName := elem.Field(index), elemType.Field(index).Name
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("field [%s] expect", fieldName)
			}
		default:
			// do nothing.
		}
	}
	return nil
}
