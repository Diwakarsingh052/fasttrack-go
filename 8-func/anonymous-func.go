package main

import "fmt"

func main() {

	i := func(a, b int) int {

		return a + b
	}(10, 20)

	fmt.Println(i)

}
