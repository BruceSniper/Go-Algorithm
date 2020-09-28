package main

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
	//// 先把root压入栈中
	//stack = append(stack, root)
	//if len(stack) > 0 || root != nil {
	//	//获取栈顶元素
	//	top := stack[len(stack)-1]
	//	//弹出栈的顶层节点
	//	stack = stack[:len(stack)-1]
	//	if top != nil {
	//		if top.Right != nil {
	//			stack = append(stack, top.Right)
	//		}
	//		if top.Left != nil {
	//			stack = append(stack, top.Left)
	//		}
	//		stack = append(stack, top, nil)
	//	} else if len(stack) > 0 {
	//		top = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		result = append(result, top.Val)
	//	}

	//}

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
