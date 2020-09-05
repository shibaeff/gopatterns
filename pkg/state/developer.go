package strategy

import "errors"

type Developer interface {
	ChangeState()
	ExecuteActivity() error
	GetName() string
}

type developer struct {
	name     string
	activity Activity
}

func (d *developer) ChangeState() {
	d.activity = d.activity.Next() // make sure that all the state make up a perfect state automata!
}

func (d *developer) ExecuteActivity() (err error) {
	if d.activity == nil {
		return errors.New("nil activity")
	}
	d.activity.Execute(d)
	return
}

func (d *developer) GetName() string {
	return d.name
}

func NewDeveloper(name string) (d Developer) {
	d = &developer{
		name:     name,
		activity: NewCoding(),
	}
	return
}
