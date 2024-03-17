package model

import (
	"fmt"
	"reflect"
)

// User 用户，可以研究下如何接入第三方登录，比如微信，qq等
type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username" gorm:"not null;unique"` // 这个要唯一
	Password   string `json:"password" gorm:"not null"`        // 存储密码是加密的
	Email      string `json:"email" gorm:"not null"`           // 邮箱，可以为空
	Phone      string `json:"phone" gorm:"not null"`           // 手机号，可以为空
	Status     int    `json:"status" gorm:"not null"`          // 状态，0：正常，1：禁用
	IsAdmin    int    `json:"is_admin" gorm:"not null"`        // 是否是管理员，0：不是，1：是
	CreateTime int64  `json:"create_time" gorm:"not null"`     // 创建时间
}

func (u *User) Check() error {
	elem := reflect.ValueOf(u).Elem()
	elemType := elem.Type()
	for index := 0; index < elem.NumField(); index++ {
		field, fieldName := elem.Field(index), elemType.Field(index).Name
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("field [%s] expect", fieldName)
			}
		case reflect.Int64:
			if fieldName == "ID" {
				continue
			}
			if field.Int() <= 0 {
				return fmt.Errorf("field [%s] expect > 0", fieldName)
			}
		default:
			// do nothing.
		}
	}
	return nil
}

// Get 获取用户信息
func (u *User) Get() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	return cursor.Where("username=? and password=?", u.Username, u.Password).First(u).Error
}

// Create 创建用户
func (u *User) Create() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	return cursor.Create(u).Error
}

// Delete 删除用户
func (u *User) Delete(username string) error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	return cursor.Where("username=?", u.Username).Delete(u).Error
}

// Update 更新用户，全量更新了
func (u *User) Update() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	return cursor.Model(u).Updates(u).Error
}
