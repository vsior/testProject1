package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type runFunc func(int) int

type Job struct {
	id int
	runFunc
}

type Data struct {
	Job
	res int
}

type WorkerPool struct {
	waitGroupPool sync.WaitGroup
	size          int
	jobs          chan Job
	data          chan Data
	runFunc
}

func NewPoolWork(size int, run runFunc) WorkerPool {
	return WorkerPool{
		size:    size,
		runFunc: run,
	}
}

func (w *WorkerPool) SetNewSize(size int) {
	fmt.Println("old size", w.size, "new size", size)
	w.size = size
	fmt.Println(strings.Repeat("*", 10))
}

func (w *WorkerPool) work() {
	for f := range w.jobs {
		w.data <- Data{f, f.runFunc(f.id)}
		time.Sleep(time.Second)
	}

	w.waitGroupPool.Done()
}

func (w *WorkerPool) createAllWork() {
	for i := 0; i < w.size; i++ {
		c := Job{runFunc: w.runFunc, id: i}
		w.jobs <- c
	}
	close(w.jobs)
}

func (w *WorkerPool) readAll() {
	for d := range w.data {
		fmt.Printf("%+v\n", d)
	}

	fmt.Println(strings.Repeat("*", 10))
}

func (w *WorkerPool) RunWorker() {
	w.jobs = make(chan Job, w.size)
	w.data = make(chan Data, w.size)

	w.createAllWork()

	for i := 0; i < w.size; i++ {
		w.waitGroupPool.Add(1)
		go w.work()
	}

	w.waitGroupPool.Wait()
	close(w.data)

	w.readAll()
}
