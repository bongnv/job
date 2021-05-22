# task

[![CI](https://github.com/bongnv/task/actions/workflows/ci.yml/badge.svg)](https://github.com/bongnv/task/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/bongnv/task/branch/main/graph/badge.svg?token=OpOoecFx9h)](https://codecov.io/gh/bongnv/task)
[![Go Report Card](https://goreportcard.com/badge/github.com/bongnv/task)](https://goreportcard.com/report/github.com/bongnv/task)

`task` is a simple library that helps to compose complex business logic from smaller steps in a maintainable way.

## Quick Start

1. Make sure [Go](https://golang.org/) is installed and import `task` package to your project via the below command:
```sh
$ go get -u github.com/bongnv/task
```

2. Import it to your code:

```go
import "github.com/bongnv/task"
```

3. Use `task.Exec` to run multiple tasks in sequence:

```go
	step1 := makeTask("Rinse the rice")
	step2 := makeTask("Use the right ratio of water")
	step3 := makeTask("Bring the water to a boil")
	step4 := makeTask("Maintain a simmer")
	step5 := makeTask("Cook without peeking or stirring")
	step6 := makeTask("Let the rice rest covered")

	err := task.Exec(
		ctx,
		step1,
		step2,
		step3,
		step4,
		step5,
		step6,
	)
```

## Usages

There are two ways to compose tasks together:

1. `task.Sequence` creates a new task by running multiple tasks in a sequence. It stops and returns error if any task returns error.

```go
	composedTask := task.Sequence(
		step1,
		step2,
		step3,
	)
```
2. `task.Concurence` creates a new task by running multiple tasks concurrently. It stops, cancels the provided context and returns error if any task return error.

```go
	composedTask := task.Concurrence(
		step1,
		step2,
		step3,
	)
```