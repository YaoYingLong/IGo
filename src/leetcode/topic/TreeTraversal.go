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

/*
1145
二叉树着色游戏
*/
var leftNum int
var rightNum int

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	getLeftAndRightCount(root, x)
	half := n / 2
	return leftNum > half || rightNum > half || (leftNum+rightNum) < half
}

func getLeftAndRightCount(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	left := getLeftAndRightCount(root.Left, x)
	right := getLeftAndRightCount(root.Right, x)
	if root.Val == x {
		leftNum = left
		rightNum = right
	}
	return left + right + 1
}

/*
105
从前序与中序遍历序列构造二叉树
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	indexMap := make(map[int]int, len(inorder))
	for index, val := range inorder {
		indexMap[val] = index
	}
	return buildTreeTrav(preorder, 0, len(preorder)-1, 0, indexMap)
}

func buildTreeTrav(preorder []int, preLeft, preRight, inLeft int, indexMap map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}

	root := &TreeNode{Val: preorder[preLeft]}
	if preLeft == preRight {
		return root
	}
	rootIndex := indexMap[root.Val]
	leftNodes := rootIndex - inLeft
	root.Left = buildTreeTrav(preorder, preLeft+1, preLeft+leftNodes, inLeft, indexMap)
	root.Right = buildTreeTrav(preorder, preLeft+leftNodes+1, preRight, rootIndex+1, indexMap)
	return root
}

func buildTreeV2(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var queue []*TreeNode
	queue = append(queue, root)
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		preVal := preorder[i]
		node := queue[len(queue)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preVal}
			queue = append(queue, node.Left)
		} else {
			for len(queue) > 0 && queue[len(queue)-1].Val == inorder[inorderIndex] {
				node = queue[len(queue)-1]
				queue = queue[0 : len(queue)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{Val: preVal}
			queue = append(queue, node.Right)
		}
	}
	return root
}

/*
114
二叉树展开为链表
*/
var cou *TreeNode

func flatten(root *TreeNode) {
	ans := &TreeNode{Val: 0}
	cou = ans
	flattenDfs(root)
	root = ans.Right
}

func flattenDfs(root *TreeNode) {
	if root == nil {
		return
	}
	leftNode := root.Left
	rightNode := root.Right
	root.Left = nil
	cou.Right = root
	cou = root
	flattenDfs(leftNode)
	flattenDfs(rightNode)
}
