package main

import (
	"fmt"
)

// expandAroundCenter expands around a center and returns the start and end indexes of the palindrome.
func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

// longestPalindrome finds the longest palindromic substring in s.
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int = 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return (s[start : end+1])
}

func main() {
	//input := "babad"
	input := "121222111"
	fmt.Println("Input:", input)
	fmt.Println("Longest Palindromic Substring:", longestPalindrome(input))

	input = "dbbd"
	fmt.Println("\nInput:", input)
	fmt.Println("Longest Palindromic Substring:", longestPalindrome(input))
}
