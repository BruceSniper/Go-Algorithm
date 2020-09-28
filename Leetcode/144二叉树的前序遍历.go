package Leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// 定义stack模拟二叉树，利用先入后出的特点，出栈顺序为中左右，进栈顺序则为右左中
	var stack []*TreeNode
	var result []int

	for 0 < len(stack) || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		top := len(stack) - 1 //栈顶
		root = stack[top]     //出栈
		stack = stack[:top]
	}

	return result
}
