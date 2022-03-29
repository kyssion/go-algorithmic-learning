package main

func main() {
	println(lengthOfLongestSubstring("abba"))
}

// todo - 运行时的边界问题考虑 - map 数据和下标需要映射和刷新
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	sb := []int32(s)
	maxLength := 0
	sbIndexMap := make(map[int32]int, len(s))
	start, end := 0, 0
	for ; end < len(sb); end++ {
		if index, ok := sbIndexMap[sb[end]]; ok && index >= start {
			start = index + 1
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		sbIndexMap[sb[end]] = end
	}
	return maxLength
}
