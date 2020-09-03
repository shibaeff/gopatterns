package strategy

import "errors"

type Developer interface {
	SetActivity(activity Activity)
	ExecuteActivity() error
	GetName() string
}

type developer struct {
	name     string
	activity Activity
}

func (d *developer) SetActivity(activity Activity) {
	d.activity = activity
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
		activity: nil,
	}
	return
}
