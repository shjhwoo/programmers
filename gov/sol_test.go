package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func setTree() *Node {
	tree := &Node{
		Value: 1,
	}

	tree.Left = NewNode(2)
	tree.Right = NewNode(3)

	tree.Left.Left = NewNode(4)
	tree.Left.Right = NewNode(5)

	tree.Right.Left = NewNode(6)
	tree.Right.Right = NewNode(7)

	tree.Right.Right.Left = NewNode(8)
	tree.Right.Right.Right = NewNode(9)

	return tree
}

func (n *Node) PrintPreOrder() []int {
	mid := []int{n.Value}

	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintPreOrder()
	rights := n.Right.PrintPreOrder()

	return append(append(mid, lefts...), rights...)
}

func (n *Node) PrintInOrder() []int {
	mid := []int{n.Value}
	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintInOrder()
	rights := n.Right.PrintInOrder()

	return append(append(lefts, mid...), rights...)
}

func (n *Node) PrintPostOrder() []int {
	mid := []int{n.Value}
	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintPostOrder()
	rights := n.Right.PrintPostOrder()

	return append(append(lefts, rights...), mid...)
}

func TestPreOrderTree(t *testing.T) {
	tree := setTree()

	preOrder := tree.PrintPreOrder()

	assert.DeepEqual(t, []int{1, 2, 4, 5, 3, 6, 7, 8, 9}, preOrder)
}

func TestInOrder(t *testing.T) {
	tree := setTree()

	inOrder := tree.PrintInOrder()

	assert.DeepEqual(t, []int{4, 2, 5, 1, 6, 3, 8, 7, 9}, inOrder)

}

func TestPostOrder(t *testing.T) {
	tree := setTree()

	inOrder := tree.PrintPostOrder()

	assert.DeepEqual(t, []int{4, 5, 2, 6, 8, 9, 7, 3, 1}, inOrder)

}
