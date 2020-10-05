package Leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 0) //存放结果
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	//将第一个元素添加到切片中，方便下面的操作
	for _, v := range m[string(digits[0])] {
		res = append(res, string(v))
	}

	//如果只有一个元素，则返回结果
	if len(digits) == 1 {
		return res
	}

	//循环，依次增加
	for i := 1; i < len(digits); i++ {
		res = combination(res, m[string(digits[i])])
	}

	return res
}

func combination(arr []string, s string) (res []string) {
	for _, v := range s {
		for i := 0; i < len(arr); i++ {
			res = append(res, arr[i]+string(v))
		}
	}
	return
}
