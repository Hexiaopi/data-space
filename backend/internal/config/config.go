package config

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	HTTP  HttpConfig  `yaml:"http"`
	MySQL MySQLConfig `yaml:"mysql"`
	Log   LogConfig   `yaml:"log"`
	JWT   JWTConfig   `yaml:"jwt"`
	Trace TraceConfig `yaml:"trace"`
}

func Init(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var conf config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return err
	}
	fmt.Println(conf)
	DBEngine, err = conf.MySQL.NewClient()
	if err != nil {
		return err
	}
	Logger = conf.Log.NewLog()
	HTTP = conf.HTTP
	conf.JWT.Init()
	conf.Trace.NewShutdown()
	return nil
}

func Close() {
	db, _ := DBEngine.DB()
	db.Close()
	if shutdown != nil {
		shutdown(context.Background())
	}
}
