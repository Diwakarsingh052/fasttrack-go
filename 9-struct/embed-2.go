package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

func (u user) show() {
	fmt.Printf("%+v\n", u)
}

type admin struct {
	user  // embedding // anonymous field because we are not giving any name to it // anonymous field name would be assigned according to the type name
	email string
	role  []string
}

func (a admin) show() {
	fmt.Printf("%+v\n", a)
}

func main() {
	a := admin{
		user: user{
			name:  "user",
			email: "user@email.com",
		},
		email: "admin@email.com",
		role:  nil,
	}

	fmt.Println(a.user.email)
	a.user.show()
}
