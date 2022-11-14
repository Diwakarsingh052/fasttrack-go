package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

// prefix your err variables with Err word

var ErrRecordNotFound = errors.New("record not found with the id provided")

func main() {
	//os.ErrClosed
	//sql.Err

	data, err := FetchRecord(100)
	if err != nil {
		log.Println(err) // log means you have handled your errors, don't do duplicate logging and once an error is handled then don't propagate that error to further chain
		return
	}
	fmt.Println(data)
}

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	return name, nil

}
