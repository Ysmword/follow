package model

import (
	"fmt"
	"reflect"
)

// Script 脚本
type Script struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Username    string `json:"username"`                    // 用户名
	Name        string `json:"name" gorm:"unique"`          // 脚本名称
	Type        string `json:"type" gorm:"not null"`        // 脚本类型
	Cycle       int    `json:"cycle"`                       // 运行周期，默认一个小时，单位s，数据库默认为6min
	Status      bool   `json:"status"`                      // 脚本状态，true为开启，false为关闭，默认不开启
	CreateTime  int64  `json:"create_time"`                 // 创建时间
	UpdateTime  int64  `json:"update_time"`                 // 更新时间
	Description string `json:"description" gorm:"not null"` // 描述
}

func (s *Script) Check() error {
	elem := reflect.ValueOf(s).Elem()
	elemType := elem.Type()
	for index := 0; index < elem.NumField(); index++ {
		field, fieldName := elem.Field(index), elemType.Field(index).Name
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("field [%s] expect", fieldName)
			}
		default:
			// do nothing
		}
	}
	return nil
}

// GetByID 根据ID获取脚本
func (s *Script) GetByID() (script Script, err error) {
	return script, err
}

// GetByUsername 获取用户名下的脚本
func (s *Script) GetByUsername() ([]Script, error) {
	cursor, err := getDB()
	if err != nil {
		return []Script{}, err
	}
	scripts := make([]Script, 0)
	if cursor.Where("username=?", s.Username).Find(&scripts).Error != nil {
		return []Script{}, fmt.Errorf("find failed: %v", err)
	}
	return scripts, nil
}

// Delete 根据ID删除脚本
func (s *Script) Delete() error {
	return nil
}

// Create 创建脚本
func (s *Script) Create() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}

	return cursor.Create(s).Error
}

// Update 更新脚本
func (s *Script) Update() error {
	return nil
}
