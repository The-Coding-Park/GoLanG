package main

import "fmt"

type profile struct {
	name    string
	address string
	email   string
	contact string
}

func (p profile) getname() string {
	return p.name
}
func (p profile) getaddrss() string {
	return p.address
}
func (p profile) getcontact() map[string]string {
	m := map[string]string{"email": p.email, "contact": p.contact}
	return m
}
func methodDemo() {
	s := profile{
		name:    "gunjan",
		address: "jaipur",
		email:   "gunjan.mtbpaul@gmail.com",
		contact: "8927510637",
	}
	fmt.Println("name: ", s.getname())
	fmt.Println("name: ", s.getaddrss())
	fmt.Println("name: ", s.getcontact())

}

func main() {
	methodDemo()
}
