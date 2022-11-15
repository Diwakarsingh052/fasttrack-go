package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	abc()
	fmt.Println("main func ends")
}

func abc() {
	defer recoverFunc() // recover func recovers from the panic and stops its further propagation

	var i []int
	i[10] = 100 // panic

	fmt.Println(i)
}

func search(name string) {
	if name == "" {
		panic("name is not found") // not a good idea , panic should be used when something critical is not working
	}
}

func recoverFunc() {
	r := recover() // nil means no panic // otherwise r would be the msg of the panic

	if r != nil {
		fmt.Println("recovered from the panic", r)
		fmt.Println(string(debug.Stack()))
	}
}
