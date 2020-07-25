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

type cusUser struct {
	name  string
	email string
	addr  []string
}

func (u cusUser) notifyV1() {
	fmt.Println("V1 Sending User Email To ", u.name, u.email, u.addr)
}

func (u *cusUser) notifyV2() {
	fmt.Println("V2 Sending User Email To ", u.name, u.email, u.addr)
}

func (u *cusUser) changeEmailV1(email string) {
	u.email = email
}

func (u cusUser) changeEmailV2(email string) {
	u.email = email
}

func (u *cusUser) changeAddrV11(addr []string, addAddr string) {
	u.addr = append(addr, addAddr)
}

func (u *cusUser) changeAddrV12(addAddr string) {
	u.addr = append(u.addr, addAddr)
}

func (u *cusUser) changeAddrV13(addr []string, addAddr string) {
	addr = append(addr, addAddr)
}

func (u cusUser) changeAddrV21(addr []string, addAddr string) {
	u.addr = append(addr, addAddr)
}

func (u cusUser) changeAddrV22(addAddr string) {
	u.addr = append(u.addr, addAddr)
}

func (u cusUser) changeAddrV23(addr []string, addAddr string) {
	addr = append(addr, addAddr)
}

func TestMethod(t *testing.T) {
	bill := cusUser{"bill", "bill@email.com", []string{"a", "b"}}
	bill.notifyV1()
	bill.notifyV2()

	bill.changeEmailV1("bill@emailv1.com")
	bill.notifyV1()
	bill.notifyV2()

	bill.changeEmailV2("bill@emailv2.com")
	bill.notifyV1()
	bill.notifyV2()

	bill.changeAddrV11(bill.addr, "c")
	bill.notifyV1()

	bill.changeAddrV12("d")
	bill.notifyV1()

	bill.changeAddrV13(bill.addr, "e")
	bill.notifyV1()

	bill.changeAddrV21(bill.addr, "f")
	bill.notifyV1()

	bill.changeAddrV22("g")
	bill.notifyV1()

	bill.changeAddrV23(bill.addr, "h")
	bill.notifyV1()

	lisa := &cusUser{"lisa", "lisa@emial.com", []string{"a", "b"}}
	lisa.notifyV1()
	lisa.notifyV2()

	lisa.changeEmailV1("lisa@emailv1.com")
	lisa.notifyV1()
	lisa.notifyV2()

	lisa.changeEmailV2("lisa@emailv2.com")
	lisa.notifyV1()
	lisa.notifyV2()
}
