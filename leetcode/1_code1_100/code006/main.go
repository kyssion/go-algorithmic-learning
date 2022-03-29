package main

import "bytes"

func main() {
	println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans := make([]int32, len(s))
	arr := []int32(s)
	nextStep := numRows*2 - 2
	indexA := 0
	for i := 0; i < numRows; i++ {
		toIndex := i
		times := 1
		for {
			if toIndex >= len(s) {
				break
			}
			ans[indexA] = arr[toIndex]
			indexA++
			if i == 0 || times == 1 {
				toIndex = toIndex + nextStep - (i * 2)
			}
			if i == numRows-1 || times == 0 {
				toIndex = toIndex + i*2
			}
			times = times ^ 1
		}
	}
	return string(ans)
}

// convert todo 另一种思路吧 - 就是模拟添加过程(二维数组从上到下遍历) 然后使用bytes.join 方法
func convertV2(s string, numRow int) string {
	r := numRow
	if r == 1 || r >= len(s) {
		return s
	}
	t, idx := r*2-2, 0
	mat := make([][]byte, r)
	for index, val := range s {
		mat[idx] = append(mat[idx], byte(val))
		if index%t < r-1 {
			idx++
		} else {
			idx--
		}
	}
	return string(bytes.Join(mat, nil))
}
