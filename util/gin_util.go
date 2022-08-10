/*
 *                                  Apache License
 *                            Version 2.0, January 2004
 *                         http://www.apache.org/licenses/
 */

package util

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

var GinUtil = context.Background()

const GinUtilKey = "GinUtilKey"

// GetUrlPath 不带参数的  /user/test?name=xx => /user/test
func GetUrlPath() string {
	split := strings.Split(GetParamsUrlPah(), "?")
	path := split[0]
	return path
}

// GetParamsUrlPah 带参数的  /user/test?name=xx
func GetParamsUrlPah() string {
	c := GetGinContext()
	return c.Request.RequestURI
}

// GetGinContext 获取 GinContext 上下文
func GetGinContext() *gin.Context {
	value := GinUtil.Value(GinUtilKey)
	if value == nil {
		panic("没有在gin环境注册！")
	}
	return value.(*gin.Context)
}

// SetGinContext 设置 GinContext 上下文
func SetGinContext(c *gin.Context) {
	GinUtil = context.WithValue(GinUtil, GinUtilKey, c)
}
