package model

// Script 脚本
type Script struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Username    string `json:"username"`                    // 用户名
	Name        string `json:"name" gorm:"unique"`          // 脚本名称
	Type        string `json:"type" gorm:"not null"`        // 脚本类型
	Cycle       int    `json:"cycle"`                       // 运行周期，默认一个小时，单位s
	Status      bool   `json:"status"`                      // 脚本状态，true为开启，false为关闭
	CreateTime  int64  `json:"create_time"`                 // 创建时间
	UpdateTime  int64  `json:"update_time"`                 // 更新时间
	Description string `json:"description" gorm:"not null"` // 描述
}

// GetByID 根据ID获取脚本
func (s *Script) GetByID(id string) (script Script, err error) {
	return script, err
}

// GetByUsername 获取用户名下的脚本
func (s *Script) GetByUsername(username string) ([]Script, error) {
	return []Script{}, nil
}

// Delete 根据ID删除脚本
func (s *Script) Delete(id string) error {
	return nil
}

// Create 创建脚本
func (s *Script) Create(s Script) error {
	return nil
}

// Update 更新脚本
func (s *Script) Update(s Script) error {
	return nil
}
