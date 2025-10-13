/*
input: cars := []string{"red", "blue", "green", "red", "yellow","green", "red"} min=2 and max=3

output= [["red","red","red"], ["green","green"]
*/

package main

import "fmt"

func CountColorCar(cars []string, minm, maxm int) [][]string {

	countColor := make(map[string]int)

	for _, color := range cars {
		countColor[color]++
	}
	fmt.Println(countColor)
	var res [][]string

	for color, count := range countColor {

		if count >= minm && count <= maxm {
			group := make([]string, count)
			for i := range group {
				group[i] = color
			}
			res = append(res, group)
		}

	}
	fmt.Println(res)
	return res

}

func main() {

	cars := []string{"red", "blue", "green", "red", "yellow", "green", "red"}
	min := 2
	max := 3
	CountColorCar(cars, min, max)
}
