package intermediate

import (
	"fmt"
	"time"
)

func workerPool(id int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		fmt.Printf("workder: %d ,Job started: %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("workder: %d ,Job finished: %d\n", id, i)
		result <- i * 2
	}
}

func Worker_pool() {

	numJobs := 5
	numWorkers := 2
	jobs := make(chan int, numJobs)
	result := make(chan int, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go workerPool(i, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for res := range result {
		fmt.Println(res)
	}
	close(result)
	fmt.Println("Program finished")

}
