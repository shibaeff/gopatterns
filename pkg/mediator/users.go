package mediator

import "fmt"

const (
	getMessageString = "User %s got message from %s: %s\n"
)

type User interface {
	GetMessage(message string, from User)
	SendMessage(message string)
	GetName() string
}

type user struct {
	name string
	chat Chat
}

func (u *user) GetMessage(message string, from User) {
	fmt.Printf(getMessageString, u.name, from.GetName(), message)
}

func (u *user) SendMessage(message string) {
	u.chat.SendMessage(message, u)
}

func (u *user) GetName() string {
	return u.name
}

func NewUser(name string, chat Chat) User {
	return &user{
		name: name,
		chat: chat,
	}
}

type admin struct {
	name string
	chat Chat
}

func (a *admin) GetMessage(message string, from User) {
	fmt.Printf(getMessageString, a.GetName(), from, message)
}

func (a *admin) SendMessage(message string) {
	a.chat.SendMessage(message, a)
}

func (a *admin) GetName() string {
	return a.name
}

func NewAdmin(chat Chat) User {
	return &user{
		name: "admin",
		chat: chat,
	}
}
