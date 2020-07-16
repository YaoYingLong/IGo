package test

import (
	"first/function"
	"testing"
)

func TestLinkedListInsertHead(t *testing.T) {
	var head = new(function.Node)
	var tail *function.Node
	tail = head
	for i := 1; i < 10; i++ {
		node := function.Node{Data: i}
		node.Next = tail
		tail = &node
	}
	function.ShowNode(tail)
}

func TestLinkedListInsertEnd(t *testing.T) {
	var head = new(function.Node)
	var tail *function.Node
	tail = head
	for i := 1; i < 10; i++ {
		node := function.Node{Data: i}
		tail.Next = &node
		tail = &node
	}
	function.ShowNode(head)
}
