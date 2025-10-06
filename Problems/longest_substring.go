package main

import "fmt"

func longestSubstring(s string) int {

	myMap := make(map[rune]int)
	maxLength := 0
	start := 0
	for i, val := range s {
		if prevIndex, ok := myMap[val]; ok && prevIndex >= start {
			start = prevIndex + 1
			fmt.Println("STart:", start)
		}
		myMap[val] = i
		fmt.Println("Map:", myMap)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		fmt.Println("MaxLength", maxLength)
	}
	return maxLength
}

func main() {
	str := "abcabcdab"
	result := longestSubstring(str)
	fmt.Println(result)
}
