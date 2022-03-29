package setting

// 负责 设置相关的数据结构定义

import "time"

type AppSettings struct {
	RunMode   string
	Duration  time.Duration
	Threshold int
}

type FileSettings struct {
	Env         string
	ServiceName string
	LogSavePath string
	LogFileName string
	LogFileExt  string
	Keyword     string
}

type DatabaseSettings struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type QywechatSettings struct {
	CorpID  string
	AgentID int32
	Secret  string
	ToUser  string
	MsgType string
}

type EmailSettings struct {
	Host string
	Port string
	From string
	To   string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.Vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
