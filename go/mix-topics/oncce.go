package main

import (
	"fmt"
	"sync"
)

func onceDemo() {
	var wg sync.WaitGroup
	wg.Add(1)
	var o sync.Once

	for i := 0; i < 10; i++ {
		go func() {
			o.Do(task)

			wg.Done()
		}()
	}

	wg.Wait()
}

func task() {
	fmt.Println("task done")
}
