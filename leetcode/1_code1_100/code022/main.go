package main

import "go-algorithmic-learning/lib"

func main() {
	lib.ShowStrArr([][]string{generateParenthesis(4)})
}

// 方法1 -> 暂时pass 有skip
// ()
// ((()) ()()
// ((())) (()()) ()(()) (()())
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{"()"}
	for n > 1 {
		n--
		nextAns := make([]string, 0)
		for _, str := range ans {
			baseStr := "(" + str
			nextAns = append(nextAns, baseStr+")")
			for i := len(str) - 1; i > 0; i-- {
				if baseStr[i] == ')' {
					continue
				}
				nextAns = append(nextAns, baseStr[:i]+")"+baseStr[i:])
			}
		}
		ans = nextAns
	}
	return ans
}
