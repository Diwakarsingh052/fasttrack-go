// Package model  is the only place that interacts with our Db
package model

import "errors"

type User struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
}

var ErrUserNotFound = errors.New("user id not found")

var users = map[uint64]User{
	123: {
		FName: "Ajay",
		LName: "Sharma",
		Email: "ajay@email.com",
	},
}

func FetchUser(id uint64) (User, error) {
	u, ok := users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
