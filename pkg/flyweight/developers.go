package flyweight

import "fmt"

type Developer interface {
	WriteCode()
}

type GolangDeveloper struct{}

func (g *GolangDeveloper) WriteCode() {
	fmt.Println("Writing in Golang")
}

type PascalDeveloper struct{}

func (p *PascalDeveloper) WriteCode() {
	fmt.Println("Writing in Pascal")
}
