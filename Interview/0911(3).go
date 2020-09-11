package main

import (
	"fmt"
	"regexp"
)

func CheckPassword(ps string) string {
	if len(ps) < 8 || len(ps) > 100 {
		return "Irregular password"
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_/\<>]{1}`
	a, _ := regexp.MatchString(num, ps)
	b, _ := regexp.MatchString(a_z, ps)
	c, _ := regexp.MatchString(A_Z, ps)
	d, _ := regexp.MatchString(symbol, ps)
	if !(a && b && c && d) {
		return "Irregular password"
	}

	return "Ok"
}

func main() {
	strArr := make([]string, 20)
	for i := 0; i < 20; i++ {
		fmt.Scanf("%s", &strArr[i])
		if strArr[i] == "" {
			break
		}
	}

	for j := 0; j < len(strArr); j++ {
		if strArr[j] != "" {
			fmt.Println(CheckPassword(strArr[j]))
		}
	}

}
