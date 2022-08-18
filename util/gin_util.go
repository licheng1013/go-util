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

// GetUrlPath 不带参数的  /user/test?name=xx => /user/test 请使用： GinUrlPath
// Deprecated
func GetUrlPath() string {
	return GinUrlPath()
}

// GinUrlPath 不带参数的  /user/test?name=xx => /user/test
func GinUrlPath() string {
	split := strings.Split(GinParamsUrlPah(), "?")
	path := split[0]
	return path
}

// GetParamsUrlPah 带参数的  /user/test?name=xx 请使用： GinParamsUrlPah
// Deprecated
func GetParamsUrlPah() string {
	return GinParamsUrlPah()
}

// GinParamsUrlPah 带参数的  /user/test?name=xx
func GinParamsUrlPah() string {
	c := GinGetContext()
	return c.Request.RequestURI
}

// GetGinContext 获取 GinContext 上下文 请使用： GinGetContext
// Deprecated
func GetGinContext() *gin.Context {
	return GinGetContext()
}

func GinGetContext() *gin.Context {
	value := GinUtil.Value(GinUtilKey)
	if value == nil {
		panic("没有在gin环境注册！")
	}
	return value.(*gin.Context)
}

// SetGinContext 设置 GinContext 上下文  请使用： GinSetContext
// Deprecated
func SetGinContext(c *gin.Context) {
	GinSetContext(c)
}

// GinSetContext 设置 GinContext 上下文
func GinSetContext(c *gin.Context) {
	GinUtil = context.WithValue(GinUtil, GinUtilKey, c)
}
