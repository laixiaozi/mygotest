package ryanutility

import (
	"fmt"
	"net"
)

/**
* 启动一个监听端口
 */
func WebServerStart() {
	fmt.Println("启动一个web服务端口")
	RyanListener ,err := net.Listen("tcp",":8081")
	if err != nil{
		fmt.Println("创建net listener", err)
	}
   defer RyanListener.Close()
	for {
       netCon , err :=  RyanListener.Accept()
		if err != nil{
			fmt.Println("链接错误", err)
		}

	}
}
