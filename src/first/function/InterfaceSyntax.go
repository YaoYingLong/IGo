package function

import "fmt"

type Inveker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

type FuncCallerV2 func(interface{})
func (f FuncCallerV2) Call(p interface{}) {
	fmt.Println("form function:", p)
}

func InterBaseFunc() {
	var invoker Inveker = new(Struct)
	invoker.Call("hello1")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("form function:", v)
	})
	invoker.Call("hello2")

	invoker = new(FuncCallerV2)
	invoker.Call("hello3")
}