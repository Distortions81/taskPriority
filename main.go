package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/remeh/sizedwaitgroup"
)

var TaskLimit int = 8

const LOWEST_PRIORITY = 100

var TaskQueue [LOWEST_PRIORITY][]*taskData
var CompletedQueue []*taskData

type taskData struct {
	Value    int
	Priority int
}

func (t *taskData) IncrementValue() {
	t.Value++
}

var TimeBudget [LOWEST_PRIORITY]time.Duration

func InitTimeBudgets() {
	for i := 0; i < LOWEST_PRIORITY; i++ {
		if i < 10 {
			TimeBudget[i] = time.Duration(50-(i*5)) * time.Millisecond // e.g. 50ms, 45ms, ...
			//fmt.Printf("%v: %v\n", i, TimeBudget[i])
		}
	}
}

func main() {

	InitTimeBudgets()

	for i := range 9 {
		for range 50 {
			MakeTask(i)
		}
	}
	RunScheduler()
}

func RunScheduler() {
	rounds := 0
	for {
		rounds++

		allEmpty := true
		wg := sizedwaitgroup.New(runtime.NumCPU())

		for priority := 0; priority < LOWEST_PRIORITY; priority++ {
			queue := TaskQueue[priority]
			qLen := len(queue)

			if qLen == 0 {
				//No tasks at this priority
				continue
			}
			allEmpty = false

			sliceDeadline := time.Now().Add(TimeBudget[priority])

			for qLen > 0 && time.Now().Before(sliceDeadline) {
				task := TaskQueue[priority][0]
				TaskQueue[priority] = TaskQueue[priority][1:]
				CompletedQueue = append(CompletedQueue, task)
				qLen--
				wg.Add()
				go func(task *taskData) {
					defer wg.Done()
					time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)
					task.IncrementValue()
				}(task)
			}
		}
		wg.Wait()

		if allEmpty {
			break
		}
	}

	for t, task := range CompletedQueue {
		fmt.Printf("Pri: %v, Task: %v, Value: %v\n", task.Priority, t, task.Value)
	}
	fmt.Printf("Total rounds: %d\n", rounds)

}

func MakeTask(priority int) *taskData {
	newTask := &taskData{
		Priority: priority,
		Value:    0,
	}
	TaskQueue[priority] = append(TaskQueue[priority], newTask)
	return newTask
}
