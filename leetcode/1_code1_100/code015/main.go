package main

import (
	"go-algorithmic-learning/lib"
	"sort"
)

func main() {
	arr := threeSum([]int{-1, 0, 1, 2, -1, -4})
	lib.ShowIntArr(arr)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		end := len(nums) - 1
		for j := i + 1; j < len(nums) && end > j; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for end > j {
				now := nums[i] + nums[j] + nums[end]
				if now > 0 {
					end--
				} else if now < 0 {
					break
				} else {
					ans = append(ans, []int{
						nums[i], nums[j], nums[end],
					})
					break
				}
			}
		}
	}
	return ans
}
