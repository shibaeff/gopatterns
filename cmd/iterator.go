package main

import (
	"fmt"

	"patterns/pkg/iterator"
)

func main() {
	tree := iterator.NewTree()
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	iter := tree.GetIterator()
	for _, item := range iter.Range() {
		fmt.Println(item.Value)
	}
}
