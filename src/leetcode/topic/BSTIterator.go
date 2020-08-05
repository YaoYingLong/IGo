package topic

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	queue []int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{queue: inorderDfs(root)}
}

func inorderDfs(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	node := root
	for node != nil || len(queue) > 0 {
		for node != nil {
			queue = append([]*TreeNode{node}, queue...)
			node = node.Left
		}
		node = queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	result := this.queue[0]
	this.queue = this.queue[1:]
	return result
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.queue) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
