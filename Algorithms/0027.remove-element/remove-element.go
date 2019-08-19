package main

import "fmt"

func removeElement(nums []int, val int) int {
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}

func removeElement1(nums []int, val int) int {
	left, right, size := 0, 0, len(nums)
	for ; right < size; right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement1(nums, 3))
	fmt.Println(nums)
}
