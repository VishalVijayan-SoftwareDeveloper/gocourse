package intermediate

import "fmt"

func Worker(jobs chan int, done chan bool) {
	for {
		j, more := <-jobs
		if more {
			fmt.Println("job recieved:", j)
		} else {
			fmt.Println("job finished")
			done <- true
			return
		}
	}
}

func CloseChannel() {

	jobs := make(chan int)
	done := make(chan bool)

	go Worker(jobs, done)

	for j := 1; j <= 3; j++ {
		fmt.Println("job sent: ", j)
		jobs <- j
	}
	close(jobs)
	fmt.Println("closed all jobs")
	<-done
	fmt.Println("sending one more job on closed channel")

	val, ok := <-jobs
	fmt.Println("recived on closed channel:", ok, val)

}
