// Package task is a simple library that helps to compose complex business logic from smaller steps in a maintainable way.
package task

import (
	"context"
	"sync"
)

// Task is a unit of work that is executable.
type Task func(ctx context.Context) error

// Sequence creates a new task by running multiple tasks in a sequence.
// It stops if any task returns error.
func Sequence(tasks ...Task) Task {
	return func(ctx context.Context) error {
		for _, t := range tasks {
			if err := t(ctx); err != nil {
				return err
			}
		}

		return nil
	}
}

// Concurrence creates a new task by running multiple tasks concurrently.
func Concurrence(tasks ...Task) Task {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		var wg sync.WaitGroup
		errCh := make(chan error, 1)

		for _, t := range tasks {
			tClone := t
			wg.Add(1)
			go func() {
				if err := tClone(ctx); err != nil {
					select {
					case errCh <- err:
					default:
					}
				}
				wg.Done()
			}()
		}

		go func() {
			wg.Wait()
			select {
			case errCh <- nil:
			default:
			}
		}()

		return <-errCh
	}
}

// Exec executes a task or multiple tasks in a sequence. It returns error if there is any.
func Exec(ctx context.Context, tasks ...Task) error {
	return Sequence(tasks...)(ctx)
}
