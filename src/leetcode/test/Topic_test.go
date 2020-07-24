package test

import (
	"fmt"
	"leetcode/topic"
	"testing"
)

func TestWordChain(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(topic.FindLadders(beginWord, endWord, wordList))

	queue := [][]int{{1, 3, 5}, {2, 4, 7}, {6, 8, 9}, {16, 18, 19}}
	fmt.Println("queue:", queue)
	fmt.Println("queue[1]:", queue[1])
	fmt.Println("len queue:", len(queue))
	fmt.Println("cap queue:", cap(queue))
	for _, val := range queue {
		fmt.Println("val:", val)
	}

	sliceB := []int{1, 2, 3, 4, 5}
	sliceA := []int{6, 7, 8}
	copyCount := copy(sliceA, sliceB)
	fmt.Println("copyCount:", copyCount)
	fmt.Println("sliceB:", sliceB)
	fmt.Println("sliceA:", sliceA)
}
