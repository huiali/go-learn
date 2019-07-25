package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			//for {
			fmt.Printf("hello from go routine %d \n", i)
			runtime.Gosched()
			//}
		}(i)
	}
	time.Sleep(time.Microsecond)
}

//轻量级 线程
//非抢占式多任务处理 ，由协程主动交出控制权
//编译器/解释器/虚拟机层面的多任务
