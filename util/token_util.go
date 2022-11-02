package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenUtil 对token工具的封装
type TokenUtil struct {
	// 密钥
	key string
	// 参数
	Claims jwt.RegisteredClaims
	// 是否检查过期时间，默认false不检查
	IsCheckTokenExpired bool
}

func NewTokenUtil(key string) *TokenUtil {
	return &TokenUtil{key: key}
}

// CustomToken 获取自定义token
func (util TokenUtil) CustomToken() string {
	mySigningKey := []byte(util.key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, util.Claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return t
}

// DefaultToken 默认token 附带颁发时间: 此刻，过期时间: 此刻延后三天
func (util TokenUtil) DefaultToken(id interface{}) string {
	mySigningKey := []byte(util.key)
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprint(id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour * 24 * 3)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return t
}

// ParseToken 获取token解析
func (util TokenUtil) ParseToken(t string) string {
	token, err := jwt.ParseWithClaims(t, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.key), nil
	})
	// 是否不检查过期判断
	if !util.IsCheckTokenExpired && errors.Is(err, jwt.ErrTokenExpired) {
		err = nil
	}
	if err != nil {
		panic(err)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Subject
}
