package syntax

import (
	"fmt"
)

/*
	定义变量：var name [type] = value
	GO支持的数据类型：
	bool
	string
	int、int8、int16、int32、int64
	uint、uint8、uint16、uint32、uint64、uintptr(只在底层编程时才需要)
	byte // uint8 的别名
	rune // int32 的别名 代表一个 Unicode 码
	float32、float64
	complex64、complex128
*/
func defineVar() {
	{
		var a, b *int
		fmt.Println("定义的指针变量默认值：", a, b)
	}
	{
		var a, b int
		fmt.Println("定义的int变量默认值：", a, b)
	}
	{
		var a, b bool
		fmt.Println("bool 变量默认值：", a, b)
	}
	{
		var a, b string
		fmt.Println("string  变量默认值：", a, b)
	}

	var num1 = 11
	/*
			使用简短格式定义变量
			- 定义变量，同时显示初始化
			- 不能提供数据类型
			- 只能用在函数内部
		    - 必须有新的变量被定义
	*/
	num1, num2 := 12, 13
	fmt.Println("使用:=简短格式定义变量：", num1, num2)

	var (
		a int
		b float32
		c string
		d [5]int
		//e func() bool
		//f struct {
		//	x int
		//}
	)
	a, b, c, d[0] = 1, 1.1, "str", 1
	fmt.Println(a, b, c, d)
}

/*
内容交换
*/
func swap() {
	var a = 100
	var b = 200
	fmt.Println("swap before:", a, b)
	a, b = b, a
	fmt.Println("swap before:", a, b)
}

/*
类型转换
*/
func trans() {

	/*
		// 字符串不能直接转int
		{
			var a = "15"
			fmt.Println("string to int: ", int(a))
		}*/
	{
		var a = 1.0
		fmt.Println("float to int: ", int(a))
	}
	{
		var a = 1
		fmt.Println("int to float32: ", float32(a))
	}
}

/*
指针
*/
func ptr() {
	var cat = 1
	var str = "banana"
	fmt.Println("%p %p", &cat, &str)
	var house = "Malibu Point 10880, 90265"
	ptr := &house
	fmt.Println("ptr type: ", ptr)
	fmt.Println("address: ", ptr)
	value := *ptr
	// 取值后的类型
	fmt.Println("value type: ", value)
}

/*
常量的定义：const name [type] = value
*/
func defineConst() {
	const (
		a = iota
		b
		c
	)
}

func BaseFunc() {
	defineVar()
	defineConst()
	swap()
	trans()
	ptr()
}
