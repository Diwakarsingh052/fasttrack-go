package main

import "fmt"

type user struct {
	name string
}

func main() {
	u := make(map[int]*user)

	u[101] = &user{name: "john"}

	v, ok := u[101] // value,bool

	//not recommended way of checking if a value is present or not
	if v == nil {
		fmt.Println("using pointers, not there")
	}
	if !ok {
		fmt.Println("not there")
		return
	}
	fmt.Printf("%T\n", v)
	fmt.Println(v)

}
