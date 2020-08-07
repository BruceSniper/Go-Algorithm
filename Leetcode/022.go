package Leetcode

import "fmt"

//左括号随时加，只要不超标
//右括号必须之前有左括号，且左个数>右个数

func generateParenthesis(n int) []string {
	list := new([]string)
	generate(0, 0, n, "", list)

	return *list

}

func generate(left, right, n int, s string, list *[]string) {
	//终止条件
	if left == n && right == n {
		*list = append(*list, s)
		return
	}

	if left < n {
		generate(left+1, right, n, s+"(", list)
	}

	if left > right && right < n {
		generate(left, right+1, n, s+")", list)
	}
}

func main() {
	fmt.Println(generateParenthesis(4))
}
