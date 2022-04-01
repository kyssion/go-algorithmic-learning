package main

func main() {
	println(longestCommonPrefix([]string{
		"flower", "flow", "flight",
	}))
	println(longestCommonPrefix([]string{
		"flower", "flow", "floight", "flo",
	}))
	println(longestCommonPrefix([]string{
		"",
	}))
}

// todo 这里有一个小技巧就是, 变量可以提前设置:
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	index := 0
	str := ""
	for len(strs[0]) > index {
		baseC := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if len(strs[i])-1 < index || strs[i][index] != baseC {
				return str
			}
		}
		str = str + string(baseC)
		index++
	}
	return str
}
