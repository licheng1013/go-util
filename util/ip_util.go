package util

import (
	"fmt"
	"net"
	"strings"
)

// IpUtil 获取本机ip
type IpUtil struct {
}

func NewIpUtil() *IpUtil {
	return &IpUtil{}
}

// GetIp 获取ip 192.168.16.100/24
func (v IpUtil) GetIp() []string {
	var list []string
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, item := range addr {
		list = append(list, item.String())
	}
	return list
}

// GetTargetMaskIp 获取指定掩码位数的ip 8 16 24 32   分割 192.168.16.100/24 =>  192.168.16.100
func (v IpUtil) GetTargetMaskIp(digit interface{}) []string {
	digit = fmt.Sprintf("%v", digit)
	var list []string
	ip := v.GetIp()
	for _, item := range ip {
		split := strings.Split(item, "/")
		if split[1] == digit {
			list = append(list, split[0])
		}
	}
	return list
}
