package mediator

type Chat interface {
	AddUser(user User)
	SendMessage(message string, from User)
}

type chat struct {
	users []User
}

func (c *chat) AddUser(user User) {
	c.users = append(c.users, user)
}

func (c *chat) SendMessage(message string, from User) {
	for _, item := range c.users {
		if item != from {
			item.GetMessage(message, from)
		}
	}
}

func NewChat() Chat {
	return &chat{}
}
