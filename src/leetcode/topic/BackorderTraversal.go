package topic

/*
894
所有可能的满二叉树
*/
var cache = make(map[int][]*TreeNode)

func allPossibleFBT(N int) []*TreeNode {
	if _, isExist := cache[N]; isExist {
		return cache[N]
	}

	nResult := make([]*TreeNode, 0)
	if N%2 == 0 {
		return nResult
	}

	if N == 1 {
		node := &TreeNode{Val: 0}
		nResult = append(nResult, node)
	} else {
		for x := 1; x < N; x++ {
			y := N - x - 1
			leftResult := allPossibleFBT(x)
			rightResult := allPossibleFBT(y)
			for _, left := range leftResult {
				for _, right := range rightResult {
					node := &TreeNode{Val: 0}
					node.Left = left
					node.Right = right
					nResult = append(nResult, node)
				}
			}
		}
	}
	cache[N] = nResult
	return nResult
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

/*
814
二叉树剪枝
*/
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := pruneTree(root.Left)
	right := pruneTree(root.Right)
	if left == nil {
		root.Left = nil
	}
	if right == nil {
		root.Right = nil
	}
	if root.Val == 0 && left == nil && right == nil {
		return nil
	}
	return root
}
