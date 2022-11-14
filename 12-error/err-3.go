package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("not able to find in root directory")

func main() {
	_, err := openFile("abc")
	if errors.Is(err, ErrFileNotFound) {
		fmt.Println("our custom error  is present in the chain", "let's create that file")
		return
	}
	fmt.Println("our custom error is not present in the chain")

}

func openFile(fileName string) (*os.File, error) { // *os.File denotes that the value returned is a pointer
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		// %w verb wraps two errors together // go 1.13 introduced wrapping and unwrapping
		return nil, fmt.Errorf("%v %w", err, ErrFileNotFound) // fmt.Errorf returns an error msg, we are trying to add some extra info to our error value
	}
	return f, nil
}
