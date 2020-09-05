package main

import strategy "patterns/pkg/state"

func main() {
	developer := strategy.NewDeveloper("Pasha")

	for i := 0; i < 10; i++ {
		developer.ExecuteActivity()
		developer.ChangeState()
	}
}
