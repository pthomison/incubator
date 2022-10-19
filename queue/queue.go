package queue

import (
	"sync"
)

func Workqueue[T any](objArray []T, numberOfWorkers int, workerFunction func(*sync.WaitGroup, chan uint64, map[uint64]T)) {
	var wg sync.WaitGroup
	index := make(map[uint64]T)

	jobChan := make(chan uint64, 2*numberOfWorkers)

	for i, obj := range objArray {
		index[uint64(i)] = obj
	}

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go workerFunction(&wg, jobChan, index)
	}

	for i := range index {
		jobChan <- i
	}

	close(jobChan)

	wg.Wait()
}
