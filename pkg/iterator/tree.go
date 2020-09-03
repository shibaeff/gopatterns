package iterator

type Tree interface {
	Insert(value int)
	GetRoot() *TreeNode
	GetIterator() Iterator
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	Value int
}

type tree struct {
	root *TreeNode
}

func (t *tree) GetIterator() Iterator {
	return NewIterator(t)
}

func (t *tree) GetRoot() *TreeNode {
	return t.root
}

func (t *tree) Insert(value int) {
	if t.root == nil {
		t.root = &TreeNode{
			Value: value,
		}
		return
	}
	t.root.insert(value)
}

func (tn *TreeNode) insert(value int) {
	if value < tn.Value {
		if tn.left == nil {
			tn.left = &TreeNode{
				Value: value,
			}
		} else {
			tn.left.insert(value)
		}
	} else {
		if tn.right == nil {
			tn.right = &TreeNode{
				Value: value,
			}
		} else {
			tn.right.insert(value)
		}
	}
}

func NewTree() Tree {
	return &tree{
		root: nil,
	}
}
