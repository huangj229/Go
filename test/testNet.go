package test

import (
    "fmt"
	"net"
	"bufio"    
)



func TestNet() {
	// resolve IP of baidu.com
	ip,err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("www.baidu.com: ", ip)


    //build connection to baidu.com
	conn,err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	//ask for response from baidu.com
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("status:", status)


}