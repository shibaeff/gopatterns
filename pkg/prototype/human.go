package prototype

import "fmt"

type Copyable interface {
	Copy() Copyable
}

type Human struct {
	age  int
	name string
}

func (h *Human) Copy() Copyable {
	return NewHuman(h.age, h.name)
}

func (h *Human) Print() {
	fmt.Printf("I'm %s.\n", h.name)
}

func NewHuman(age int, name string) *Human {
	return &Human{
		age:  age,
		name: name,
	}
}
