package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	u    *user // normal field
	role []string
}

func (a admin) show() {
	fmt.Printf("%+v\n", a)
}

func main() {
	var u *user = &user{} // if we are storing struct instance in pointer than we have to pass it's address using &
	_ = u
	a := admin{
		u: &user{
			name:  "user",
			email: "user@email.com",
		},
		role: nil,
	}

	a.u.updateEmail("user123@email.com")

}
