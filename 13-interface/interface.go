package main

import (
	"fmt"
)

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike
// Bigger the interface weaker the abstraction // Rob Pike

// interfaces are abstract type // it does not store anything of their own
type reader interface {
	read(b []byte) (int, error)
	//abc() // all the methods should be impl over the struct to use the interface
}

type file struct {
	name string
}

// read method signature has to be exactly the same as defined in the interface to impl that
func (f file) read(b []byte) (int, error) {
	fmt.Println("inside file read")
	s := "hello go devs"
	copy(b, s)
	return len(b), nil
}

func (f file) print() {
	fmt.Println("i am a method of file struct")
}

type jsonObject struct {
	data string
}

func (j jsonObject) read(b []byte) (int, error) {
	s := `{name:"abc"}`
	fmt.Println("inside json read")
	copy(b, s)
	return len(s), nil
}

// fetch is a polymorphic func
// fetch() will accept any type of value which implements reader interface
func fetch(r reader) {

	fmt.Printf("%T\n", r)

	data := make([]byte, 50)
	x, _ := r.read(data)
	//r.print() // we can only call methods part of the interface signature

	fmt.Println(x)
	fmt.Println(string(data))
}

func main() {
	f := file{name: "abc.txt"} // concrete data
	j := jsonObject{data: "any json"}

	//fetch(f)
	//fetch(j)

	//var r reader // nil is the default value of the interface
	//s := make([]byte, 90)
	//x, err := r.read(s) // we can't call the read as no struct or concrete value was passed to it.

	var r reader = f //passing in file var as it impls the interface
	s := make([]byte, 90)
	r.read(s)
	fmt.Println(string(s))

	fmt.Printf("%T\n", r)
	r = j     // passing in jsonObject as it impls the interface
	r.read(s) // jsonObject read method would be called
	fmt.Println(string(s))
	fmt.Printf("%T\n", r)

}

//func FetchDataFile(f file) {
//
//}
//
//func FetchDataJson(j jsonObject) {
//}
