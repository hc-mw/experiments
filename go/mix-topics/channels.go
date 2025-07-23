package main

import (
	"fmt"
	"time"
)

func channelDemo() {
	bch := make(chan string, 2)
	go foo(bch, "1", 0)
	go foo(bch, "2", 0)

	for val := range bch {
		fmt.Println(val)
	}
}

func selectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go foo(ch1, "1", 0)
	go foo(ch2, "2", 1)

	select {
	case res := <-ch1:
		fmt.Println("1:", res)
	case res := <-ch2:
		fmt.Println("2:", res)
	default:
		fmt.Println("default")
	}

	time.Sleep(5 * time.Second)
}

func foo(ch chan<- string, str string, d int) {
	//time.Sleep(time.Duration(d) * time.Second)
	ch <- str
}

func roughlyFair() {
	ch1 := make(chan any)
	close(ch1)
	ch2 := make(chan any)
	close(ch2)

	var n1, n2 int
	for i := 0; i < 100; i++ {
		select {
		case <-ch1:
			n1++
		case <-ch2:
			n2++
		}
	}
	fmt.Printf("n1: %d, n2: %d\n", n1, n2)
}
