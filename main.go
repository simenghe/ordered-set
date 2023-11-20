package main

import (
	"fmt"
	"orderedset/taskscheduler"
	"time"
)

func main() {
	ts := taskscheduler.New()
	ts.AddJob(func() { fmt.Println("Hello in 1 second") }, time.Second)
	ts.AddJob(func() { fmt.Println("Hello in 2 second") }, 2*time.Second)

	time.Sleep(10 * time.Second)
	fmt.Println(ts.Timers)
}
