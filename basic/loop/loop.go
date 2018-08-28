package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

// 整数转二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0 ; n /=2 {
		lsb := n %2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever()  {
	for  {
		fmt.Println("abc")
	}

}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1011 -- > 1101
	)

	printFile("basic/abc.txt")

	s := `abc"d"
		kkk`
	printFileContents(strings.NewReader(s))

	//forever()
}
