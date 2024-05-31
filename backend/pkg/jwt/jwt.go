package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

type JWT struct {
	// 过期时间
	Expire time.Duration
	// 签发人
	Issuer string
	// 秘钥
	Secret string
}

func NewJWT(expire time.Duration, issuer, secret string) JWT {
	return JWT{
		Expire: expire,
		Issuer: issuer,
		Secret: secret,
	}
}

func (j JWT) GenerateToken(id int64, username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(j.Expire)
	claims := Claims{
		UserId:   id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  nowTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(j.Secret))
	return token, err
}

func (j JWT) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
