package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeTree struct {
	Root *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val, Left: new(Node), Right: new(Node)}
}

func (n *NodeTree) Add(val int) {
	if n.Root == nil {
		n.Root = NewNode(val)
	} else {
		n.Root.Add(val)
	}
}

func (n *Node) Add(val int) {
	if val <= n.Value {
		if n.Left != nil {
			n.Left.Add(val)
		} else {
			n.Left = NewNode(val)
		}
	} else {
		if n.Right != nil {
			n.Right.Add(val)
		} else {
			n.Right = NewNode(val)
		}
	}
}

func (n *NodeTree) Contains(target int) bool {
	node := n.Root
	for node != nil {
		if target < node.Value {
			node = node.Left
		} else if target > node.Value {
			node = node.Right
		} else {
			return true
		}
	}
	return false
}

func main() {
	nt := new(NodeTree)
	nt.Add(1)
	nt.Add(2)
	nt.Add(3)
	nt.Add(100)

	if !nt.Contains(4) {
		fmt.Println("Not Found.")
	}

	if nt.Contains(100) {
		fmt.Println("Found.")
	}
}
