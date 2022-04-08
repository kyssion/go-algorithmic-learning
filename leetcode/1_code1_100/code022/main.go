package main

import "go-algorithmic-learning/lib"

func main() {
	lib.ShowStrArr([][]string{generateParenthesis(4)})
}

// 方法1 -> 暂时pass 有skip
// ()
// ((()) ()()
// ((())) (()()) ()(()) (()())
func generateParenthesisV1(n int) []string {
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

func generateParenthesis(n int) []string {
	return do(n, n, "", []string{})
}

func do(i int, j int, s string, ans []string) []string {
	if i == 0 && j == 0 {
		ans = append(ans, s)
		return ans
	}
	if i > 0 {
		ans = do(i-1, j, s+"(", ans)
	}
	if j > 0 && j > i {
		ans = do(i, j-1, s+")", ans)
	}
	return ans
}
