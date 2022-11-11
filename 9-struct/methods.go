package main

import (
	"fmt"
)

type student struct {
	name string
}

func (s student) print() { // value from the main func of the user var would be copied here in the receiver
	fmt.Println(s.name)
}

func (s student) update(name string) {
	s.name = name

}

func main() {

	s := student{name: "Bob"}
	s.print()
	s.update("John")
	s.print()
	//os.OpenFile()

}
