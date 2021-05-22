package task

import (
	"context"
	"errors"
	"testing"
	"time"
)

func createTask(err error) (Task, *bool) {
	called := false
	return func(ctx context.Context) error {
		called = true
		return err
	}, &called
}

func Test_Sequence_happy(t *testing.T) {
	task1, task1Called := createTask(nil)
	task2, task2Called := createTask(nil)

	err := Exec(context.Background(), Sequence(task1, task2))
	if err != nil {
		t.Log("test failed with error:", err)
		t.FailNow()
	}

	if !*task1Called {
		t.Log("test failed as task1 is not called")
		t.FailNow()
	}

	if !*task2Called {
		t.Log("test failed as task2 is not called")
		t.FailNow()
	}
}

func Test_Sequence_error(t *testing.T) {
	task1, task1Called := createTask(errors.New("random error"))
	task2, task2Called := createTask(nil)

	err := Exec(context.Background(), Sequence(task1, task2))
	if err == nil {
		t.Log("test failed as there is no error returned")
		t.FailNow()
	}

	if !*task1Called {
		t.Log("test failed as task1 is not called")
		t.FailNow()
	}

	if *task2Called {
		t.Log("test failed as task2 should not be called")
		t.FailNow()
	}
}

func Test_Concurrence_happy(t *testing.T) {
	task1, task1Called := createTask(nil)
	task2, task2Called := createTask(nil)

	err := Exec(context.Background(), Concurrence(task1, task2))
	if err != nil {
		t.Log("test failed with error:", err)
		t.FailNow()
	}

	if !*task1Called {
		t.Log("test failed as task1 is not called")
		t.FailNow()
	}

	if !*task2Called {
		t.Log("test failed as task2 is not called")
		t.FailNow()
	}
}

func Test_Concurrence_error(t *testing.T) {
	task1, task1Called := createTask(errors.New("random error"))
	task2, _ := createTask(nil)

	err := Exec(context.Background(), Concurrence(task1, task2))
	if err == nil {
		t.Log("test failed as there is no error returned")
		t.FailNow()
	}

	if !*task1Called {
		t.Log("test failed as task1 is not called")
		t.FailNow()
	}
}

func Test_Concurrence_both_error(t *testing.T) {
	task1, _ := createTask(errors.New("random error"))
	task2, _ := createTask(errors.New("another error"))

	err := Exec(context.Background(), Concurrence(task1, task2))
	if err == nil {
		t.Log("test failed as there is no error returned")
		t.FailNow()
	}
}

func Test_Concurrence_context_cancelled(t *testing.T) {
	task1, task1Called := createTask(errors.New("random error"))

	task2Done := make(chan struct{})
	task2 := func(ctx context.Context) error {
		<-ctx.Done()
		close(task2Done)
		return nil
	}

	err := Exec(context.Background(), Concurrence(task1, task2))
	if err == nil {
		t.Log("test failed as there is no error returned")
		t.FailNow()
	}

	if !*task1Called {
		t.Log("test failed as task1 is not called")
		t.FailNow()
	}

	select {
	case <-task2Done:
	case <-time.After(time.Second):
		t.Log("test timed out")
		t.FailNow()
	}
}
