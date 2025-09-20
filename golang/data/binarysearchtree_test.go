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
)

const (
	LeftDirection  string = "left"
	RightDirection string = "right"
)

// todo: change the signature of the function...
func InsertionBST() *BSTNode {
	node := &BSTNode{}
	node = node.Insert(10, LeftDirection)
	node = node.Insert(20, RightDirection)
	node = node.Insert(30, LeftDirection)
	node.Insert(40, RightDirection)
	return node
}

func (n *BSTNode) Insert(data int, direction string) *BSTNode {
	n.Data = data

	if RootNode == nil {
		RootNode = n
		CurrentNode = RootNode
		return RootNode
	}

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

func TestInsertionBST(t *testing.T) {
	InsertionBST()
}
