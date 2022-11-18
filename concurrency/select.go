package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	wg.Add(3)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("hello this will run later , and not going to affect our program when it proceed further")
	//}()

	go func() {
		defer wg.Done()

		time.Sleep(3 * time.Second)
		c1 <- "one" // send
	}()

	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	go func() {
		defer wg.Done()

		c3 <- "three"
	}()
	//
	//fmt.Println(<-c1)
	//fmt.Println(<-c2) // uncomment to see the problem
	//fmt.Println(<-c3)

	for i := 1; i <= 3; i++ {

		select {
		// whichever case is not blocking exec that first
		//whichever case is ready first exec that.
		case x := <-c1: // recv over the channel
			fmt.Println("send the result in the further pipeline", x)
		case y := <-c2:
			fmt.Println(y)
		case z := <-c3:
			fmt.Println(z)

		}
	}

	fmt.Println("end of the main")
	wg.Wait()

}
