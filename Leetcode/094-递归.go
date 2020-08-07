package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := append(inorderTraversal(root.Left), root.Val)
	r = append(r, inorderTraversal(root.Right)...) //...代表右子树元素被打散一个个append到根节点后面
	return r
}
