package test

import (
	"fmt"
	"leetcode/topic/graph"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	data := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(graph.AllPathsSourceTarget(data))
}
