package main

func maxArea(height []int) int {
	area := 0
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			area = max(area, min(height[i], height[j])*(j-i))
		}
	}
	return area
}

func maxArea1(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

func maxArea2(height []int) int {
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		// 朝着area具有变大的可能性方向变化。
		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func maxArea3(height []int) int {
	area := 0
	start, end := 0, len(height)-1
	for start < end {
		area = max(area, min(height[start], height[end])*(end-start))
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return area
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func main() {
	maxArea([]int{1, 2, 3, 1})
}
