package config

import (
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/pkg/mysql"
)

var DBEngine *gorm.DB

type MySQLConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  string        `mapstructure:"port"`
	UserName              string        `mapstructure:"username"`
	PassWord              string        `mapstructure:"password"`
	DataBase              string        `mapstructure:"database"`
	Charset               string        `mapstructure:"charset"`
	MaxIdleConnections    int           `mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `mapstructure:"max-connection-life-time"`
	LogLevel              int           `mapstructure:"log-level"`
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
