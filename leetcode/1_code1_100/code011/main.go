package main

func main() {
	println(maxArea([]int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}))
}

// todo , 控制变量法方法, 容积= 长度* 高度 , 最大值寻找方法就是控制高度最高, 或者长度变低导致的数据变化
// 方法1 , 控制长度
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for {
		if left >= right {
			break
		}
		leftNum := height[left]
		rightNum := height[right]
		if leftNum > rightNum {
			max = maxInt(max, rightNum*(right-left))
			rightNow := right
			for {
				if rightNow <= left || height[rightNow] > height[right] {
					break
				}
				rightNow--
			}
			right = rightNow
		} else {
			max = maxInt(max, leftNum*(right-left))
			leftNow := left
			for {
				if leftNow >= right || height[leftNow] > height[left] {
					break
				}
				leftNow++
			}
			left = leftNow
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
