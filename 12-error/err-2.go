package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("not able to find in root directory")

func main() {
	_, err := openFile("abc")
	//if err contains ErrFileNotFound { this is not possible with %v verb as it meges the both error value together
	//
	//}
	_ = err
}

func openFile(fileName string) (*os.File, error) { // *os.File denotes that the value returned is a pointer
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		// %v verb merges two errors together // go 1.13 this was replaced with wrapping
		return nil, fmt.Errorf("%v %v", err, ErrFileNotFound) // fmt.Errorf returns an error msg, we are trying to add some extra info to our error value
	}
	return f, nil
}
