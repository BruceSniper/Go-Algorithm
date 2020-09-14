package main

import "fmt"

func GetCommon(str1 string, str2 string) string {
	chs1 := len(str1)
	chs2 := len(str2)
	maxLength := 0
	end := 0

	rows := 0
	cols := chs2 - 1
	for rows < chs1 {
		i, j := rows, cols
		length := 0
		for i < chs1 && j < chs2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}
	return str1[(end - maxLength + 1):(end + 1)]
}

func main() {
	var str1, str2 string
	fmt.Scanf("%s,%s", &str1, &str2)
	fmt.Println(GetCommon(str1, str2))
}
