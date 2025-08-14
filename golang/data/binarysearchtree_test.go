package data

import (
	"fmt"
	"testing"
)

// todo: make it non-generic and then implement the generic part.
type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

var (
	RootNode    *BSTNode
	CurrentNode *BSTNode
	// using constants for better readability
	LeftDirection  string = "left"
	RightDirection string = "right"
)

func (n *BSTNode) Insert(data int, direction string) *BSTNode {
	n.Data = data
	if direction == LeftDirection {
		n.Left = &BSTNode{}
		CurrentNode = n.Left
		return n.Left
	} else {
		n.Right = &BSTNode{}
		CurrentNode = n.Right
		return n.Right
	}
}

func (n *BSTNode) IsEmptyBSTTree() bool {
	if RootNode == nil {
		fmt.Println("root node is empty")
		return true
	}
	return false
}

func add() int {
	return 1 + 2
}

func TestAdd(t *testing.T) {
	sum := add()
	fmt.Println(sum)
}
