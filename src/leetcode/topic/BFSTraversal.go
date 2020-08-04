package topic

import "fmt"

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLen := len(queue)
		var levelResult []int
		for levelLen > 0 {
			treeNode := queue[0]
			queue = queue[1:]
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
			levelResult = append(levelResult, treeNode.Val)
			levelLen--
		}
		result = append(result, levelResult)
	}
	return result
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLen := len(queue)
		var levelResult []int
		for levelLen > 0 {
			treeNode := queue[0]
			queue = queue[1:]
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
			levelResult = append(levelResult, treeNode.Val)
			levelLen--
		}
		result = append([][]int{levelResult}, result...)
	}
	return result
}

/*
剑指 Offer 32 - I
从上到下打印二叉树
*/
func levelOrderV2(root *TreeNode) []int {
	var resultList []int
	if root == nil {
		return resultList
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLen := len(queue)
		for levelLen > 0 {
			node := queue[0]
			resultList = append(resultList, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelLen--
		}
	}
	return resultList
}

/*
剑指 Offer 32 - III
从上到下打印二叉树 III
*/
func levelOrderV3(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	levelNum := 0
	for len(queue) > 0 {
		levelLen := len(queue)
		var subResult []int
		for levelLen > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if levelNum%2 == 0 {
				subResult = append(subResult, node.Val)
			} else {
				subResult = append([]int{node.Val}, subResult...)
			}
			levelLen--
		}
		result = append(result, subResult)
		levelNum++
	}
	return result
}

/*
623
在二叉树中增加一行
*/
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		node := &TreeNode{Val: v}
		node.Left = root
		return node
	}
	var queue []*TreeNode
	queue = append(queue, root)
	levelNum := 1
	for len(queue) > 0 {
		levelLen := len(queue)
		for levelLen > 0 {
			node := queue[0]
			queue = queue[1:]
			if levelNum == d-1 {
				nodeVLeft := &TreeNode{Val: v}
				nodeVLeft.Left = node.Left
				node.Left = nodeVLeft
				nodeVRight := &TreeNode{Val: v}
				nodeVRight.Right = node.Right
				node.Right = nodeVRight
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelLen--
		}
		if levelNum == d-1 {
			return root
		}
		levelNum++
	}
	return root
}

/*
863
二叉树中所有距离为 K 的结点
*/
var containSet map[*TreeNode]int

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	containSet = make(map[*TreeNode]int)
	parentMap := make(map[*TreeNode]*TreeNode)
	parent(root, nil, parentMap)

	containSet[target] = 1

	if containSet[target] != 1 {
		fmt.Println(1111)
	}

	var queue []*TreeNode
	queue = append(queue, target)
	levelNum := 0
	var result []int
	for len(queue) > 0 {
		levelLen := len(queue)
		for levelLen > 0 {
			node := queue[0]
			queue = queue[1:]
			if levelNum == K {
				result = append(result, node.Val)
			}
			if node.Left != nil && containSet[node.Left] != 1 {
				containSet[node.Left] = 1
				queue = append(queue, node.Left)
			}
			if node.Right != nil && containSet[node.Right] != 1 {
				containSet[node.Right] = 1
				queue = append(queue, node.Right)
			}
			parent := parentMap[node]
			if containSet[parent] != 1 {
				containSet[parent] = 1
				queue = append(queue, parent)
			}
			levelLen--
		}
		if levelNum == K {
			return result
		}
	}
	return result
}

func parent(node, parentNode *TreeNode, parentMap map[*TreeNode]*TreeNode) {
	if node != nil {
		parentMap[node] = parentNode
		containSet[node] = 0
		parent(node.Left, node, parentMap)
		parent(node.Right, node, parentMap)
	}
}

/*
102
二叉树的层序遍历
*/
func levelOrder102(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLen := len(queue)
		var levelResult = []int{}
		for levelLen > 0 {
			node := queue[0]
			queue = queue[1:]
			levelResult = append(levelResult, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelLen--
		}
		result = append(result, levelResult)
	}
	return result
}
