package main

import "fmt"

func main() {
	var str1, str2 string
	for i := 0; i < 10; i++ {
		fmt.Scanf("%s %s", &str1, &str2)
		if str1 == "" && str2 == "" {
			break
		}
		fmt.Println(isMatch(str1, str2))
	}

}

func isMatch(str1, str2 string) bool {
	if str1 == str2 {
		return true
	}

	if len(str2) > 1 && str2[1] == '*' {
		switch str2[0] {
		case '.':
			for i := 0; i <= len(str1); i++ {
				if isMatch(str1, str2[2:]) {
					return true
				}
			}
		default:
			if isMatch(str1, str2[2:]) {
				return true
			}
			for i := 0; i < len(str1) && str1[1] == str2[0]; i++ {
				if isMatch(str1[i+1:], str2[2:]) {
					return true
				}
			}
		}
	} else {
		if str1 == "" || str2 == "" {
			return false
		}
		if str1[0] == str2[0] || str2[0] == '.' {
			return isMatch(str1[1:], str2[1:])
		}
	}
	return false
}
