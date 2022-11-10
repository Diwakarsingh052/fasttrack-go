package main

import (
	"fmt"
	"log"
)

func main() {
	msg, ok := hello("ajay", 23)
	if !ok {
		log.Println("process failed", msg)
		return
	}
	fmt.Println(msg)

}

func hello(name string, age int) (string, bool) {
	if name == "" {
		return "please provide a name", false // it will stop the current func and return values
	}
	if age == 0 {
		return "please provide your age", false
	}
	fmt.Println(name, age)
	return "success", true
}
