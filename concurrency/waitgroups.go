package main

import (
	"fmt"
	"sync"
	"time"
)

//var wg = sync.WaitGroup{}

func main() {
	wg := &sync.WaitGroup{}
	//wg.Add(10) // counter to add number of goroutines that we are starting or spinning up
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg) // 10 goroutines
	}
	//wg.Add(1) // it is a deadlock as we are running just 10 goroutines not 11

	wg.Add(1)
	go calc(10, 20, wg) // from goroutines, we can't return values // so we create a blanket func and do our actual work there

	fmt.Println("some other task is scheduled")

	wg.Wait() // it will wait until counter resets to zero

}

func work(i int, wg *sync.WaitGroup) {
	//panic("error") // panic reveals go routines id
	defer wg.Done() // decrements the counter by one
	fmt.Println("I am doing some work", i)
}

func calc(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	s := add(a, b)
	fmt.Println("calc job is finished", s)

}

func add(a, b int) int {
	return a + b
}
