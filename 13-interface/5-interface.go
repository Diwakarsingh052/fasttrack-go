package main

import (
	"fmt"
)

type Runner interface {
	Run()
}
type Walker interface {
	Walk()
}

type WalkRunner interface {
	Runner
	Walker
}

type Human struct{ name string }

func (h Human) Walk() {
}
func (h Human) Run() {
}

type Robo struct{ name string }

func (r Robo) Run() {}

func main() {

	h := Human{name: "John"}
	robot := Robo{name: "r1"}
	_ = h

	var r Runner = robot

	r.Run()

	//type assertion
	i, ok := r.(Human) // type assertion // check whether a type exists in the interface or not // if it is there than that struct would be returned

	if !ok {
		fmt.Println("you are not human") // ok would be true if the human struct is present otherwise it would be false
		//return
	}
	fmt.Println(i)
	i.Walk()

}
