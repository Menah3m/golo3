package setting

// 负责设置读取相关

import (
	"github.com/spf13/viper"
)

type Setting struct {
	Vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("golo-config")
	//如果不配置path，默认会从程序执行的当前目录寻找配置文件
	vp.AddConfigPath("configs/")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	//监听配置文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
