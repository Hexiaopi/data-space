package config

import (
	"context"

	"github.com/hexiaopi/data-space/pkg/trace"
)

type TraceConfig struct {
	URL        string `yaml:"url"`
	ServerName string `yaml:"server-name"`
}

var shutdown func(context.Context) error

func (conf TraceConfig) NewShutdown() error {
	var err error
	shutdown, err = trace.InitTrace(conf.URL, conf.ServerName)
	if err != nil {
		return err
	}
	return nil
}
