package store

import (
	"github.com/jinzhu/gorm"
	"golo3/model"
)

/*
负责 将日志信息存储到数据库
  -- Mysql
*/

//WriteLogInfo
func WriteLogInfo(db *gorm.DB, info *model.LogInfo) error {
	err := db.LogMode(false).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.LogInfo{}).Error
	if err != nil {
		return err
	}
	err = db.Create(info).Error
	if err != nil {
		return err
	}
	return nil
}
