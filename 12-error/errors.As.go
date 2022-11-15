package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// this method is compulsory to format your err msg // and all the struct field should be used inside the error msg
func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input: " + q.Input + " " + q.Err.Error() //error msg format

}

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

func SearchSomething(s string) error {
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   ErrNotFound,
	}

}

func SearchName(name string) error {
	return &QueryError{
		Func:  "SearchName",
		Input: name,
		Err:   ErrMismatch,
	}
}

func main() {
	var q *QueryError // nil by default

	err := SearchSomething("abc")

	if errors.As(err, &q) { // reference imp // it checks whether struct is inside the error chain or not
		fmt.Println("true", q.Func)
		return
	}
	fmt.Println("not")

	//err = SearchName("xyz")
	//fmt.Println(err)
}
