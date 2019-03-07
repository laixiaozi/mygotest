package main

import (
	"./ryanutility"
	"fmt"
)

func main() {
	fmt.Println("启动时间:", ryanutility.GetTodayString())
	ryanutility.WebServerStart()
}
