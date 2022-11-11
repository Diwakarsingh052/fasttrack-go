package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	//do work with your file

}
