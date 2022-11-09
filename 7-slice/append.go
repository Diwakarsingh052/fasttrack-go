package main

import "fmt"

func main() {
	var a []int
	b := []int{90, 809}

	a = append(a, 10, 20, 30, 40)
	a = append(a, b...) // we are unpacking the b slice

	fmt.Println(a)
}
