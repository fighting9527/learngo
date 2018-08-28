package main

import (
	"fmt"
)

/**
    @author good mood
    2018/7/15 下午6:57
*/

func tryRecover()  {
	defer func() {
		r := recover()
		fmt.Println(r)
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't knoe what to do %v", r))
		}
	}()

	//panic(errors.New("this is an error"))
	panic(123)
}

func main() {
	//tryRecover()
	//rand.Seed(time.Now().Unix())
	//str := ""
	//for i := 0; i < 100 ; i++ {
	//	str += strconv.Itoa(rand.Intn(9))
	//}
	//fmt.Println(str)

	str4 := "0885561236802855410406168118126810542687524615265466722730743412010643175484874043723174016886720255"

	str3 := "08855612368028554104061681181268" +
			"105426875246152654667227307434120106" +
			"43175484874043723174016886720255"
	fmt.Println(str3 == str4)

	
	//str1 := "5622467846573248022520510771634655386730741010711307064783758633188852108560147860712725812461043072"
	//
	//str2 := "5622467846573248022520510771634655386730741010711307064783758633188852108560147860712725812461043072"
	//fmt.Println(str2 == str1)
}
