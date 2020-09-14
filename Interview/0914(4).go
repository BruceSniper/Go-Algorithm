package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printNode(node *TreeNode) [][]int {
	// write code here
	var ret [][]int
	if node == nil {
		return ret
	}
	q := []*TreeNode{node}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			Node := q[j]
			ret[i] = append(ret[i], Node.Val)
			if Node.Left != nil {
				p = append(p, Node.Left)
			}
			if Node.Right != nil {
				p = append(p, Node.Right)
			}
		}
		q = p
	}
	return ret
}
