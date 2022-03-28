package parse

/*
负责 读取日志文件内容并解析
*/

import (
	"fmt"
	"github.com/hpcloud/tail"
	"golo3/global"
	"golo3/internal/alert"
	"golo3/internal/process"
	"golo3/internal/store"
	"golo3/model"
	"golo3/pkg/setting"
	"log"
	"strings"
	"time"
)

type Logfile struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
}

//bindLogfileSetting 绑定日志文件的设置信息
func bindLogfileSetting(fileSetting *setting.FileSettings) *Logfile {
	return &Logfile{
		LogSavePath: fileSetting.LogSavePath,
		LogFileName: fileSetting.LogFileName,
		LogFileExt:  fileSetting.LogFileExt,
	}
}

//getLogfilePath 获取日志文件完整路径
func getLogfilePath() string {
	fs := bindLogfileSetting(global.FileSetting)
	filepath := fs.LogSavePath + fs.LogFileName + fs.LogFileExt
	return filepath
}

//ReadLine 按行读取日志文件内容
func ReadLine() error {
	fp := getLogfilePath()
	fmt.Println(fp)
	t, err := tail.TailFile(fp, tail.Config{
		Follow: true,
	})
	if err != nil {
		log.Fatalf("parse.ReadLine.TailFile err: %v", err)
		return err
	}

	// 监听
	for line := range t.Lines {
		logInfo := &model.LogInfo{}
		now := time.Now()
		formatNow := now.Format("2006-01-02 15:04:05")

		s := strings.Split(line.Text, " ")
		fmt.Println(s)
		if len(s) >= 5 && s[1] == "ERROR" {
			logInfo.Env = global.FileSetting.Env
			logInfo.ServiceName = global.FileSetting.ServiceName
			logInfo.Timestamp = s[0]
			logInfo.LogLevel = s[1]
			logInfo.LogKeyword = s[7]
			logInfo.LogInfo = s[7:]
			logInfo.StartAt = formatNow

			//写入数据库
			err := store.WriteLogInfo(global.DBEngine, logInfo)
			if err != nil {
				return err
			}

			var c int32
			var duration time.Duration = 5
			d := duration * time.Minute
			start := time.Now().Add(-d).Format("2006-01-02 15:04:05")
			end := time.Now().Format("2006-01-02 15:04:05")
			//统计次数
			c = process.Count(global.DBEngine, logInfo.LogKeyword, start, end)

			if c > 100 {
				fmt.Println(logInfo.LogKeyword, c)
				// 发送通知
				_, err = alert.QywechatAlert(logInfo)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
