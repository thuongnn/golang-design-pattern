package main

import "fmt"

type UserService interface {
	load()
	insert()
}

type UserServiceImpl struct {
	name string
}

func (us *UserServiceImpl) load() {
	fmt.Printf("%s loaded\n", us.name)
}

func (us *UserServiceImpl) insert() {
	fmt.Printf("%s inserted\n", us.name)
}

type UserServiceProxy struct {
	role        string
	userService *UserServiceImpl
}

func NewUserServiceProxy(name, role string) *UserServiceProxy {
	return &UserServiceProxy{
		role:        role,
		userService: &UserServiceImpl{name},
	}
}

func (us *UserServiceProxy) load() {
	us.userService.load()
}

func (us *UserServiceProxy) insert() {
	if us.isAdmin() {
		us.userService.insert()
	} else {
		fmt.Println("Access denied")
	}
}

func (us *UserServiceProxy) isAdmin() bool {
	return us.role == "admin"
}

func main() {
	admin := NewUserServiceProxy("thuongnn", "admin")
	admin.load()
	admin.insert()

	client := NewUserServiceProxy("customer", "client")
	client.load()
	client.insert()
}