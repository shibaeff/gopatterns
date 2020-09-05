package strategy

import "fmt"

const (
	activityString = "Developer %s is now %s\n"

	runningString = "running"
	codingString  = "coding"
)

type Activity interface {
	Execute(dev Developer)
	Next() Activity
}

type run struct{}

func (r *run) Execute(dev Developer) {
	fmt.Printf(activityString, dev.GetName(), runningString)
}

func (r *run) Next() Activity {
	return NewCoding()
}

func NewRun() Activity {
	return &run{}
}

type coding struct{}

func (c *coding) Execute(dev Developer) {
	fmt.Printf(activityString, dev.GetName(), codingString)
}

func (c *coding) Next() Activity {
	return NewRun()
}

func NewCoding() Activity {
	return &coding{}
}
