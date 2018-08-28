package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

/**
	变量定义
 */

var (
	aa = 3
 	ss = "kkk"
 	bb = true
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter()  {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 复数
func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%.3f", cmplx.Exp(1i * math.Pi) + 1)
}

// 强制类型转换
func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTaiangle(a, b))
}

func calcTaiangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

// 常量，
func consts()  {
	const (
		filename  = "abc.txt"
		a, b  = 3, 4 // 没定义类型可以当成任意类型
	)

	var c int
	c = int(math.Sqrt(a * a + b * b)) // 这里当成float64
	fmt.Println(filename, c)
}

// 枚举
func enums()  {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()

	triangle()

	consts()

	enums()
}
