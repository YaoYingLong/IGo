package function

import (
	"flag"
	"fmt"
)

func funcA() (a, b int) {
	a = 1
	b = 2
	return
}

func funcB() (int, int) {
	return 3, 4
}

func funcC() (a, b string, c int) {
	return "a1", "b2", 5
}

func funcD(arr []int, f func(int)) {
	for _, value := range arr {
		f(value)
	}
}

func test() {
	{
		a, b := funcA()
		fmt.Println("a, b:", a, b)
	}
	{
		a, b := funcB()
		fmt.Println("a, b:", a, b)
	}
	{
		a, b, c := funcC()
		fmt.Println("a, b, c:", a, b, c)
	}
	// 函数变量
	{
		var f func() (int, int)
		f = funcA
		a, b := f()
		fmt.Println("a, b:", a, b)
	}
	{
		f := funcB
		a, b := f()
		fmt.Println("a, b:", a, b)
	}
	// 匿名函数
	{
		func(data int) {
			fmt.Println("inner func:", data)
		}(100)

		f := func(data int) {
			fmt.Println("inner func:", data)
		}
		f(500)
	}
	{
		funcD([]int{1, 2, 3, 4}, func(data int) {
			fmt.Println("this value:", data)
		})
	}
}

func funcE() {
	skillParam := flag.String("skill", "fire", "skill to perform")
	flag.Parse()

	skill := map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	f, ok := skill[*skillParam]
	fmt.Println("skill ok value:", ok)
	if ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func notFixedParam(args ...int) {
	for _, i := range args {
		fmt.Println("notFiexedParam:", i)
	}
}

func notFixedParamV2(format string, args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func rawPrint(rawList ...interface{}) {
	for _, raw := range rawList {
		fmt.Println(raw)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}


func BaseFunc() {
	test()
	funcE()
	accumulator := accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	accumulator2 := accumulate(10)
	fmt.Println(accumulator2())

	accumulator3 := accumulate(5)()
	fmt.Println(accumulator3)

	notFixedParam(1, 2)
	notFixedParamV2("kk", 1, 234, "hello", 3.14)

	print(1, 2, 3)
	//defer fmt.Println("宕机后要做的事情")
	//panic("宕机")
}
