package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool

func gettingReadyForMission() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReady(cond)
	workIntervals := 0

	cond.L.Lock()
	for !ready {
		cond.Wait()
		workIntervals++
	}
	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}
func gettingReady(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}
func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
