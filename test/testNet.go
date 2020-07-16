package test

import (
    "fmt"
	"net"
	"io"
	"strings"    
    "../config"
)



func TestNet() {
	config := config.LoadConfig()
	
	// resolve IP of baidu.com
	ip,err := net.ResolveIPAddr("ip", config.Site.Domain)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(config.Site.Domain, ": ", ip)


    //build connection to baidu.com
	conn,err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	//close tcp con
	defer conn.Close()

	//send get to baidu.com ,HTTP 1.0
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	
	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF{
				fmt.Println("read error :", err)
			}
			break
		} 
		sb.Write(buf[:n])
	}
	// display result
	// http introduce https://www.ruanyifeng.com/blog/2016/08/http.html
	fmt.Println("response:", sb.String())
	fmt.Println("total response size:", sb.Len())


}