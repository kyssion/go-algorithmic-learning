package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Printf("%f", findMedianSortedArrays([]int{1, 2}, []int{3}))
	fmt.Printf("%f", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

}

// 1. 方法1 - 简单双指针遍历
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
	}
	start1, start2, step := -1, -1, 0
	needSkip := int(math.Ceil(float64(len(nums1)+len(nums2)) / 2))
	for start1+1 < len(nums1) && start2+1 < len(nums2) && step < needSkip {
		if nums1[start1+1] > nums2[start2+1] {
			start2++
		} else {
			start1++
		}
		step++
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		if start1 == -1 {
			if start2 < len(nums2)-1 {
				return float64(nums2[start2]+nums2[start2+1]) / 2.0
			}
			return float64(nums2[start2]+nums1[0]) / 2.0
		} else if start2 == -1 {
			if start1 < len(nums1)-1 {
				return float64(nums1[start1]+nums1[start1+1]) / 2.0
			}
			return float64(nums1[start1]+nums2[0]) / 2.0
		}
		return float64(nums2[start2]+nums1[start1]) / 2.0
	}
	if start1 == -1 {
		return float64(nums2[start2])
	} else if start2 == -1 {
		return float64(nums1[start1])
	}
	return math.Max(float64(nums1[start1]), float64(nums2[start2]))
}

// 2. 方法2 -> 从num1 中获取一个数字 在num2 中查看位置 ， 将 1的位置+ 2的位置和目标位置比较获取最终结论

func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	start1, start2, step := 0, 0, 0
	length := (len(nums1) + len(nums2)) / 2
	for start1 < len(nums1) && start2 < len(nums2) && step < length {
		if nums1[start1] > nums2[start2] {
			start2++
		} else {
			start1++
		}
		step++
	}
	if length%2 == 0 {
		return float64(nums2[start2]+nums1[start1]) / 2.0
	}
	return math.Max(float64(nums1[start1]), float64(nums2[start2]))
}

// 3 。 中位数前缀支持

func findMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	start1, start2, step := 0, 0, 0
	length := (len(nums1) + len(nums2)) / 2
	for start1 < len(nums1) && start2 < len(nums2) && step < length {
		if nums1[start1] > nums2[start2] {
			start2++
		} else {
			start1++
		}
		step++
	}
	if length%2 == 0 {
		return float64(nums2[start2]+nums1[start1]) / 2.0
	}
	return math.Max(float64(nums1[start1]), float64(nums2[start2]))
}
