package Leetcode

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return subResult(x, n)
	}
	return 1.0 / subResult(x, -n)
}

func subResult(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := subResult(x, n/2)
	if n%2 == 0 {
		return result * result
	} else {
		return result * result * x
	}
}
