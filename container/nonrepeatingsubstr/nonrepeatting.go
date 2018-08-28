package main

import "fmt"

//func lengthOfNonRepeatingSubStr(s string) int {
//	lastOccured := make(map[rune]int)
//	start := 0
//	maxLength := 0
//	for i, ch := range []rune(s){
//
//		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
//			start = lastI + 1
//		}
//		if i - start + 1 > maxLength {
//			maxLength = i - start + 1
//		}
//		lastOccured[ch] = i
//	}
//	return maxLength
//}
var lastOccured = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int {
	// stores last occurred pos + 1.
	// 0 means not seen.

	for i := range lastOccured {
		lastOccured[i] = 0
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s){

		if lastI := lastOccured[ch]; lastI >= start {
			start = lastI
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i + 1
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}
