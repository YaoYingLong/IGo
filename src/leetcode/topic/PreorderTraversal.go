package topic

/*
144
二叉树的前序遍历
*/
func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	left := preorderTraversal(root.Left)
	result = append(result, left...)
	right := preorderTraversal(root.Right)
	result = append(result, right...)
	return result
}

/*
剑指 Offer 26
树的子结构
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return isSubTrue(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubTrue(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return isSubTrue(A.Left, B.Left) && isSubTrue(A.Right, B.Right)
}
