package main

import (
	"patterns/pkg/strategy"
)

func main() {
	developer := strategy.NewDeveloper("Pasha")
	developer.SetActivity(strategy.NewCoding())
	developer.ExecuteActivity()
	developer.SetActivity(strategy.NewRun())
	developer.ExecuteActivity()
}
