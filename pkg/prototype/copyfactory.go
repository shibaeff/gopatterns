package prototype

type CopyFactory interface {
	SetPrototype(proto Copyable)
	GetCopy() Copyable
}

type humanFactory struct {
	human Copyable
}

func (h *humanFactory) SetPrototype(proto Copyable) {
	h.human = proto
}

func (h *humanFactory) GetCopy() Copyable {
	return h.human.Copy()
}

func NewHumanFactory() CopyFactory {
	return &humanFactory{}
}
