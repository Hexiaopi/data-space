package config

import (
	"time"

	"github.com/hexiaopi/data-space/pkg/jwt"
)

type JWTConfig struct {
	// 过期时间
	Expire time.Duration `yaml:"expire"`
	// 签发人
	Issuer string `yaml:"issuer"`
	// 秘钥
	Secret string `yaml:"secret"`
}

var JWT jwt.JWT

func (c JWTConfig) Init() {
	JWT = jwt.NewJWT(c.Expire, c.Issuer, c.Secret)
}
