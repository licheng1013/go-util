/*
 *                                  Apache License
 *                            Version 2.0, January 2004
 *                         http://www.apache.org/licenses/
 */

package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const key = "izVguZPRsBQ5Rqw6dhMvcIwy8_9lQnrO3vpxGwPxfAxDs"
const UserIdKey = "userId"
const expireTimeKey = "expireTime"

func GetToken(id interface{}) string {
	mySigningKey := []byte(key)
	claims := &jwt.MapClaims{
		UserIdKey:     id,
		expireTimeKey: time.Now(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return t
}

// GetUserId  获取token解析，请使用 GetTokenParse 替代
// Deprecated
func GetUserId(t string) string {
	return GetTokenParse(t)
}

// GetTokenParse 获取token解析
func GetTokenParse(t string) string {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	//log.Println(token.Header)
	//log.Println(token.Claims)
	if err != nil {
		panic(err)
	}
	id := (token.Claims.(jwt.MapClaims))["userId"]
	return fmt.Sprintf("%v", id)
}
