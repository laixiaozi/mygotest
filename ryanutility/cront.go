package ryanutility

//计时器
import (
	"fmt"
	"time"
)

//无限循环的任务
func HandlCront() {
	timeer := time.Tick(2 * time.Second)
	for now := range timeer {
		fmt.Printf("now : %v %s", now, " \n")
	}
}

/**
一个可以停止的计时器
*/
func HandlCrontAndStop() {

	Myticker := time.NewTicker(time.Second)
	defer Myticker.Stop()
	var d chan bool = make(chan bool, 1)
	//5秒后执行 <- d 这个部分的代码。读取  <- d 以后。这个地方的代码就不会再执行了
	go func() {
		time.Sleep(5 * time.Second)
		d <- true
	}()
	for {
		select {
		case <-d:
			fmt.Printf("done: \n")
		case t := <-Myticker.C:
			fmt.Println("current :", t)
			//default:
			//	fmt.Println("default now ", time.Now().Unix())
		}
	}

}
