package main

import (
	"fmt"
	"strconv"
	"sync"
)

// restrict two things to exec at a time

var wg = &sync.WaitGroup{}

func main() {

	const work = 10
	const cap = 2
	ch := make(chan string, work)
	sem := make(chan bool, cap)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range ch { // recv until the channel is not closed // on the closed chan we can recv value but we cannot send to it
			fmt.Println(v)
		}
	}()

	go func() {
		defer wg.Done()

		for e := 1; e <= work; e++ {
			//this go func is the sender
			go func(e int) {
				sem <- true // this line will block if we don't have room in the buffer
				ch <- "task" + strconv.Itoa(e)
			}(e)
			//time.Sleep(1 * time.Second)
			<-sem // freeing the space in the buffer
		}
		close(ch) // close channels from where you are sending the data
		//ch <- "hello" // send on closed channel is not allowed
	}()
	wg.Wait()
}
