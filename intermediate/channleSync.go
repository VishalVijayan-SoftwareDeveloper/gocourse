package intermediate

import (
	"fmt"
	"time"
)

func employer(done chan bool) {
	fmt.Print("I am working and wait for to finish...")

	time.Sleep(10 * time.Second)

	fmt.Print("I am finished, exit!")

	done <- true
}

func ChannelSyncTest() {
	done := make(chan bool)

	go employer(done)
	<-done

}
