package main

import (
	"./ryanutility"
	"fmt"
)

func main() {
	fmt.Println("启动时间:", ryanutility.GetTodayString())
	//ryanutility.WebServerStart()
	//ryanutility.RpcServerStart()
	//ryanutility.HandlCront()  //计时器
	//ryanutility.HandlCrontAndStop()  //计时器
}
