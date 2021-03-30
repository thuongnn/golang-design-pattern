package main

import (
	"fmt"
	"strings"
)

func main() {
	user := NewUser("Cristiano Ronaldo")
	customer := NewCustomer(user)

	fmt.Printf("First Name: %s\n", customer.getFirstName())
	fmt.Printf("Last Name: %s\n", customer.getLastName())
}

/*------------------- user define -------------------*/
func NewUser(name string) UserInterface {
	return &user{name}
}

type UserInterface interface {
	getName() string
	setName(name string)
}

type user struct {
	name string
}

func (u *user) getName() string {
	return u.name
}

func (u *user) setName(name string) {
	u.name = name
}

/*------------------- customer define -------------------*/
func NewCustomer(user UserInterface) CustomerInterface {
	split := strings.Split(user.getName(), " ")
	var c customer
	c.setFirstName(split[0])
	c.setLastName(split[1])
	return &c
}

type CustomerInterface interface {
	getFirstName() string
	setFirstName(fistName string)
	getLastName() string
	setLastName(lastName string)
}

type customer struct {
	fistName string
	lastName string
}

func (c *customer) getFirstName() string {
	return c.fistName
}

func (c *customer) setFirstName(fistName string) {
	c.fistName = fistName
}

func (c *customer) getLastName() string {
	return c.lastName
}

func (c *customer) setLastName(lastName string) {
	c.lastName = lastName
}
