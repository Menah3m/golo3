package global

//定义一些数据库相关的全局变量

import "github.com/jinzhu/gorm"

var (
	DBEngine *gorm.DB
)
