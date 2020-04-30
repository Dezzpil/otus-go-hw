package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	var errorsCount int32
	wg := &sync.WaitGroup{}
	for len(tasks) > 0 {
		for _, task := range tasks[:N] {
			wg.Add(1)

			go func(t Task) {
				if err := t(); err != nil {
					if M >= 0 { // оптимизируем, что при -1 не будем считать
						atomic.AddInt32(&errorsCount, 1)
					}
				}
				wg.Done()
			}(task)
		}
		wg.Wait()
		if M >= 0 && int(errorsCount) > M {
			return ErrErrorsLimitExceeded
		}
		tasks = tasks[N:]
	}
	return nil
}
