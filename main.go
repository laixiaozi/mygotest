package main

import (
	"./ryanutility"
	"fmt"
	"sync"
	"time"
)

var wg  sync.WaitGroup

type Job struct {
	Value int
}

func main() {
	fmt.Println("启动时间:", ryanutility.GetTodayString())
	//ryanutility.WebServerStart()
	//ryanutility.RpcServerStart()
	//ryanutility.HandlCront()  //计时器
	//ryanutility.HandlCrontAndStop()  //计时器var wg sync.WaitGroup
	  arr := []int{1,2,3,4,5,6,7,8,10}
	  ch := make(chan int)
	  wg.Add(2)
	  go worker(ch)
	  go worker(ch)
	  for _, value := range arr {
	     ch <- value
	  }
	  close(ch)
	  wg.Wait()
	  fmt.Println("Done")



}
var i int = 1
var tm *time.Ticker = time.NewTicker(time.Second * 1)
func worker(ch chan int){
	for {
		 select {
		    case c := <-ch :{
		    	fmt.Println("print : ",c)
		    	if c < 10 {
		    		continue
				}
		    	wg.Done()
		    	break
			}
		    case <- tm.C :
		      wg.Done()
		      break
		 }

	}
}

