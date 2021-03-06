package global

// 定义一些设置相关的全局变量

import (
	"golo3/pkg/setting"
)

var (
	AppSetting      *setting.AppSettings
	FileSetting     *setting.FileSettings
	DatabaseSetting *setting.DatabaseSettings
	QywechatSetting *setting.QywechatSettings
	EmailSetting    *setting.EmailSettings
)
