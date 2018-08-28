package main

import (
	"fmt"
	"sync"
)

/**
  @author good mood
  2018/7/16 下午11:45
*/
func doWorker(id int, c chan int,
	//done chan bool
	//wg *sync.WaitGroup,
	w worker,
	) {
	go func() {
		for n := range c{
			//n, ok := <-c
			//if !ok {
			//	break
			//}
			fmt.Printf("Worker %d received %d\n", id, n)

			//done <- true
			//wg.Done()
			w.done()
		}
	}()
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}


func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w)
	return w
}

func chanDemo() {
	//var c chan int // c == nil
	var wg sync.WaitGroup


	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers{
		worker.in <- 'a' + i
	}

	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i, worker := range workers{
		worker.in <- 'A' + i
	}

	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}
}
func bufferedChannel()  {
	//c := make(chan int, 3)
	//go worker(0, c)
	//c <- 'a'
	//c <- 'b'
	//c <- 'c'
	//c <- 'd'
	//time.Sleep(time.Millisecond)
}

func channelClose()  {
	//c := make(chan int)
	//go worker(0, c)
	//c <- 'a'
	//c <- 'b'
	//c <- 'c'
	//c <- 'd'
	//close(c)
	//time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
