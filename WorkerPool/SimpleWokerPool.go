package main

import (
	"sync"
	"time"
)

type job struct {
	id int
}

func doWork(jobs <-chan job, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Second * 3)
		println(job.id, "processing this")
	}

}

func main() {

	const numWokers int = 5

	jobs := make(chan job, 100)

	wg := sync.WaitGroup{}

	// create a worker pool
	for range numWokers {
		wg.Add(1)
		go doWork(jobs, &wg)
	}

	// populate jobs
	for i := range 5 {
		jobs <- job{id: i}
	}
	close(jobs)

	wg.Wait()

}
