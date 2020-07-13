package container

import (
	"container/list"
	"fmt"
	"sync"
)

func forArray() {
	var numArr = [3]int{2, 3, 4,}
	for i, v := range numArr {
		fmt.Println("i, v", i, v)
	}
}

func slice() {
	var arr = [...]int{1, 2, 3, 4, 5,}
	{
		slice := arr[0:]
		fmt.Println("slice[0:], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[:]
		fmt.Println("slice[:], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[1:3]
		fmt.Println("slice[1:3], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[1:]
		fmt.Println("slice[1:], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[:3]
		fmt.Println("slice[:3], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[:len(arr)]
		fmt.Println("slice[:len(arr)], len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := arr[:len(arr)-1]
		fmt.Println("slice[:len(arr)-1], len, cap:", slice, len(slice), cap(slice))
	}
	{
		var slice []int
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
		slice = append(slice, 1)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
	}
	{
		var slice = make([]int, 2, 5)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
		slice = append(slice, 1, 2, 3, 4)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
		slice = append(slice, []int{5, 6, 7}...)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
		slice = append([]int{-3, -2, -1}, slice...)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
	}
	{
		slice := []int{1, 2, 3, 4}
		slice = append(slice[:2], append([]int{6, 7}, slice[2:]...)...)
		fmt.Println("slice, len, cap:", slice, len(slice), cap(slice))
	}
	{
		sliceA := []int{1, 2, 3, 4, 5}
		sliceB := []int{6, 7, 8}
		copyCount := copy(sliceA, sliceB)
		fmt.Println("slice, len, cap, copyCount:", sliceA, len(sliceA), cap(sliceA), copyCount)
	}
	{
		sliceB := []int{1, 2, 3, 4, 5}
		sliceA := []int{6, 7, 8}
		copyCount := copy(sliceA, sliceB)
		fmt.Println("slice, len, cap, copyCount:", sliceA, len(sliceA), cap(sliceA), copyCount)
	}
}

func mapBase() {
	{
		var mapLit = map[int]string{}
		fmt.Println("mapLit1 len:", mapLit, len(mapLit))
	}
	{
		var mapLit map[int]string
		fmt.Println("mapLit2 len:", mapLit, len(mapLit))
	}
	{
		var mapLit = make(map[int]string, 16)
		fmt.Println("mapLit3 len:", mapLit, len(mapLit))
	}
	{
		mapLit := map[int]string{1: "a", 2: "b"}
		fmt.Println("mapLit len:", mapLit, len(mapLit))
		fmt.Println("mapLit 1:", mapLit[1])
		fmt.Println("mapLit 2:", mapLit[2])
		fmt.Println("mapLit 2:", mapLit[5])
		mapLit[3] = "c"
		fmt.Println("mapLit len:", mapLit, len(mapLit))
		for key, value := range mapLit {
			fmt.Println("key value:", key, value)
		}
		for key := range mapLit {
			fmt.Println("key:", key)
		}
		for _, value := range mapLit {
			fmt.Println("value:", value)
		}
		delete(mapLit, 1)
		fmt.Println("mapLit len:", mapLit, len(mapLit))
	}
	{
		mapLit := map[string]int{
			"hello": 100,
			"world": 200,
		}
		for key, value := range mapLit {
			fmt.Println("key value:", key, value)
		}
		{
			value, isExist := mapLit["hello"]
			fmt.Println("value, isExist:", value, isExist)
		}
		{
			value, isExist := mapLit["kk"]
			fmt.Println("value, isExist:", value, isExist)
		}
	}
}

func syncMap() {
	var syncMap sync.Map
	syncMap.Store(1, "a")
	syncMap.Store(2, "b")
	syncMap.Store(1, "C")
	fmt.Println("syncMap", syncMap)
	//syncMap.Load(func(key interface{}) {
	//	fmt.Println("syncMap 1", key)
	//})

	value, ok := syncMap.Load(2)
	fmt.Println("syncMap value, ok:", value, ok)

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("key, value", key, value)
		return true
	})
}

func listSimple() {
	lit := list.New()
	lit.PushBack("AA")
	lit.PushFront("BB")
	element := lit.PushFront("CC")
	lit.InsertBefore("DD", element)
	for i := lit.Front(); i != nil; i = i.Next() {
		fmt.Println("lit value:", i.Value)
	}
	lit.Remove(element)
	for i := lit.Front(); i != nil; i = i.Next() {
		fmt.Println("lit value:", i.Value)
	}
}

func BaseFunc() {
	forArray()
	slice()
	mapBase()
	syncMap()
	listSimple()
}
