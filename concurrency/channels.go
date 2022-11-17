package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
//A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
// unbuffered channel gives you a guarantee that the work would be always received by someone

// A send on a buffered channel can proceed if there is room in the buffer
var wg = &sync.WaitGroup{}

//channel -> unbuffered chan, buffered chan

func main() {
	c := make(chan int) // unbuffered chan

	wg.Add(3)
	go addNum(10, 20, c)
	go mult(2, 2, c)
	go sub(9, 6, c)
	//go calcAll(c)

	x, y, z := <-c, <-c, <-c // recv from the channel // it is a blocking call // this line will block until we recv all the values // it blocks the go routine where it is being used
	fmt.Println(x, y, z)

	// please create a new channel when we have a new series of task or different task than other go routines
	//go workOnJson()

	wg.Wait()
}

func addNum(a, b int, c chan int) {
	defer wg.Done()
	sum := a + b
	// in case of an unbuffered chan , receiver must be ready otherwise send will block
	c <- sum // send operation signal on the channel  // signaling with data
}

func sub(a, b int, c chan int) {
	time.Sleep(2 * time.Second)
	defer wg.Done()
	sum := a - b
	c <- sum // send
}

func mult(a, b int, c chan int) {
	defer wg.Done()
	sum := a * b
	c <- sum // send
}

//func calcAll(c chan int) {
//	defer wg.Done()
//	x, y, z := <-c, <-c, <-c // recv from the channel
//	fmt.Println(x, y, z)
//}
