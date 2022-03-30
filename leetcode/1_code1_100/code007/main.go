package main

import "math"

func main() {
	println(reverse(-3201000))
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

func reverseV2(x int) int {
	if x == 0 {
		return 0
	}
	var isF bool
	if x < 0 {
		isF = true
		x = -x
	}
	for {
		now := x % 10
		if now != 0 {
			break
		}
		x = x / 10
	}
	item, _ := getReverse(x)
	if isF {
		item = -item
	}
	if item >= math.MaxInt32 || item <= math.MinInt32 {
		return 0
	}
	return item
}

func getReverse(x int) (int, int) {
	if x == 0 {
		return x, 1
	}
	now := x % 10
	all, times := getReverse(x / 10)
	return now*times + all, times * 10
}
