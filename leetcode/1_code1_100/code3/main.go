package main

/**
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成

*/

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
