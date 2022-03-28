package main

import "fmt"

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
*/

func main() {
	fmt.Printf("%+v\n", twoSumV3([]int{
		2, 7, 5, 43, 2,
	}, 9))
}

// 方法1 : 构建map 查找
func twoSumV1(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for i, j := range nums {
		indexMap[j] = i
	}
	for i, j := range nums {
		if index, ok := indexMap[target-j]; ok && index != i {
			return []int{
				i, index,
			}
		}
	}
	return nil
}

// 方法2 简单暴力
func twoSumV2(nums []int, target int) []int {
	for i, j := range nums {
		for p := i + 1; p < len(nums); p++ {
			if j+nums[p] == target {
				return []int{
					i, p,
				}
			}
		}
	}
	return nil
}

// todo 方法3: 构建map 查找 - 优化, 其实在构建过程中就可以构建map了 - 不用完全构建
func twoSumV3(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for i, j := range nums {
		try := target - j
		if index, ok := indexMap[try]; ok {
			return []int{
				index, i,
			}
		}
		indexMap[j] = i
	}
	return nil
}
