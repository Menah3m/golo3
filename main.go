package main

import (
	"golo3/global"
	"golo3/internal/parse"
	"golo3/model"
	"golo3/pkg/setting"
	"log"
	"time"
)

// 初始化工作，包括 读取设置， 连接数据库..等等
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err :%v", err)
	}

}

func main() {

	err := parse.ReadLine()
	//_, err := alert.QywechatAlert()
	if err != nil {
		log.Println(err)
	}

}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
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

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
