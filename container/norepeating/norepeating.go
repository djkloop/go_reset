package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i

	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("111222333"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkewa"))
	fmt.Println(lengthOfNonRepeatingSubStr("对啊对啊"))
}
