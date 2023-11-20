package taskscheduler

import (
	"fmt"
	"sync/atomic"
	"time"
)

type TaskScheduler struct {
	Name   string
	Timers map[int64]*time.Timer
	Count  atomic.Int64
}

func New() *TaskScheduler {
	return &TaskScheduler{Name: "Generic Scheduler", Timers: map[int64]*time.Timer{}}
}

func (ts *TaskScheduler) AddJob(f func(), duration time.Duration) {
	id := ts.Count.Load()

	// Add cleanup function
	newFunc := func() {
		f()
		delete(ts.Timers, id)
	}
	timer := time.AfterFunc(duration, newFunc)
	ts.Timers[id] = timer
	ts.Count.Add(1)
	fmt.Println(ts.Timers)
}
