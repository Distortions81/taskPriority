package main

import "fmt"

var TaskLimit int = 8

var TaskQueue []*taskData

type taskData struct {
	Value    int
	Priority int
	Do       func()
}

func (t *taskData) IncrementValue() {
	t.Value++
}

func main() {
	newTask := &taskData{
		Priority: 0,
	}
	TaskQueue = append(TaskQueue, newTask)

	for range 10 {
		newTask.IncrementValue()
	}

	for t, task := range TaskQueue {
		fmt.Printf("Task %v: %v\n", t, task.Value)
	}
}
