package main

func main() {

}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	itemMap := make(map[int]int)
	for index, item := range nums {
		itemMap[item] = index
	}
	for i := 0; i < len(nums); i++ {
	}
	return ans
}
