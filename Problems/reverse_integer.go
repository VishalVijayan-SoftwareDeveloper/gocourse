package main

import (
	"fmt"
	"math"
)

func reverseInteger(num int) int {

	reversed := 0
	for num != 0 {
		rem := num % 10
		num /= 10
		if reversed > math.MaxInt32/10 || reversed < math.MinInt32/10 {
			return 0
		}
		reversed = reversed*10 + rem
	}

	return reversed

}

func main() {
	res := reverseInteger(123)
	fmt.Println(res)
}
