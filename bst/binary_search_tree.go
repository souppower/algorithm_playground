package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.Value {
		return n.Left.Search(value)
	} else if n.Value < value {
		return n.Right.Search(value)
	}

	return true
}

func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		}
		n.Left.Insert(value)
	} else if n.Value < value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		}
		n.Right.Insert(value)
	}
}

func (n *Node) Delete(value int) *Node {
	if n.Value < value {
		n.Right = n.Right.Delete(value)
	} else if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		min := n.Right.Min()
		n.Value = min
		n.Right = n.Right.Delete(min)
	}

	return n
}

func (n *Node) Min() int {
	if n.Left == nil {
		return n.Value
	}
	return n.Left.Min()
}
