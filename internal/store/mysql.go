package store

import (
	"github.com/jinzhu/gorm"
	"golo3/model"
)

/*
负责 将日志信息存储到数据库
  -- Mysql
*/

//WriteRecord
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

func WriteCountInfo(db *gorm.DB, countInfo *model.CountInfo) error {
	err := db.LogMode(false).Error
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.CountInfo{}).Error
	if err != nil {
		return err
	}
	err = db.Create(countInfo).Error
	if err != nil {
		return err
	}
	return nil
}
