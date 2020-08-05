package test

import (
	"leetcode/topic"
	"testing"
)

/*
[3,5,1,6,2,0,8,null,null,7,4]
5
2
*/
func TestDistanceK(t *testing.T) {
	node0 := new(topic.TreeNode)
	node0.Val = 0
	node4 := new(topic.TreeNode)
	node4.Val = 4
	node6 := new(topic.TreeNode)
	node6.Val = 6
	node7 := new(topic.TreeNode)
	node7.Val = 7
	node8 := new(topic.TreeNode)
	node8.Val = 8
	node2 := new(topic.TreeNode)
	node2.Val = 2
	node2.Left = node7
	node2.Right = node4
	node1 := new(topic.TreeNode)
	node1.Val = 1
	node1.Left = node0
	node1.Right = node8
	node5 := new(topic.TreeNode)
	node5.Val = 5
	node5.Left = node6
	node5.Right = node2
	node3 := new(topic.TreeNode)
	node3.Val = 3
	node3.Left = node5
	node3.Right = node1
	topic.DistanceK(node3, node5, 2)
}
