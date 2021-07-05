package binarySearch

import (
	"errors"
	"fmt"
	"log"
)

var count int

type Tree struct {
	Root *Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (t *Tree) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{Key: val}
	} else {
		t.Root.Insert(val)
	}
}

func (t *Tree) Delete(s int) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	fakeParent := &Node{Right: t.Root}
	t.Root.Delete(s, fakeParent)

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax(n)
}

func (n *Node) ReplaceNode(parent, replacement *Node) {
	if n == parent.Left {
		parent.Left = replacement
		return
	}

	parent.Right = replacement
}

func (n *Node) Delete(val int, parent *Node) {
	switch {
	case val < n.Key:
		n.Left.Delete(val, n)
	case val > n.Key:
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

		replacement, replParent := n.Left.findMax(n)

		n.Key = replacement.Key

		replacement.Delete(replacement.Key, replParent)
	}
}

func (n *Node) Insert(val int) {
	if n.Key < val {
		if n.Right == nil {
			n.Right = &Node{Key: val}
			return
		} else {
			n.Right.Insert(val)
		}
	} else if n.Key > val {
		if n.Left == nil {
			n.Left = &Node{Key: val}
			return
		} else {
			n.Left.Insert(val)
		}
	}
}

func InOrderShow(n *Node) {
	if n == nil {
		return
	} else {
		InOrderShow(n.Left)
		fmt.Printf("%d ", n.Key)
		InOrderShow(n.Right)
	}
}

func PreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Key)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func (n *Node) Find(val int) {
	count++
	if n.Key < val {
		if n.Right == nil {
			fmt.Printf("not found, finding in %d step", count)
			return
		} else {
			n.Right.Find(val)
		}
	} else if n.Key > val {
		if n.Left == nil {
			fmt.Printf("not found, finding in %d step", count)
			return
		} else {
			n.Left.Find(val)
		}
	} else {
		fmt.Printf("found value in %d step", count)
		return
	}
}

func (n *Node) Change(new int, k int) string {

	if n == nil {
		return fmt.Sprintf("not found nodes %v", k)
	}

	if n.Key < k {
		return n.Right.Change(new, k)
	} else if n.Key > k {
		return n.Right.Change(new, k)
	}

	n.Key = new

	return fmt.Sprintf("success change %v to %v", k, new)
}

func BinaryTree() {
	var tree Tree

	data := []int{24, 20, 21, 23, 10, 5, 11, 30, 26, 35, 40}
	for i := 0; i < len(data); i++ {
		tree.Insert(data[i])
	}

	PreOrder(tree.Root)
	fmt.Println()

	err := tree.Delete(24)
	if err != nil {
		log.Fatal(err)
	}

	// tree.Delete(24)

	PreOrder(tree.Root)
	fmt.Println()

}
