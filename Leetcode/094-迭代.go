package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversalX(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}
		index := len(stack) - 1                   //栈顶
		result = append(result, stack[index].Val) //中序输出
		root = stack[index].Right                 //右节点会进入下次循环，如果 ==nil，继续出栈
		stack = stack[:index]                     //出栈
	}
	return result
}
