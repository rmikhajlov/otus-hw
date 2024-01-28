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
	finishedBecauseErrors := false

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mutex.Lock()
				if finishedBecauseErrors {
					mutex.Unlock()
					return
				}
				mutex.Unlock()
				select {
				case job := <-jobs:
					if job == nil {
						s.Do(func() {
							returnChannel <- 0
						})
						return
					}
					mutex.Lock()
					if m == 0 {
						finishedBecauseErrors = true
						mutex.Unlock()
						return
					}
					mutex.Unlock()
					if err := job(); err != nil {
						mutex.Lock()
						m--
						if m == 0 {
							finishedBecauseErrors = true
							mutex.Unlock()
							return
						}
						mutex.Unlock()
					}
				default:
					s.Do(func() {
						returnChannel <- 0
					})
					return
				}
			}
		}()
	}

	wg.Wait()

	if finishedBecauseErrors {
		return ErrErrorsLimitExceeded
	}

	return nil
}
