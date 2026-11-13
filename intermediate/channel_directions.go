package intermediate

import "fmt"

func ping(ping chan<- string, msg string) {
	ping <- msg

}
func pong(ping <-chan string, pong chan<- string) {
	msgs := <-ping
	pong <- msgs
}

func Channel() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "Hello")
	go pong(pings, pongs)
	fmt.Println(<-pongs)

}
