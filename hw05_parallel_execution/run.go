package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	jobs := make(chan Task, len(tasks))
	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var errorsCount int

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				if err := job(); err != nil {
					mutex.Lock()
					errorsCount++
					if errorsCount >= m {
						mutex.Unlock()
						return
					}
					mutex.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	if errorsCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
