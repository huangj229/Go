package test

import (
    "fmt"
	"net"    
)

func TestNet() {
	ip,err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("www.baidu.com: ", ip)
}