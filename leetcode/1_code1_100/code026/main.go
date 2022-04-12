package main

func main() {
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	for base, i := 0, 1; i < len(nums); i++ {
		if nums[base] == nums[i] {
			length--
		} else {
			base++
			nums[base] = nums[i]
		}
	}
	return length
}
