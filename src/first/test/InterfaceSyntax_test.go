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
		value := x.(string)
		fmt.Println(value)
	}
}
