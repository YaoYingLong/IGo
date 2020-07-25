package test

import (
	"fmt"
	"testing"
)

func TestArrary(t *testing.T) {
	array := [5]int{1: 10, 3: 30}
	fmt.Println("array:", array)
}

func TestSlice(t *testing.T) {
	{
		var slice []int
		fmt.Println("slice == nil:", slice == nil, len(slice), cap(slice))
		slice = append(slice, 1)
		fmt.Println("slice:", slice)
		fmt.Println("slice == nil:", slice == nil, len(slice), cap(slice))
	}
	{
		fmt.Println()
		slice := []int{}
		fmt.Println("slice == nil:", slice == nil, len(slice), cap(slice))
		slice = append(slice, 1)
		fmt.Println("slice:", slice)
		fmt.Println("slice == nil:", slice == nil, len(slice), cap(slice))
	}
	{
		fmt.Println()
		var arr = [...]int{1, 2, 3, 4, 5}
		{
			slice := arr[:]
			fmt.Println("slice len5 cap5:", slice, len(slice), cap(slice))
		}
		{
			slice := arr[1:3]
			fmt.Println("slice len2 cap4:", slice, len(slice), cap(slice))
		}
		{
			slice := arr[1:]
			fmt.Println("slice len4 cap4:", slice, len(slice), cap(slice))
		}
		{
			slice := arr[:3]
			fmt.Println("slice len3 cap5:", slice, len(slice), cap(slice))
		}
	}
	{
		fmt.Println()
		var arr = [...]int{1, 2, 3, 4, 5}
		slice := arr[1:3]

		fmt.Println("slice[0]:", slice[0])
		fmt.Println("slice[1]:", slice[1])
		slice[0] = 12
		fmt.Println("arr:", arr)

		slice = append(slice, 13)
		slice[0] = 14
		fmt.Println("arr1:", arr)
	}
	{
		fmt.Println()
		var arr = [...]int{1, 2, 3, 4, 5}
		slice := arr[1:3:3]
		fmt.Println("slice len2 cap2:", slice, len(slice), cap(slice))

		fmt.Println("slice[0]:", slice[0])
		fmt.Println("slice[1]:", slice[1])
		slice[0] = 12
		fmt.Println("arr:", arr)

		slice = append(slice, 13)
		slice[0] = 14
		fmt.Println("slice:", slice)
		fmt.Println("arr1:", arr)
	}
	{
		fmt.Println()
		var arr = []int{1, 2, 3, 4, 5}
		slice := arr[1:3:3]
		fmt.Println("slice len2 cap2:", slice, len(slice), cap(slice))

		fmt.Println("slice[0]:", slice[0])
		fmt.Println("slice[1]:", slice[1])
		slice[0] = 12
		fmt.Println("arr:", arr)

		slice = append(slice, 13)
		slice[0] = 14
		fmt.Println("slice:", slice)
		fmt.Println("arr1:", arr)
	}
}
