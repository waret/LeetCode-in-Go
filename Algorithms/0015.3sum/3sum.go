package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					a, b, c := nums[i], nums[j], nums[k]
					if a < b {
						a, b = b, a
					}
					if a < c {
						a, c = c, a
					}
					if b < c {
						b, c = c, b
					}
					if len(res) == 0 {
						res = append(res, []int{a, b, c})
					} else {
						exist := false
						for _, r := range res {
							if r[0] == a && r[1] == b && r[2] == c {
								exist = true
							}
						}
						if !exist {
							res = append(res, []int{a, b, c})
						}
					}
				}
			}
		}
	}
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s > 0:
				r--
			case s < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r {
					if nums[l] == nums[l+1] {
						l++
					} else if nums[r] == nums[r-1] {
						r--
					} else {
						l++
						r--
						break
					}
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s > 0:
				r--
			case s < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				c := true
				for l < r && c {
					switch {
					case nums[l] == nums[l+1]:
						l++
					case nums[r] == nums[r-1]:
						r--
					default:
						l++
						r--
						c = false
					}
				}
			}
		}
	}
	return res
}

func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s > 0:
				r--
			case s < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && (nums[l] == nums[l+1] || nums[r] == nums[r-1]) {
					if nums[l] == nums[l+1] {
						l++
					}
					if nums[r] == nums[r-1] {
						r--
					}
				}
				l++
				r--
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%+v", []int{-1, 0, 1, 2, -1, -4})
	threeSum1([]int{-1, 0, 1, 2, -1, -4})
}
