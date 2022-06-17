package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job ", j)
		res <- j * 2
	}
}

func main() {
	start := time.Now()
	const numjobs int = 5
	jobs := make(chan int, numjobs)
	res := make(chan int, numjobs)

	var wg sync.WaitGroup

	for i := 1; i <= numjobs; i++ {
		wg.Add(1)
		jobs <- i
		i := i
		go func() {
			defer wg.Done()
			worker(i, jobs, res)

		}()

	}
	close(jobs)
	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration)
}
