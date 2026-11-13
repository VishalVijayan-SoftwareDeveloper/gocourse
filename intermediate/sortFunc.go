package intermediate

import (
	"cmp"
	"fmt"
	"slices"
)

func SortFunc() {

	fruits := []string{"Kiwi", "Banana", "Apple", "Orange"}

	slices.Sort(fruits)
	fmt.Println("AFter sorting:", fruits)

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)

	fmt.Println("After sorting with length:", fruits)

}
