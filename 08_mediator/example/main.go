package main

import "fmt"

/*------------------ implement chat ------------------*/
func NewChat() Chat {
	return Chat{}
}

type Chat struct {
	users []*User
}

func (c *Chat) sendMessage(msg string, user *User) {
	for i, _ := range c.users {
		if c.users[i].name != user.name {
			c.users[i].receive(msg)
		}
	}
}

func (c *Chat) addUser(user *User) {
	fmt.Printf("%s joined this group\n", user.name)
	c.users = append(c.users, user)
}

/*------------------ implement user ------------------*/
func NewUser(chat *Chat, name string) User {
	return User{name, chat}
}

type User struct {
	name string
	chat *Chat
}

func (u *User) send(msg string) {
	fmt.Println("---")
	fmt.Printf("%s is sending the message: %s\n", u.name, msg)
	u.chat.sendMessage(msg, u)
}

func (u *User) receive(msg string) {
	fmt.Printf("%s received the message: %s\n", u.name, msg)
}

func main() {
	chat := NewChat()

	user1 := NewUser(&chat, "thuongnn7")
	user2 := NewUser(&chat, "tanpv5")
	user3 := NewUser(&chat, "anhbd1")

	chat.addUser(&user1)
	chat.addUser(&user2)
	chat.addUser(&user3)

	user1.send("hi")
	user2.send("test thoi")
	user1.send("hi")
}
