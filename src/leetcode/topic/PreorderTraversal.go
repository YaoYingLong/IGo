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
