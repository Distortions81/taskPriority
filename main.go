package main

import (
	"fmt"
	"time"
)

var TaskLimit int = 8

const LOWEST_PRIORITY = 100

var TaskQueue [LOWEST_PRIORITY][]*taskData

type taskData struct {
	Value    int
	Priority int
}

func (t *taskData) IncrementValue() {
	t.Value++
}

func main() {
	MakeTask(0)
	MakeTask(0)
	MakeTask(0)

	MakeTask(1)

	start := time.Now()
	for {
		for taskPriority := range LOWEST_PRIORITY {
			for _, task := range TaskQueue[taskPriority] {
				task.IncrementValue()
			}
		}
		if time.Since(start) > time.Second {
			break
		}
	}

	for taskPriority := range LOWEST_PRIORITY {
		for t, task := range TaskQueue[taskPriority] {
			fmt.Printf("Priority: %v, Task %v: Value: %v\n", taskPriority, t+1, task.Value)
		}
	}
}

func MakeTask(priority int) *taskData {
	newTask := &taskData{
		Priority: priority,
	}
	TaskQueue[priority] = append(TaskQueue[priority], newTask)
	return newTask
}
