package config

import (
	"time"

	"github.com/hexiaopi/data-space/pkg/jwt"
)

type JWTConfig struct {
	Expire time.Duration `yaml:"expire"` // 过期时间
	Issuer string        `yaml:"issuer"` // 签发人
	Secret string        `yaml:"secret"` // 秘钥
}

var JWT jwt.JWT

func (c JWTConfig) Init() {
	JWT = jwt.NewJWT(c.Expire, c.Issuer, c.Secret)
}
