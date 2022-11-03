package util

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomUtil 随机工具
type RandomUtil struct {
	alphabetMap       string
	alphabetMapLength int
}

// NewRandomUtil 构建对象
func NewRandomUtil() *RandomUtil {
	r := &RandomUtil{}
	r.alphabetMap = r.alphabet()
	r.alphabetMapLength = len(r.alphabetMap)
	return r
}

// 字母表
func (r RandomUtil) alphabet() string {
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
func (r RandomUtil) RandomString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		v := r.alphabetMap[r.RandomRangeNum(r.alphabetMapLength)]
		str += fmt.Sprintf("%c", v)
	}
	return str
}

// RandomRangeNum 获取指定返回的数字  输入 30 返回 0-29
func (r RandomUtil) RandomRangeNum(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// RandomRangeNumPlus 获取指定返回的数字  输入 10,30 返回 10-29
func (r RandomUtil) RandomRangeNumPlus(min, max int) int {
	num := r.RandomRangeNum(max)
	if num <= min {
		num += min
	}
	return num
}

// RandomNumber 获取指定长度的数字字符串
func (r RandomUtil) RandomNumber(length int) string {
	var strNumber string
	for i := 0; i < length; i++ {
		strNumber += fmt.Sprintf("%v", r.RandomRangeNum(10))
	}
	return strNumber
}
