package main

import (
	"fmt"
	"time"
)

/**
  @author good mood
  2018/7/16 下午11:45
*/
func worker(id int, c chan int) {
	go func() {
		for n := range c{
			//n, ok := <-c
			//if !ok {
			//	break
			//}
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {

			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func bufferedChannel()  {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
