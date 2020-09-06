package main

import "patterns/pkg/mediator"

func main() {
	chat := mediator.NewChat()
	admin := mediator.NewAdmin(chat)
	chat.AddUser(admin)
	user1 := mediator.NewUser("Pasha", chat)
	chat.AddUser(user1)
	user2 := mediator.NewUser("Sasha", chat)
	chat.AddUser(user2)

	chat.SendMessage("Hello", user1)
	chat.SendMessage("Hi", user2)
	chat.SendMessage("Privet", admin)
}
