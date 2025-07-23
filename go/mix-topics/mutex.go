package main

import (
	"fmt"
	"sync"
)

var (
	lock  sync.Mutex
	count int
	wg    sync.WaitGroup
)

func mutexDemo() {
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go inc()
	}

	wg.Wait()
	fmt.Println(count)
}

func inc() {
	lock.Lock()
	count++
	lock.Unlock()
	wg.Done()
}
