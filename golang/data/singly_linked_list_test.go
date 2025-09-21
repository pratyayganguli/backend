package data

import (
	"fmt"
	"log"
	"testing"
)

type SllNode struct {
	Data int
	Next *SllNode
}

var HeadSLL *SllNode
var TailSLL *SllNode

func InsertSinglyLinkedList(data int) {
	if HeadSLL == nil {
		HeadSLL = &SllNode{
			Data: data,
		}
		TailSLL = HeadSLL
		HeadSLL.Next = TailSLL
		return
	} else {
		n := &SllNode{
			Data: data,
		}
		TailSLL.Next = n
		TailSLL = n
		return
	}
}

func TestInsertSinglyLinkedList(t *testing.T) {
	InsertSinglyLinkedList(10)
	InsertSinglyLinkedList(20)
	InsertSinglyLinkedList(30)
	InsertSinglyLinkedList(40)
	traverse()
	top()
}

func traverse() {
	if HeadSLL == nil {
		log.Fatalf("head of the singly linked list is null")
		return
	} else {
		n := HeadSLL
		for n != nil {
			fmt.Println(n.Data)
			n = n.Next
		}
	}
}

func top() {
	if HeadSLL != nil {
		fmt.Println(HeadSLL.Data)
	} else {
		fmt.Println("No data present in the linked list!")
	}
}
