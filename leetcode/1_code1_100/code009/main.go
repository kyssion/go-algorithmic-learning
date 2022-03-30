package main

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	p := reverse(x)
	return p == x
}

func reverse(x int) int {
	ans := 0
	for {
		// todo 这里要注意下 , 先check 最大最小值的边界问题
		if ans >= math.MaxInt32 || ans <= math.MinInt32 {
			return 0
		}
		if x == 0 {
			break
		}
		next := x % 10
		ans = ans*10 + next
		x = x / 10
	}
	return ans
}
