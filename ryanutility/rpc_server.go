package ryanutility

import (
	"../server"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//RpcServerStart 启动一个rpc http服务
func RpcServerStart() {
	arg := new(server.RpcTest)
	rpc.Register(arg)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("rpc http serve err :", err)
		return
	}
	 http.Serve(listen, nil)
}
