package model

import (
	"github.com/lib/pq"
	"time"
)

//负责定义日志信息的数据结构

/*
建表sql

create table log_info
(
    id          int(10) auto_increment,
    log_level   varchar(100)         not null comment '日志级别',
    log_keyword varchar(255)         not null comment '日志关键词',
    log_info    longtext             null comment '日志内容详情',
    start_at    int(10)              not null comment '日志出现时间',
    end_at      int                  not null comment '日志内容监控 结束时间',
    is_silence  tinyint(3) default 0 not null comment '0 表示未静默  1 表示静默',
    constraint log_info_pk
        primary key (id)
)
    comment '存放日志信息';
*/

type LogInfo struct {
	// id
	Id int32 `json:"id"`
	// 环境
	Env string `json:"env"`
	// 服务名
	ServiceName string `json:"service_name"`
	// 日志时间戳
	Timestamp string `json:"timestamp"`
	// 日志级别
	LogLevel string `json:"log_level"`
	// 日志关键词
	LogKeyword string `json:"log_keyword"`
	// 日志内容详情
	LogInfo pq.StringArray `json:"log_info"`
	// 日志出现时间
	StartAt string `json:"start_at"`
	// 0 表示未静默  1 表示静默
	IsSilence int8 `json:"is_silence"`
}

type CountInfo struct {
	// id
	Id int32 `json:"id"`
	// 环境
	Env string `json:"env"`
	// 服务名
	ServiceName string `json:"service_name"`
	// 日志级别
	LogLevel string `json:"log_level"`
	// 日志关键词
	LogKeyword string `json:"log_keyword"`
	// 日志内容详情
	LogInfo pq.StringArray `json:"log_info"`
	// 日志出现时间
	StartAt string `json:"start_at"`
	// 时间跨度
	Duration time.Duration `json:"duration"`
	// 统计次数
	Count int32 `json:"count"`
	// 0 表示未静默  1 表示静默
	IsSilence int8 `json:"is_silence"`
}

func (model LogInfo) TableName() string {
	return "log_info"
}

func (model CountInfo) TableName() string {
	return "count_info"
}
