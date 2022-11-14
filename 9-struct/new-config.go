package main

import (
	"errors"
	"log"
	"os"
)

var DB string = "mongo"

type Config struct {
	db string
}

var ErrCon = errors.New("please provide a valid connection")

// NewConfig func is initializing the Config struct
func NewConfig(conn string) (*Config, error) {

	if conn == "" {

		return nil, ErrCon // without pointer // Config{},ErrCon
	}
	//db.Ping()
	//c := Config{db: conn}
	//return &c, nil

	return &Config{db: conn}, nil
}

func main() {

	log.Println("hello ")
	l := log.New(os.Stdout, "sales-app ", log.LstdFlags)
	l.Println("welcome to this app")

	//ioutil.ReadAll()
	//ioutil.ReadFile()
}
