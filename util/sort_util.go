package util

import "sort"

// SortInt 排序int，示例：3,2,1 => 1,2,3
func SortInt(t []int) []int {
	list := sort.IntSlice(t[0:])
	sort.Sort(list)
	return list
}

// SortString 排序字符串，示例：b,c,a => a,b,c
func SortString(t []string) []string {
	list := sort.StringSlice(t[0:])
	sort.Sort(list)
	return list
}
