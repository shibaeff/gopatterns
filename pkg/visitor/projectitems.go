package visitor

import "errors"

const (
	Test    = "test"
	DB      = "db"
	Cluster = "cluster"
)

// ProjectItem ...
type ProjectItem interface {
	// StartWork triggers dev to work on the item instance
	StartWork(dev Developer)
}

type db struct{}

func (d *db) StartWork(dev Developer) {
	dev.SetupDB()
}

type test struct{}

func (t *test) StartWork(dev Developer) {
	dev.WriteTests()
}

type cluster struct{}

func (c *cluster) StartWork(dev Developer) {
	dev.SetupCluster()
}

func NewProjectItem(which string) (ProjectItem, error) {
	// factory method to create specific project item
	switch which {
	case Test:
		return &test{}, nil
	case DB:
		return &db{}, nil
	case Cluster:
		return &cluster{}, nil
	}
	return nil, errors.New("no such project item")
}
