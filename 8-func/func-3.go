package main

import (
	"fmt"
)

type sum func(x, y int)

func main() {

	add := func(a, b int) {
		fmt.Println(a + b)
	}
	sub := func(a, b int) {
		fmt.Println(a - b)
	}
	add(70, 80)
	calc(add)
	calc(sub)
	calc2(add)
	//	http.HandleFunc() // /home , homeHandleFunc
	//http.HandlerFunc()
}

func calc(sum func(x, y int)) {
	sum(10, 40)
}

func calc2(s sum) {
	s(90, 90)
}
