package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := (len(nums1) + len(nums2)) / 2
	if len(nums1) < length {
		return findMedianSortedArrays(nums2, nums1)
	}
	start1, end1 := 0, length
	start2, end2 := 0, 0

}
