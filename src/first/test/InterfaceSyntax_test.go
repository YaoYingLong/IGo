package test

import (
	"first/function"
	"fmt"
	"testing"
)

func TestInterBaseFunc(t *testing.T) {
	function.InterBaseFunc()
}

func TestType(t *testing.T) {
	{
		var x interface{}
		x = 10
		value, ok := x.(int)
		fmt.Println(value, ",", ok)
	}
	{
		var x interface{}
		x = 10
		value, ok := x.(string)
		fmt.Println(value, ",", ok)
	}
	{
		var x interface{}
		x = "Hello"
		value, ok := x.(string)
		fmt.Println(value, ",", ok)
	}
	{
		var x interface{}
		x = 10
		value, ok := x.(string)
		fmt.Println(value, ",", ok)
	}
	{
		var x interface{} = nil
		{
			value, ok := x.(string)
			fmt.Println(value, ",", ok)
		}
		{
			value, ok := x.(int)
			fmt.Println(value, ",", ok)
		}
	}
}

func TestTypeInterface(t *testing.T) {
	var tk interface{} = new(function.NetWriter)
	{
		rw2, ok2 := tk.(function.NetWriter)
		fmt.Println(ok2, ";", rw2.WriterClose(), rw2.WriteData("rw2"), rw2.CanWriter())
	}
	{
		rw2, ok2 := tk.(function.DataWriter)
		fmt.Println(ok2, ";", rw2.WriteData("rw2"), rw2.CanWriter())
	}
	{
		rw2, ok2 := tk.(function.Closer)
		fmt.Println(ok2, ";", rw2.WriterClose())
	}
	{
		rw2, ok2 := tk.(function.DataWriteCloser)
		fmt.Println(ok2, ";", rw2.WriterClose(), rw2.WriteData("rw2"), rw2.CanWriter())
	}
	{
		rw2, ok2 := tk.(function.FileWriter)
		fmt.Println(ok2, ";", rw2.WriterClose(), rw2.WriteData("rw2"))
	}
	{
		rw2, ok2 := tk.(function.FileWriter)
		fmt.Println(ok2, ";", rw2, rw2.WriterClose(), rw2.WriteData("rw2"))
	}
}

func TestNullInterface(t *testing.T) {
	{
		var a interface{} = 100
		var b interface{} = "hi"
		fmt.Println("a==b:", a == b)
	}
	{
		var a interface{} = []int{10}
		var b interface{} = []int{10}
		// 不能比较空接口中的动态值
		fmt.Println("a==b:", a == b)
	}
}

func TestAnyTypeDictionary(t *testing.T) {
	dict := function.NewDictionary()
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Dont Hungry", 24)
	dict.Set("Dont Hungry2", "68")

	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	dict.Visit(func(key, value interface{}) bool {
		if val, ok := value.(int); ok && val > 40 {
			fmt.Println(key, "is expensive")
			return true
		}

		fmt.Println(key, "is cheap")
		return true
	})
}

func TestTypeSwitch(t *testing.T) {
	var tk interface{} = new(function.NetWriter)
	switch tk.(type) {
	case function.NetWriter:
		fmt.Println("this type is function.NetWriter")
	case function.FileWriter:
		fmt.Println("this type is function.FileWriter")
	case function.DataWriteCloser:
		fmt.Println("this type is function.DataWriteCloser")
	case function.DataWriter:
		fmt.Println("this type is function.DataWriter")
	case function.Closer:
		fmt.Println("this type is function.Closer")
	default:
		fmt.Println("this type is nil")
	}
}

func TestCustError(t *testing.T) {
	result, error := function.Sqrt(-13)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

}
