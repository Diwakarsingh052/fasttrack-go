package main

import (
	"fmt"
)

// https://go.dev/doc/faq#methods_on_values_or_pointers
type movies struct {
	title string
}

// any changes that we would do in update method , then those changes would be reflected back to the caller variable
func (m *movies) update(name string) {
	m.title = name // *(dereference) is not required // automatically done
}

func (m *movies) print() { // keep receiver names consistent
	fmt.Println(m)
}

type user struct {
	name string
}

func (u *user) update(name string) {

}

func main() {
	var m movies
	m.update("xyz")
	fmt.Println(m)

}
