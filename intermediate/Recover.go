package intermediate

import "fmt"

func myPanic() {
	panic("a problem")
}

func RecoverFunc() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:", r)
		}
	}()
	myPanic()
	fmt.Println("Done and exit") // this will not execute, because main functions stops at the point of panic and resumes in deferred closure
}
