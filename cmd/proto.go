package main

import "patterns/pkg/prototype"

func main() {
	dev1 := prototype.NewHuman(20, "Alex")
	dev1.Print()
	factory := prototype.NewHumanFactory()
	factory.SetPrototype(dev1)

	dev1Copy := factory.GetCopy()
	dev1Copy.(*prototype.Human).Print()
}
