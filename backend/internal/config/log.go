package config

import (
	log "github.com/hexiaopi/data-space/pkg/logger"
)

type LogConfig struct {
	FileName  string `mapstructure:"filename"`  // 日志文件名称
	LogLevel  string `mapstructure:"log-level"` // debug, info, warn, error, fatal, panic
	MaxSize   int    `mapstructure:"max-size"`  // 日志文件最大大小
	MaxAge    int    `mapstructure:"max-age"`   // 日志文件最大保存时间
	MaxBackup int    `mapstructure:"max-back"`  // 日志文件最大保存数量
	Compress  bool   `mapstructure:"compress"`  // 是否压缩日志文件
	Encoding  string `mapstructure:"encoding"`  // console, json
	Env       string `mapstructure:"env"`       // 日志文件环境

}

var Logger log.Logger

func (o *LogConfig) NewLog() log.Logger {
	conf := &log.Config{
		FileName:  o.FileName,
		LogLevel:  o.LogLevel,
		MaxSize:   o.MaxSize,
		MaxBackup: o.MaxBackup,
		MaxAge:    o.MaxAge,
		Compress:  o.Compress,
		Encoding:  o.Encoding,
		Env:       o.Env,
	}
	return log.New(conf)
}
