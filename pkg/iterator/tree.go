package iterator

type Tree interface {
	Insert(value int)
}

type treeNode struct {
	Left  *treeNode
	Right *treeNode
	Value int
}

type tree struct {
	root *treeNode
}

func (t *tree) Insert(value int) {
	if t.root == nil {
		t.root = &treeNode{
			Value: value,
		}
		return
	}
	t.root.insert(value)
}

func (tn *treeNode) insert(value int) {
	if tn.Value < value {
		if tn.Left == nil {
			tn.Left = &treeNode{
				Value: value,
			}
		} else {
			tn.Left.insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &treeNode{
				Value: value,
			}
		} else {
			tn.Right.insert(value)
		}
	}
}

func NewTree() Tree {
	return &tree{
		root: nil,
	}
}
