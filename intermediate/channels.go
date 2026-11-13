package intermediate

import (
	"fmt"
	"time"
)

func ChannelTest() {

	messsages := make(chan string)
	go func() {
		messsages <- "ping"
		messsages <- "pong"
	}()

	recieveMessage := <-messsages
	fmt.Println(recieveMessage)
	time.Sleep(2 * time.Second)
	recieveMessage2 := <-messsages
	fmt.Println(recieveMessage2)
}
