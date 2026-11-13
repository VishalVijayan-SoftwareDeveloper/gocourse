package intermediate

import "fmt"

func BufChannel() {

	messages := make(chan string, 2)

	messages <- "first"
	messages <- "second message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
