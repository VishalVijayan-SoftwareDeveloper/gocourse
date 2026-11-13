package intermediate

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSync() {

	done := make(chan bool)
	go worker(done)
	res := <-done
	if res {
		fmt.Println("done from worker")
	}

}
