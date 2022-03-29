package main

import (
	"golo3/configs"
	"golo3/internal/parse"
	"log"
)

// 初始化工作，包括 读取设置， 连接数据库..等等
func init() {
	err := configs.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	err = configs.SetupDBEngine()
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
