package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	jobs := make(chan Task, len(tasks))
	returnChannel := make(chan int, 1)
	mutex := sync.Mutex{}
	//wg := sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		//wg.Add(1)
		go func(i int) {
			for job := range jobs {
				fmt.Printf("job started by goroutine %d\n", i)
				if err := job(); err != nil {
					fmt.Printf("job finished by goroutine %d with error: %v\n", i, err)
					mutex.Lock()
					m--
					if m == 0 {
						mutex.Unlock()
						//wg.Done()
						returnChannel <- 1
						close(jobs)
						fmt.Printf("goroutine %d closed jobs channel\n", i)
						return
					}
					mutex.Unlock()
				}
				//if len(jobs) == 0 {
				//	fmt.Printf("goroutine %d empty jobs channel\n", i)
				//	returnChannel <- 0
				//}
			}
		}(i)
	}

	for job := range tasks {
		jobs <- tasks[job]
	}

	returnValue, _ := <-returnChannel

	if returnValue == 1 {
		//close(returnChannel)
		return ErrErrorsLimitExceeded
	} else if returnValue == 0 {
		close(jobs)
		return nil
	}

	//wg.Wait()

	return nil
}
