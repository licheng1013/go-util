/*
 *                                  Apache License
 *                            Version 2.0, January 2004
 *                         http://www.apache.org/licenses/
 */

package util

import (
	"fmt"
	"math/rand"
	"time"
)

var alphabetMap = alphabet()
var alphabetMapLength = len(alphabetMap)

// 字母表
func alphabet() string {
	var str string
	for i := 'a'; i <= 'z'; i++ {
		str += fmt.Sprintf("%c", i)
	}
	for i := 'A'; i <= 'Z'; i++ {
		str += fmt.Sprintf("%c", i)
	}
	return str
}

// RandomString 获取指定长度的随机字符串  输入 3 返回 xna
func RandomString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		v := alphabetMap[RandomRangeNum(alphabetMapLength)]
		str += fmt.Sprintf("%c", v)
	}
	return str
}

// RandomRangeNum 获取指定返回的数字  输入 30 返回 0-29
func RandomRangeNum(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// RandomRangeNumPlus 获取指定返回的数字  输入 10,30 返回 10-29
func RandomRangeNumPlus(min, max int) int {
	num := RandomRangeNum(max)
	if num <= min {
		num += min
	}
	return num
}

// RandomNumber 获取指定长度的随机数字字符串  输入 5 返回 16523
func RandomNumber(length int) string {
	var strNumber string
	for i := 0; i < length; i++ {
		strNumber += fmt.Sprintf("%v", RandomRangeNum(10))
	}
	return strNumber
}
