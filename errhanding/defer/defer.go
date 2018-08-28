package main

import (
	"fmt"
	"os"
	"bufio"
	"learngo/functional/fib"
)

/**
    @author good mood
    2018/7/15 下午5:52
*/

func tryDefer()  {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	//writeFile("fib.txt")
}
