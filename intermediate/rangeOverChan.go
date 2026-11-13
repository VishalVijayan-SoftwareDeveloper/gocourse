package intermediate

import "fmt"

func RangeChan() {

	msg := make(chan string, 2)
	msg <- "Message1"
	msg <- "Message2"
	close(msg)
	//msg <- "Message3" // panic: send on closed channel
	for m := range msg {
		fmt.Println("recieved message:", m)
	}

}
