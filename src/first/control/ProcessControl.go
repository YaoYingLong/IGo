package control

import "fmt"

func ifSimple() {
	if index := 12; index > 10 {
		index--
		fmt.Println("index > 10", index)
		index++
		fmt.Println("index > 10", index)
	} else {
		fmt.Println("index < 10")
	}
}

func forSimple() {
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}

	var index int
	for index < 10 {
		index++
	}
	fmt.Println("index:", index)

	str := "12456789"
	for pos, char := range str {
		fmt.Println("pos, char:", pos, char)
	}
}

func forChannel() {
	channel := make(chan int)
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()
	for value := range channel {
		fmt.Println("channel value:", value)
	}
}

func switchSimple() {
	{
		str := "kk"
		switch str {
		case "hello", "kk":
			fmt.Println(1)
		case "world":
			fmt.Println(2)
		default:
			fmt.Println(3)
		}
	}
	{
		r := 11
		switch {
		case r > 10 && r < 20:
			fmt.Println(r)
		}
	}
	{
		str := "hello"
		switch {
		case str == "hello":
			fmt.Println("hello")
			fallthrough
		case str != "world":
			fmt.Println("world")

		}
	}
}

func gotoSimple() {
	if index := 10; index == 10 {
		goto onExit
	}
onExit:
	fmt.Println("exit")
}

func BaseFunc() {
	ifSimple()
	forSimple()
	forChannel()
	switchSimple()
	gotoSimple()
}
