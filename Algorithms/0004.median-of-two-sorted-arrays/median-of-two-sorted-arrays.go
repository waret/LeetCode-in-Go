package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = len(nums1), len(nums2)
	}

	imin, imax, half := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := half - i

		if i > imin && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else if i < imax && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else {
			// 找到位置，计算中位数
			var maxleft, minright int
			if i == 0 {
				maxleft = nums2[j-1]
			} else if j == 0 {
				// m == n 时出现
				maxleft = nums1[i-1]
			} else {
				maxleft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				// 总数为奇数时，左侧比右侧多的一个元素，也就是目标中位数
				return float64(maxleft)
			}

			if i == m {
				minright = nums2[j]
			} else if j == n {
				// m == n 时出现
				minright = nums1[i]
			} else {
				minright = min(nums1[i], nums2[j])
			}

			return float64(maxleft+minright) / 2.0
		}
	}
	return 0.0
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		} else if j == lenNjs ||
			(i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		} else {
			panic("不会进入该分支")
		}
	}

	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	var left []int
	var index1, index2, count int

	total := len(nums1) + len(nums2)
	flag := total / 2

	for count <= flag {
		if index1 < len(nums1) && index2 < len(nums2) {
			if nums1[index1] < nums2[index2] {
				left = append(left, nums1[index1])
				index1++
			} else {
				left = append(left, nums2[index2])
				index2++
			}
			count++
		}
		if index1 == len(nums1) && len(left) <= flag {
			left = append(left, nums2[index2:flag-count+index2+1]...)
			break
		}
		if index2 == len(nums2) && len(left) <= flag {
			left = append(left, nums1[index1:flag-count+index1+1]...)
			break
		}
	}
	if total%2 != 0 {
		return float64(left[len(left)-1])
	} else {
		return float64(left[len(left)-2]+left[len(left)-1]) / 2
	}
}

func main() {
	findMedianSortedArrays([]int{}, []int{})
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 5}, []int{2, 3, 4, 6}))
}
