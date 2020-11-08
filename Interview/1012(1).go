package main

import "fmt"

func main() {
	var n int
	var str string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &str)
	fmt.Println(Jump(str, n))
}

func Jump(str string, n int) int {
	for i := 0; i < n; i++ {
		if str[i] < 0 || str[i] > 9 {
			return -1
		}
	}
	m := map[int][]int{}
	for i := 0; i < n; i++ {
		m[int(str[i])] = append(m[int(str[i])], i)
	}

	re := 0
	queue := []int{0}
	vis := map[int]bool{}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			a := queue[0]
			if a == n-1 {
				return re
			}
			for _, v := range m[int(str[a])] {
				if !vis[v] {
					queue = append(queue, v)
					vis[v] = true
				}
			}
			delete(m, int(str[a]))
			if !vis[a+1] {
				queue = append(queue, a+1)
				vis[a+1] = true
			}
			if !vis[a-1] && a > 0 {
				queue = append(queue, a-1)
				vis[a-1] = true
			}
			queue = queue[1:]
		}
		re++
	}
	return 0

}
