package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f", findMedianSortedArrays([]int{}, []int{2, 3}))

	//fmt.Printf("%f", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	//fmt.Printf("%f", findMedianSortedArrays([]int{1, 2}, []int{3}))
	//fmt.Printf("%f", findMedianSortedArrays([]int{}, []int{3}))
	//fmt.Printf("%f", findMedianSortedArrays([]int{}, []int{}))
}

// 1. 方法1 - 简单双指针遍历
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allLength := len(nums1) + len(nums2)
	if allLength == 0 {
		return 0
	}
	nums1 = append(nums1, math.MaxInt)
	nums2 = append(nums2, math.MaxInt)
	needFind := allLength/2 + 1
	itemIndex1, itemIndex2, index := -1, -1, 0
	for index < needFind {
		if nums1[itemIndex1+1] >= nums2[itemIndex2+1] {
			itemIndex2++
		} else {
			itemIndex1++
		}
		index++
	}
	return getAnsNum(allLength, nums1, nums2, itemIndex1, itemIndex2)
}

func getAnsNum(length int, nums1 []int, nums2 []int, index1 int, index2 int) float64 {
	return 0
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
