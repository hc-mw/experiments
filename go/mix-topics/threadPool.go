package main

import (
	"fmt"
	"sync"
)

type Job func()

type Pool struct {
	workerQueue chan Job
	wg          sync.WaitGroup
}

func (p *Pool) AddJob(job Job) {
	p.workerQueue <- job
}

func (p *Pool) Wait() {
	close(p.workerQueue)
	p.wg.Wait()
}

func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workerQueue: make(chan Job),
	}
	// add job
	pool.wg.Add(workerCount)
	// spin up threads
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			defer pool.wg.Done()
			// wait for a job to come
			for job := range pool.workerQueue {
				fmt.Printf("%dth Worker doing job\n", id)
				job()
			}
		}(i)
	}

	return pool
}
