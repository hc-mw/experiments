package main

import (
	"fmt"
	"sync"
)

func waitGroupDemo() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo1("hello", &wg)
	wg.Wait()
	//fmt.Println("done")
}

func foo1(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	wg.Done()
}
