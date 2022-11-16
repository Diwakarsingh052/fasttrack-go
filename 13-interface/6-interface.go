package main

import (
	"fmt"
	"log"
)

type student struct {
	name string
}

var data map[int]student

type DB interface {
	Read(id int) (student, error)
	Delete(id int) error
}

type postgres struct {
	con string
}

func (p postgres) Read(id int) (student, error) {
	s, ok := data[id]
	if !ok {

		return student{}, fmt.Errorf("record not found in postgres")
	}
	return s, nil
}

func (p postgres) Delete(id int) error {
	//TODO implement me
	delete(data, id)
	return nil
}

type redis struct {
	con string
}

func (r redis) Read(id int) (student, error) {
	s, ok := data[id]
	if !ok {
		return student{}, fmt.Errorf("record not found in redis")
	}
	return s, nil
}

func (r redis) Delete(id int) error {
	delete(data, id)
	return nil
}

func DoSomethingInDb(d DB) {
	s, err := d.Read(10)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s)
}

func main() {
	data = make(map[int]student)
	data[100] = student{name: "John"}

	p := postgres{con: "root:localhost:postgres"}
	r := redis{con: "localhost:redis"}
	DoSomethingInDb(p)
	DoSomethingInDb(r)

}
