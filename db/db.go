package db

import "errors"

var DB string = "mongo"

type Config struct {
	db string // this field is not exported so other packages cannot access this

}

var ErrCon = errors.New("please provide a valid connection")

// NewConfig func is initializing the Config struct and set a db connection
func NewConfig(conn string) (*Config, error) {

	if conn == "" {

		return nil, ErrCon // without pointer // Config{},ErrCon
	}
	//db.Ping()
	//c := Config{db: conn}
	//return &c, nil

	return &Config{db: conn}, nil
}
