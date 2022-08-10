/*
 * 作者：Lc
 */

package util

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// GetIp 获取ip 192.168.16.100/24
func GetIp() []string {
	var list []string
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, item := range addr {
		log.Println(item.String())

		list = append(list, item.String())
	}
	return list
}

// GetTargetMaskIp 指定掩码 8 16 24 32   分割/ 192.168.16.100/24 =>  192.168.16.100
func GetTargetMaskIp(str interface{}) []string {
	str = fmt.Sprintf("%v", str)
	var list []string
	ip := GetIp()
	for _, item := range ip {
		split := strings.Split(item, "/")
		if split[1] == str {
			list = append(list, split[0])
		}
	}
	return list
}
