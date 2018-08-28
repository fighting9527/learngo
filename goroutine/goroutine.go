package main

import (
	"time"
	"fmt"
)

/**
    @author good mood
    2018/7/16 下午10:08
*/

// main也是一个协程在执行
func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for  {
				fmt.Printf("Hello from goroutine %d\n", i)
				//a[i]++
				//// 控制权
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
