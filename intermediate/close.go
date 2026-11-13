package intermediate

import "fmt"

func CloseTest() {
	ch := make(chan int)
	//sender goroutine. Sender should close the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	//ch <- 10, panic send on closed channel

	//one method to receive
	/*for val := range ch {
		fmt.Println("Recieved value:", val)
	}*/
	//another method receive
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Recieved:", val)
	}
}
