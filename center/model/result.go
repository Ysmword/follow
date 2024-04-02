package model

import "fmt"

// Result 爬虫结果
type Result struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	Type       string `json:"type"`
	ScriptName string `json:"script_name"`
	Username   string `json:"userame"` // 用户名
	Link       string `json:"link"`    // 请求连接
	Header     string `json:"header"`  // 标题
	Image      string `json:"image"`   // 图像连接
	Content    string `json:"content"` // 内容
	CreateTime int64  `json:"create_time"`
}

func (r *Result) GetAll() ([]Result, error) {
	cursor, err := getDB()
	if err != nil {
		return nil, err
	}
	all := make([]Result, 0)
	if err := cursor.Find(&all).Error; err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	return all, nil
}

func (r *Result) GetByUsername() ([]Result, error) {
	cursor, err := getDB()
	if err != nil {
		return nil, err
	}
	all := make([]Result, 0)
	if err := cursor.Where("username=?", r.Username).Find(&all).Error; err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	return all, nil
}

func (r *Result) GetByUT() ([]Result, error) {
	cursor, err := getDB()
	if err != nil {
		return nil, err
	}
	all := make([]Result, 0)
	if err := cursor.Where("username=? and type=?", r.Username, r.Type).Find(&all).Error; err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	return all, nil
}

func (r *Result) Create() error {
	cursor, err := getDB()
	if err != nil {
		return err
	}
	return cursor.Create(r).Error
}
