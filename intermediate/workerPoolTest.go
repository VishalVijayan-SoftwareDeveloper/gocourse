package intermediate

import (
	"fmt"
	"time"
)

func WorkerPool(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Workers id:", id, "started job:", j)
		time.Sleep(2 * time.Second)
		fmt.Println("Workers id:", id, "finished job:", j)
		results <- j * 2
	}
}

func WorkerPoolTes() {
	numJobs := 5
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 1; i <= numWorkers; i++ {
		go WorkerPool(i, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	for r := 1; r <= numJobs; r++ {

		fmt.Println("Job finished and results:", <-results)
	}

}
