package main

func main() {
	//A send on an unbuffered channel can proceed if a receiver is ready.
	// A send on a buffered channel can proceed if there is room in the buffer
	ch := make(chan int, 1)
	//<-ch // it will block here as there is no value to recv
	//ch <- 20
	ch <- 20 // we can proceed if here as we have room in the buffer
	//fmt.Println(<-ch) // recv will take the value out of the buffer, and it will empty the space taken up by the value so we can push more new values
	//ch <- 30
	//fmt.Println(<-ch)

}
