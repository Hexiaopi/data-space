package config

import (
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/pkg/mysql"
)

var DBEngine *gorm.DB

type MySQLConfig struct {
	Host                  string        `yaml:"host"`
	Port                  string        `yaml:"port"`
	UserName              string        `yaml:"username"`
	PassWord              string        `yaml:"password"`
	DataBase              string        `yaml:"database"`
	Charset               string        `yaml:"charset"`
	MaxIdleConnections    int           `yaml:"max-idle-connections"`
	MaxOpenConnections    int           `yaml:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `yaml:"max-connection-life-time"`
	LogLevel              int           `yaml:"log-level"`
}

func (o *MySQLConfig) NewClient() (*gorm.DB, error) {
	conf := &mysql.Config{
		Host:                  o.Host,
		Port:                  o.Port,
		UserName:              o.UserName,
		PassWord:              o.PassWord,
		DataBase:              o.DataBase,
		Charset:               o.Charset,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
	}
	return mysql.New(conf)
}
