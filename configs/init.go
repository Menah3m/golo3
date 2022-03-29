package configs

import (
	"github.com/fsnotify/fsnotify"
	"golo3/global"
	"golo3/model"
	"golo3/pkg/setting"
	"time"
)

func SetupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = ReadSection(newSetting)
	if err != nil {
		return err
	}
	//监听配置文件变化，如果发生变化，则重新读取配置
	newSetting.Vp.WatchConfig()
	newSetting.Vp.OnConfigChange(func(e fsnotify.Event) {
		err := ReadSection(newSetting)
		if err != nil {
			return
		}
	})
	return nil
}

func ReadSection(newSetting *setting.Setting) error {
	err := newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("File", &global.FileSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Qywechat", &global.QywechatSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.Duration *= time.Minute
	return nil
}

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
