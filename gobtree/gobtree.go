package gobtree

type Node struct {
	value int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

func Create() *Tree {
    return &Tree{nil}
}

func (n *Node) Insert(value int) {
	if n == nil {
		return
	} else if value < n.value {
		if n.left == nil {
			n.left = &Node{
				value: value,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{
				value: value,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.Insert(value)
		}
	}
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{
			value: value,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.Insert(value)
	}
}

func (n *Node) Exists(value int) bool {
	if n == nil {
		return false
	} else if n.value == value {
		return true
	} else if value < n.value {
		return n.left.Exists(value)
	} else {
		return n.right.Exists(value)
	}
}

func (t *Tree) Exists(value int) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.Exists(value)
	}
}
