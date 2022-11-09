package main

import "fmt"

func main() {

	var a []int
	a[0] = 100 // this will cause panic as length is not available to store the value
	if a == nil {
		fmt.Println("it is a nil slice")
	}

	fmt.Printf("%#v", a)

}
