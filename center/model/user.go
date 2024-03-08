package model

// User 用户，可以研究下如何接入第三方登录，比如微信，qq等
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`    // 存储密码是加密的
	Email      string `json:"email"`       // 邮箱，可以为空
	Phone      string `json:"phone"`       // 手机号，可以为空
	Status     int    `json:"status"`      // 状态，0：正常，1：禁用
	IsAdmin    int    `json:"is_admin"`    // 是否是管理员，0：不是，1：是
	CreateTime int64  `json:"create_time"` // 创建时间
}

// Get 获取用户信息
func (u *User) Get(username string, password string) (User, error) {
	return User{}, nil
}

// Create 创建用户
func (u *User) Create(user User) error {
	return nil
}

// Delete 删除用户
func (u *User) Delete(usernamet string) error {
	return nil
}

// Update 更新用户，全量更新了
func (u *User) Update() error {
	return nil
}
