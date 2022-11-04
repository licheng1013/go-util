package test

import (
	"errors"
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
)

func TestSortInt(t *testing.T) {
	v := []int{3, 2, 1}
	v = common.SortUtil.SortInt(v)
	log.Println(v)
	for i := 0; i < 3; i++ {
		if i+1 != v[i] {
			panic(errors.New("校对失败！"))
		}
	}
}

func TestSortStr(t *testing.T) {
	check := []string{"a", "b", "c"}

	v := []string{"b", "c", "a"}
	v = common.SortUtil.SortString(v)
	log.Println(v)
	for i := 0; i < 3; i++ {
		if v[i] != check[i] {
			panic(errors.New("校对失败！"))
		}
	}
}
