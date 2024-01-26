package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	jobs := make(chan Task, len(tasks))
	returnChannel := make(chan int, 1)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	s := sync.Once{}

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case job := <-jobs:
					if job == nil {
						//wg.Done()
						s.Do(func() {
							returnChannel <- 0
						})
						return
					}
					if err := job(); err != nil {
						mutex.Lock()
						m--
						if m == 0 {
							mutex.Unlock()
							//wg.Done()
							returnChannel <- 1
							return
						}
						mutex.Unlock()
					}
				default:
					//wg.Done()
					s.Do(func() {
						returnChannel <- 0
					})
					return
				}
			}
		}()
	}

	returnValue := <-returnChannel

	if returnValue == 1 {
		return ErrErrorsLimitExceeded
	} else if returnValue == 0 {
		wg.Wait()
		return nil
	}

	return nil
}
