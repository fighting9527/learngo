package main

import (
	"fmt"

)

func main() {
	m := map[string]string {
		"name" : "cao",
		"study" : "go",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int      // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name := m["name"]
	fmt.Println(name)
	if noName, ok := m["noname"]; ok {
		fmt.Println(noName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("deleteing values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println("longStr: ", getLognStr("abcbctyurpbbdfg"))


}

// abcbctyurpbbdfg -> abc
func getLognStr(str string) string {
	fmt.Println("start:", str)
	maxStr := string(str[0])
	maxIndex := 0
	for i := 1; i < len(str); i ++ {
		for j := maxIndex; j < i ; j ++  {
			if str[j] == str[i] {
				fmt.Println("i = ", i, " == str[maxIndex:i]:", str[maxIndex:i])
				fmt.Println("== maxStr:", maxStr)
				if len(str[maxIndex:i]) > len(maxStr) {
					maxStr = str[maxIndex : i]
					fmt.Println("maxStr:", maxStr)
				}
				fmt.Println("下标：",i,"，子串：",str[maxIndex : i])
				maxIndex = i
				break
			}
		}
		fmt.Println("i = ", i, " for over str[maxIndex:i]:", str[maxIndex:i])
		fmt.Println("for over maxStr:", maxStr)
		if len(str[maxIndex:i + 1]) > len(maxStr) {
			maxStr = str[maxIndex : i + 1]
			fmt.Println("maxStr:", maxStr)
			fmt.Println("下标：",i,"，子串：",maxStr)
		}
	}
	return maxStr
}
