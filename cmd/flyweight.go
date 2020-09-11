package main

import (
	"patterns/pkg/flyweight"
)

const (
	golangDev = "golangDeveloper"
	pascalDev = "pascalDeveloper"
)

func main() {
	devFactory := flyweight.NewDeveloperFactory()
	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			devFactory.GetDeveloper(golangDev).WriteCode()
		} else {
			devFactory.GetDeveloper(pascalDev).WriteCode()
		}
	}
}
