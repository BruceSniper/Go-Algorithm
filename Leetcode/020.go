package Leetcode

func isValid(s string) bool {
	stack := make([]byte, 0)
	strMap := map[byte]byte{')': '(', ']': '[', '}': '{'}

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == strMap[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0

}

/*
func isValid(s string) bool {
    hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
    stack := make([]byte, 0)
    if s == "" {
        return true
    }

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}

作者：qi-ang-qi-ang
链接：https://leetcode-cn.com/problems/valid-parentheses/solution/jian-dan-ming-bai-shuang-bai-jie-fa-by-qi-ang-q-19/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
