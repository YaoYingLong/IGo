package test

import (
	"first/function"
	"fmt"
	"reflect"
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

func TestStructTag(t *testing.T) {
	type Ins struct {
		in1 int `key1:"value1" key2:"value2"`
		in2 int `key1:"value1" key2:"value2"`
	}
	ins := Ins{in1: 1, in2: 2}
	fmt.Println(ins)

	typeOfIns := reflect.TypeOf(Ins{})
	if catType, ok := typeOfIns.FieldByName("in2"); ok {
		fmt.Println(catType.Tag.Get("key1"))
	}
}
