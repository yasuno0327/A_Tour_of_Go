package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	name             string
	process_millisec int
}

func (t *Task) ProcessTask(wg *sync.WaitGroup) {
	fmt.Printf("Start %s\n", t.name)
	for i := 0; i < t.process_millisec; i++ {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Printf("Finished %s\n", t.name)
	wg.Done()
}

var TaskA = Task{name: "A", process_millisec: 1000}
var TaskB = Task{name: "B", process_millisec: 2000}
var TaskC = Task{name: "C", process_millisec: 3000}
var Tasks []Task = []Task{TaskA, TaskB}

func GoShori() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	for _, task := range Tasks {
		wg1.Add(1)
		go func(task Task) {
			go task.ProcessTask(&wg1)
		}(task)
	}
	wg1.Wait()

	wg2.Add(1)
	go TaskC.ProcessTask(&wg2)
	wg2.Wait()
}
