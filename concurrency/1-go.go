// 1 core , chrome, code, firefox
// concurrency means dealing with a lot of things at once and cpu switches between the available apps
// parallelism is doing a lot of things at once

package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("hey") // separate line exec // new goroutine  // not care the order of exec // spinning of the goroutine
	go hello()            // each func call creates a go routine // we are not sure when this go routine gets chance to run, most of the cases at later stage
	fmt.Println("end of main")

	// time.sleep is not a good idea in production
	//time.Sleep(2 * time.Second)// sleeping is a non productive cpu activity so cpu will make the switch to the another go routine

}

func hello() {
	time.Sleep(2 * time.Second)
	fmt.Println("hello from the hello func")
}
