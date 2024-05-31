package config

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	MySQL MySQLConfig `yaml:"mysql"`
	JWT   JWTConfig   `yaml:"jwt"`
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
	conf.JWT.Init()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)
	return nil
}
