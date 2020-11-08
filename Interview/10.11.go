package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 找到输入字符串中连续不含重复字符的最长子串。如果有多个相同长度的，只需取第一个。
 * @param str string字符串 非空字符串
 * @return string字符串
 */
func findMaxSubstr(str string) string {
	// write code here
	var result string
	var maxLength int
	switch {
	case len(str) == 0:
		return ""

	case len(str) == 1:
		return str

	default:
		low, fast := 0, 1
		for i := 0; i < len(str)-1; i++ {
			for j := low; j < fast; j++ {
				if str[j] == str[fast] {
					maxLength = len(str[low : fast+1])
					low = j + 1
				}
				if len(str[low:fast+1]) > maxLength {
					result = str[low:fast]
				}
			}
			fast++
		}

	}
	return result
}
