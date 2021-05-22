package task_test

import (
	"context"
	"fmt"

	"github.com/bongnv/task"
)

func makeTask(text string) task.Task {
	return func(ctx context.Context) error {
		fmt.Println(text)
		return nil
	}
}

func ExampleTask() {
	step1 := makeTask("Rinse the rice")
	step2 := makeTask("Use the right ratio of water")
	step3 := makeTask("Bring the water to a boil")
	step4 := makeTask("Maintain a simmer")
	step5 := makeTask("Cook without peeking or stirring")
	step6 := makeTask("Let the rice rest covered")

	err := task.Exec(
		context.Background(),
		step1,
		step2,
		step3,
		step4,
		step5,
		step6,
	)

	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Rinse the rice
	// Use the right ratio of water
	// Bring the water to a boil
	// Maintain a simmer
	// Cook without peeking or stirring
	// Let the rice rest covered
}
