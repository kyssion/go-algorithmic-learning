package main

func main() {

}

func removeElement(nums []int, val int) int {
	length := len(nums)
	for changeIndex, i := 0, 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[changeIndex] = nums[i]
			changeIndex++
			continue
		}
		length--
	}
	return length
}
