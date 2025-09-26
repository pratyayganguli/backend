package data

type DllNode struct {
	Data int
	Next *DllNode
	Prev *DllNode
}

var HeadDLL *DllNode
var TailDLL *DllNode

func insertDoublyLinkedList(data int) {
	// todo: implement the business logic
}

func TestInsertDoublyLinkedList() {
	insertDoublyLinkedList(10)
	insertDoublyLinkedList(20)
	insertDoublyLinkedList(30)
}

func doublyLinkedListTop() {
	// print the top element in the list
}
