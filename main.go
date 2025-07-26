package main

import "fmt"

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
	newTask := MakeTask(0)

	for range 10 {
		newTask.IncrementValue()
	}

	for taskPriority := range LOWEST_PRIORITY {
		for t, task := range TaskQueue[taskPriority] {
			fmt.Printf("Priority: %v, Task %v: Value: %v\n", taskPriority, t, task.Value)
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
