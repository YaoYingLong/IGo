package topic

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	maxDepth := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			treeNode := queue[0]
			queue = queue[1:]
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
			sz--
		}
		maxDepth++
	}
	return maxDepth
}

func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := recur(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := recur(root.Right)
	if rightDepth == -1 {
		return -1
	}
	if abs(leftDepth, rightDepth) < 2 {
		return max(leftDepth, rightDepth) + 1
	}
	return -1
}

func abs(left, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}

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

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	result = append(result, root.Val)
	childrens := root.Children
	for _, node := range childrens {
		result = append(result, preorder(node)...)
	}
	return result
}

func postorder(root *Node) []int {
	var result []int
	if root == nil {
		return []int{}
	}
	childrens := root.Children
	for _, node := range childrens {
		result = append(result, postorder(node)...)
	}
	result = append(result, root.Val)
	return result
}

func isUnivalTree(root *TreeNode) bool {
	isUnival := true
	if root == nil {
		return isUnival
	}
	if root.Left != nil {
		isUnival = isUnival && isUnivalTree(root.Left) && root.Val == root.Left.Val
	}
	if root.Right != nil {
		isUnival = isUnival && isUnivalTree(root.Right) && root.Val == root.Right.Val
	}
	return isUnival
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	isPathSum := hasPathSum(root.Left, sum)
	if isPathSum {
		return true
	}
	isPathSum = hasPathSum(root.Right, sum)
	if isPathSum {
		return true
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLen := len(queue)
		var levelResult []int
		for levelLen > 0 {
			treeNode := queue[0]
			queue = queue[1:]
			if treeNode == nil {
				levelResult = append(levelResult, -1)
			} else {
				queue = append(queue, treeNode.Left)
				queue = append(queue, treeNode.Right)
				levelResult = append(levelResult, treeNode.Val)
			}
			levelLen--
		}
		levelLen = len(levelResult)
		if levelLen == 1 {
			continue
		}
		if levelLen%2 != 0 {
			return false
		}
		for i := 0; i < levelLen/2; i++ {
			if levelResult[i] != levelResult[levelLen-i-1] {
				return false
			}
		}
	}
	return true
}
