package iterator

type Iterator interface {
	Next() *TreeNode
	HasNext() bool
	Range() []*TreeNode
}

type treeInorderIterator struct {
	index         int
	iterationList []*TreeNode
}

func (it *treeInorderIterator) Next() *TreeNode {
	if it.index == len(it.iterationList) {
		return nil
	}
	it.index++
	return it.iterationList[it.index-1]
}

func (it *treeInorderIterator) HasNext() bool {
	return it.index+1 == len(it.iterationList)
}

func (it *treeInorderIterator) Range() []*TreeNode {
	return it.iterationList
}

func (it *treeInorderIterator) createList(node *TreeNode) {
	if node == nil {
		return
	}
	it.createList(node.left)
	it.iterationList = append(it.iterationList, node)
	it.createList(node.right)
}

func NewIterator(tree Tree) Iterator {
	it := &treeInorderIterator{}
	root := tree.GetRoot()
	it.createList(root)
	return it
}
