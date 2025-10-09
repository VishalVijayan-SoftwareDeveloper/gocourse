package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var primary, secondary int32
	n := len(arr)
	for i := 0; i < n; i++ {
		primary += arr[i][i]
		secondary += arr[i][n-i-1]
	}
	return int32(math.Abs(float64(primary - secondary)))
}

func main() {
	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	fmt.Println(arr)
	res := diagonalDifference(arr)
	fmt.Println(res)
}
