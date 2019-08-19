package main

import "fmt"

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := index[target - b]; ok {
			return []int{i, j}
		}
		index[b] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if i != j && x + y == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, x := range nums {
		index[x] = i
	}

	for i, b := range nums {
		if j, ok := index[target - b]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))
}