package main

func main() {
	print(isValid("()[]{}"))
}

func isValid(s string) bool {
	arr := []byte(s)
	statusArr := make([]byte, len(s))
	index := 0
	for i := 0; i < len(s); i++ {
		if arr[i] == '(' {
			statusArr[index] = ')'
			index++
			continue
		}
		if arr[i] == '[' {
			statusArr[index] = ']'
			index++
			continue
		}
		if arr[i] == '{' {
			statusArr[index] = '}'
			index++
			continue
		}
		if index == 0 || statusArr[index-1] != arr[i] {
			return false
		}
		index--
	}
	// todo 这里需要注意一下， 并不是循环走出来了就能返回true， 还要保证 左括号已经被处理干净了
	return index == 0
}
