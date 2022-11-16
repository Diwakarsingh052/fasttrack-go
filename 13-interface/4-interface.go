package main

import (
	"io"
)

type user struct {
	name string
}

func (u user) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (u user) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func main() {

	var u user
	ReadSomething(u)
	WriteSomething(u)
	ReadAndWrite(u)
}

func ReadSomething(r io.Reader)     {}
func WriteSomething(w io.Writer)    {}
func ReadAndWrite(wr io.ReadWriter) {} //  ReadWriter embeds reader and writer interface and form a new interface
