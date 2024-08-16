package config

import (
	log "github.com/hexiaopi/data-space/pkg/logger"
)

type LogConfig struct {
	FileName  string `yaml:"file-name"` // 日志文件名称
	LogLevel  string `yaml:"log-level"` // debug, info, warn, error, fatal, panic
	MaxSize   int    `yaml:"max-size"`  // 日志文件最大大小
	MaxAge    int    `yaml:"max-age"`   // 日志文件最大保存时间
	MaxBackup int    `yaml:"max-back"`  // 日志文件最大保存数量
	Compress  bool   `yaml:"compress"`  // 是否压缩日志文件
	Encoding  string `yaml:"encoding"`  // console, json
	Env       string `yaml:"env"`       // 日志文件环境

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
