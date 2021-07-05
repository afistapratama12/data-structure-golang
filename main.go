package main

import "fmt"

var count int

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (t *Tree) Delete(n int) {
	if t.Root == nil {
		fmt.Println("error cannot delete from an empty tree")
	}

	fakeParent := &Node{Right: t.Root}

	t.Root.Delete(n, fakeParent)

	if fakeParent.Right == nil {
		t.Root = nil
	}
}

func (n *Node) Delete(val int, parent *Node) {
	switch {
	case val < n.Value:
		n.Left.Delete(val, n)
	case val > n.Value:
		n.Right.Delete(val, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.ReplaceNode(parent, nil)
			return
		}

		if n.Left == nil {
			n.ReplaceNode(parent, n.Right)
			return
		}

		if n.Right == nil {
			n.ReplaceNode(parent, n.Left)
			return
		}

		replacement, replParent := n.Left.FindMax(n)

		n.Value = replacement.Value

		replacement.Delete(replacement.Value, replParent)
	}
}

func (n *Node) ReplaceNode(parent, replacement *Node) {
	if n == parent.Left {
		parent.Left = replacement
		return
	}

	parent.Right = replacement
}

func (n *Node) FindMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.FindMax(n)
}

func (t *Tree) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{Value: val}
		return
	} else {
		t.Root.Insert(val)
	}
}

func (n *Node) Insert(val int) {
	if n.Value > val {
		if n.Left == nil {
			n.Left = &Node{Value: val}
			return
		} else {
			n.Left.Insert(val)
		}
	} else if n.Value < val {
		if n.Right == nil {
			n.Right = &Node{Value: val}
			return
		} else {
			n.Right.Insert(val)
		}
	}
}

func PreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Value)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}

}

func InOrder(n *Node) {
	if n == nil {
		return
	} else {
		InOrder(n.Left)
		fmt.Printf("%d ", n.Value)
		InOrder(n.Right)
	}
}

func main() {
	var tree Tree

	nums := []int{25, 10, 5, 6, 30, 27, 35, 31, 40, 45}

	for i := 0; i < len(nums); i++ {
		tree.Insert(nums[i])
	}

	// PreOrder(tree.Root)
	// fmt.Println()
	InOrder(tree.Root)
	fmt.Println()

	tree.Delete(10)

	InOrder(tree.Root)
	fmt.Println()
}
