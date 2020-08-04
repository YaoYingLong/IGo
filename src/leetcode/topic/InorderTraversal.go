package topic

import "sort"

/*
1305
两棵二叉搜索树中的所有元素
*/
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	one := inorderBst(root1)
	two := inorderBst(root2)
	one = append(one, two...)
	sort.Ints(one)
	return one
}

func inorderBst(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	left := inorderBst(root.Left)
	result = append(result, left...)
	result = append(result, root.Val)
	right := inorderBst(root.Right)
	result = append(result, right...)
	return result
}

/*
94
二叉树的中序遍历
*/
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	left := inorderTraversal(root.Left)
	result = append(result, left...)
	result = append(result, root.Val)
	right := inorderTraversal(root.Right)
	result = append(result, right...)
	return result
}
