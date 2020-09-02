package visitor

// Project ...
type Project interface {
	// StartWork triggers the Developer on work on ALL project items
	StartWork(dev Developer)
}

type project struct {
	items []ProjectItem
}

func (p *project) StartWork(dev Developer) {
	// iterate over items to trigger work
	for _, item := range p.items {
		item.StartWork(dev)
	}
}

func NewProject(items ...ProjectItem) Project {
	p := project{}
	// iterate over items and add the mto them project
	for _, item := range items {
		p.items = append(p.items, item)
	}
	return &p
}
