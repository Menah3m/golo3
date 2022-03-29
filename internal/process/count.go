package process

import (
	"github.com/jinzhu/gorm"
	"golo3/model"
)

/*
负责 处理日志信息
  -- 记数
*/
func Count(db *gorm.DB, keyword string, start, end string) int {
	var count int
	db.Model(&model.LogInfo{}).Where("log_keyword = ? AND start_at BETWEEN  ? AND ? ", keyword, start, end).Count(&count)
	return count
}
