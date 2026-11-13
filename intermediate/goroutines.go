package intermediate

import (
	"fmt"
	"time"
)

func display(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func Goroutines() {

	go display("goroutine")
	//display("direct")
	time.Sleep(1 * time.Second)
}
