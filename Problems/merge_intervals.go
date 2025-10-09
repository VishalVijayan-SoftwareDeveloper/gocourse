/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
*/

package main

import (
	"fmt"
	"sort"
)

func mergeInterval(interval [][]int) [][]int {
	fmt.Println(interval)

	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	fmt.Println("after sort:", interval)
	merged := [][]int{interval[0]}
	for _, val := range interval[1:] {

		last := merged[len(merged)-1]
		if val[0] <= last[1] {
			fmt.Println("overlap")
			if val[1] > last[1] {
				last[1] = val[1]
			}
			merged[len(merged)-1] = last
		} else {
			merged = append(merged, val)
		}
	}
	return merged

}
func main() {
	interval := [][]int{{1, 3}, {2, 6}, {15, 18}, {8, 10}}

	res := mergeInterval(interval)
	fmt.Println("Result:", res)
}
