package main

import (
	"patterns/pkg/visitor"
)

const (
	Test    = "test"
	DB      = "db"
	Cluster = "cluster"
)

func main() {
	test, _ := visitor.NewProjectItem(Test)
	db, _ := visitor.NewProjectItem(DB)
	cluster, _ := visitor.NewProjectItem(Cluster)
	project := visitor.NewProject(test, db, cluster)

	junior := visitor.NewDeveloper(false)
	senior := visitor.NewDeveloper(true)

	project.StartWork(junior)
	project.StartWork(senior)
}
