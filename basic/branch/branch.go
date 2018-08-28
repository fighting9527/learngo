package main

import (
	"io/ioutil"
	"fmt"
)

// switch后可以没有表达式，不需要break
func grade(score int) string  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	// 条件语句，if条件里可以定义变量，没有括号
	const filename  = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(101),
		)
	
}
