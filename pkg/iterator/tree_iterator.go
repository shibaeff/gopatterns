package iterator

type Iterator interface {
	Next()
	HasNext() bool
}

type treeInorderIterator struct {
	iterationList []*treeNode
}

func (it *treeInorderIterator) createList(node *treeNode) {

}
