package test

import (
	"../server"
	"fmt"
	"net/rpc"
)

type RpcTest struct {
	A, B int64
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	test := &server.RpcTest{A: 10, B: 60}
	var b int64
	err = client.Call("RpcTest.Add", test, &b)
	fmt.Println(err)
	fmt.Println(b)
}
