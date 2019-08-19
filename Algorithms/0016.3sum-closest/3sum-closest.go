package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if delta > target-s {
					delta = target - s
					res = s
				}
			case s > target:
				r--
				if delta > s-target {
					delta = s - target
					res = s
				}
			default:
				return s
			}
		}
	}
	return res
}

func main() {
	threeSumClosest([]int{-1, 0, 1, 1, 55}, 3)
}
