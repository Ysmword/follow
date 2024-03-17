package model

// CrawlerRes 爬虫结果
type CrawlerRes struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Type        string `json:"type"`        // 关联类型
	ScriptName  string `json:"script_name"` // 关联脚本名称
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreateTime  int64  `json:"create_time"`
}

// Get 分页获取爬虫结果
func (c *CrawlerRes) Get(page, pageSize int64) ([]CrawlerRes, int64) {
	return nil, 0
}

// DeleteByID 批量根据ID删除
func (c *CrawlerRes) DeleteByID(ids []int64) error {
	return nil
}

// DeleteExpire 删除过期数据
func (c *CrawlerRes) DeleteExpire() error {
	return nil
}
